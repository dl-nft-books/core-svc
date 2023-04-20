package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/core-svc/solidity/generated/marketplace"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"math/big"
	"net/http"
	"time"
)

func BuyWithVoucher(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewBuyWithVoucherRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign mint request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Getting task's mintInfo
	task, err := helpers.DB(r).Tasks().GetById(request.Data.Attributes.TaskId)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Getting book's mintInfo
	book, err := helpers.Booker(r).GetBookById(task.BookId, task.ChainId)
	if err != nil {
		logger.WithError(err).Error("failed to get a book")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		logger.Warnf("Book with specified id %d in network with chain id %d was not found", task.BookId, task.ChainId)
		ape.RenderErr(w, problems.NotFound())
		return
	}
	// Forming signature mintInfo
	mintConfig := helpers.Minter(r)

	network, err := helpers.Networker(r).GetNetworkDetailedByChainID(task.ChainId)
	if err != nil {
		logger.WithError(err).Error("failed to get network")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if network == nil {
		logger.Error("network with such id doesn't exists")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	contractRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
	if err != nil {
		logger.WithError(err).Error("failed to create contract registry")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	marketplaceContractAddress, err := contractRegistry.GetMarketplaceContract(nil)
	if err != nil {
		logger.WithError(err).Error("failed to get marketplace contract name")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	domainData := signature.EIP712DomainData{
		VerifyingAddress: marketplaceContractAddress.String(),
		ContractName:     "Marketplace",
		ContractVersion:  "1",
		ChainID:          task.ChainId,
	}

	mintInfo := signature.MintInfo{
		TokenContract:    book.Data.Attributes.Networks[0].Attributes.ContractAddress,
		TokenRecipient:   task.Account,
		TokenId:          task.TokenId,
		TokenAddress:     request.Data.Attributes.VoucherAddress,
		TokenURI:         task.MetadataIpfsHash,
		EndTimestamp:     time.Now().Add(mintConfig.Expiration).Unix(),
		PricePerOneToken: big.NewInt(0),
		Discount:         big.NewInt(0),
	}

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	publicKeyECDSA, ok := helpers.Transacter(r).PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		logger.WithError(err).Error("failed to convert public key to ECDSA")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	nonce, err := network.RpcUrl.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		logger.WithError(err).Error("failed to get nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	gasPrice, err := network.RpcUrl.SuggestGasPrice(context.Background())
	if err != nil {
		logger.WithError(err).Error("failed to get gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(helpers.Transacter(r).PrivateKey, big.NewInt(task.ChainId))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = helpers.Transacter(r).GasLimit
	auth.GasPrice = gasPrice

	if err != nil {
		logger.WithError(err).Error("failed to create transaction signer")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	permitSig := marketplace.IMarketplaceSigData{
		EndSigTimestamp: request.Data.Attributes.EndSigTimestamp,
		V:               uint8(request.Data.Attributes.PermitSig.Attributes.V),
	}
	copy(permitSig.R[:], request.Data.Attributes.PermitSig.Attributes.R)
	copy(permitSig.S[:], request.Data.Attributes.PermitSig.Attributes.S)

	sig := marketplace.IMarketplaceSigData{
		EndSigTimestamp: big.NewInt(mintInfo.EndTimestamp),
		V:               uint8(mintSignature.V),
	}
	copy(sig.R[:], mintSignature.R)
	copy(sig.S[:], mintSignature.S)

	buyParams := marketplace.IMarketplaceBuyParams{
		TokenContract: common.HexToAddress(mintInfo.TokenContract),
		Recipient:     common.HexToAddress(mintInfo.TokenRecipient),
		PaymentDetails: marketplace.IMarketplacePaymentDetails{
			PaymentTokenAddress: common.HexToAddress(mintInfo.TokenAddress),
			PaymentTokenPrice:   mintInfo.PricePerOneToken,
			Discount:            mintInfo.Discount,
			NftTokenId:          big.NewInt(0),
		},
		TokenData: marketplace.IERC721MintableTokenTokenMintData{
			TokenId:  big.NewInt(mintInfo.TokenId),
			TokenURI: mintInfo.TokenURI,
		},
	}

	transactor, err := marketplace.NewMarketplaceTransactor(marketplaceContractAddress, network.RpcUrl)
	if err != nil {
		logger.WithError(err).Error("failed to create marketplace transactor")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	fmt.Println("domainData")
	spew.Dump(domainData)
	fmt.Println("mintInfo")
	spew.Dump(mintInfo)
	fmt.Println("mintSignature")
	spew.Dump(mintSignature)
	fmt.Println("sig")
	spew.Dump(sig)
	fmt.Println("buyParams")
	spew.Dump(buyParams)

	tx, err := transactor.BuyTokenWithVoucher(auth, buyParams, sig, permitSig)

	if err != nil {
		logger.WithError(err).Error("failed to generate send transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	receipt, err := bind.WaitMined(context.Background(), network.RpcUrl, tx)
	fmt.Printf("transaction mined in block %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("transaction mined with hash %v\n", receipt.TxHash)
	fmt.Printf("transaction mined with status %v\n", receipt.Status)
	if err != nil {
		logger.WithError(err).Error("failed to submit transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/core-svc/solidity/generated/marketplace"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"math/big"
	"net/http"
	"time"
)

const receiptStatusOk = 1

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

	gasPrice, err := getGasPrice(r, network.RpcUrl)
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
	decodeR, err := hexutil.Decode(request.Data.Attributes.PermitSig.Attributes.R)
	decodeS, err := hexutil.Decode(request.Data.Attributes.PermitSig.Attributes.S)
	copy(permitSig.R[:], decodeR)
	copy(permitSig.S[:], decodeS)

	sig := marketplace.IMarketplaceSigData{
		EndSigTimestamp: big.NewInt(mintInfo.EndTimestamp),
		V:               mintSignature[64],
	}
	copy(sig.R[:], mintSignature[:32])
	copy(sig.S[:], mintSignature[32:64])

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

	helpers.Log(r).Debug(fmt.Sprintf("Traing to send transaction. Tx gas price is %v", gasPrice))
	transactor, err := marketplace.NewMarketplaceTransactor(marketplaceContractAddress, network.RpcUrl)
	if err != nil {
		logger.WithError(err).Error("failed to create marketplace transactor")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	tx, err := transactor.BuyTokenWithVoucher(auth, buyParams, sig, permitSig)

	if err != nil {
		logger.WithError(err).Error("failed to generate send transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewBuyVoucherResponse(tx.Hash().String()))

}
func getGasPrice(r *http.Request, rpc *ethclient.Client) (*big.Int, error) {
	gasPrice, err := rpc.SuggestGasPrice(context.Background())
	helpers.Log(r).Debug(fmt.Sprintf("On chain gas price is %v ", gasPrice))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas price")
	}

	if helpers.Transacter(r).MaxGasPrice.Cmp(gasPrice) != -1 {
		return gasPrice, nil
	}

	return helpers.Transacter(r).MaxGasPrice, nil

}

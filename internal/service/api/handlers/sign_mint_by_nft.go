package handlers

import (
	"fmt"
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"time"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func SignMintByNft(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewSignMintByNftRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign mint by nft request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Getting task's mintInfo
	task, err := helpers.DB(r).Tasks().GetById(request.TaskID)
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
	contractName, err := contractRegistry.MARKETPLACENAME(nil)
	if err != nil {
		logger.WithError(err).Error("failed to get marketplace contract name")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	domainData := signature.EIP712DomainData{
		VerifyingAddress: book.Data.Attributes.Networks[0].Attributes.ContractAddress,
		ContractName:     contractName,
		ContractVersion:  "1",
		ChainID:          book.Data.Attributes.Networks[0].Attributes.ChainId,
	}
	mintInfo := signature.MintInfo{
		TokenAddress: request.NftAddress,
		Discount:     big.NewInt(0),
		TokenURI:     task.MetadataIpfsHash,
	}

	// Getting price per token in dollars
	priceResponse, err := helpers.Pricer(r).GetNftPrice(request.Platform, request.NftAddress)
	if err != nil {
		logger.WithError(err).Error("failed to get nft floor price")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	// Converting price
	mintInfo.PricePerOneToken, err = helpers.ConvertPrice(fmt.Sprintf("%f", priceResponse.Data.Attributes.Usd), mintConfig.Precision)
	if err != nil {
		logger.WithError(err).Error("failed to convert price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	mintInfo.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignMintResponse(
		mintInfo.PricePerOneToken.String(),
		mintInfo.Discount.String(),
		mintSignature,
		mintInfo.EndTimestamp,
		mintInfo.TokenId,
	))
}

package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/jsonerrors"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"time"
)

func SignAcceptNftRequest(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewSignAcceptNftRequestRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign accept nft_request request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequest, err := helpers.DB(r).NftRequests().FilterById(request.NftRequestID).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get nft_request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nftRequest == nil {
		ape.RenderErr(w, jsonerrors.WithDetails(problems.NotFound(), jsonerrors.NftRequestNotFound))
		return
	}
	if nftRequest.Status != resources.RequestAccepted {
		ape.RenderErr(w, jsonerrors.WithDetails(problems.Forbidden(), jsonerrors.NftRequestNotApprovedByAdmin))
		return
	}

	// Getting task's info
	task, err := helpers.DB(r).Tasks().GetById(request.TaskID)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, jsonerrors.WithDetails(problems.NotFound(), jsonerrors.TaskNotFound))
		return
	}

	// Forming signature acceptInfo
	acceptConfig := helpers.Accepter(r)

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
	acceptInfo := signature.AcceptInfo{
		RequestId:      nftRequest.Id,
		TokenId:        task.TokenId,
		TokenURI:       task.MetadataIpfsHash,
		TokenRecipient: task.Account,
		EndTimestamp:   time.Now().Add(acceptConfig.Expiration).Unix(),
	}

	signerDataConfig := helpers.SignererData(r)
	// Signing the mint transaction
	acceptSignature, err := signature.SignAcceptInfo(&acceptInfo, &domainData, &signerDataConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 accept signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	RSV, err := signature.ParseSignatureParameters(acceptSignature)
	if err != nil {
		logger.WithError(err).Error("failed to parse signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignAcceptResponse(RSV))
}

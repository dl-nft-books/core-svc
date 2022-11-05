package handlers

import (
	"math/big"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/signature"
)

const tokenIdIncrementKey = "token_id_increment"

func SignCreate(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewSignCreateRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// generating id
	tokenKV, err := helpers.GeneratorDB(r).KeyValue().Get(tokenIdIncrementKey)
	if err != nil {
		logger.WithError(err).Debug("failed to generate token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if tokenKV == nil {
		tokenKV = &data.KeyValue{
			Key:   tokenIdIncrementKey,
			Value: "0",
		}
	}

	lastTokenContractID, err := strconv.ParseInt(tokenKV.Value, 10, 64)
	if err != nil {
		logger.WithError(err).Debug("failed to parse token id from kv")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// forming signature createInfo
	mintConfig := helpers.Minter(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: mintConfig.TokenFactoryAddress,
		ContractName:     mintConfig.TokenFactoryName,
		ContractVersion:  mintConfig.TokenFactoryVersion,
		ChainID:          mintConfig.ChainID,
	}

	tokenPrice, ok := big.NewInt(0).SetString(req.PricePerOneToken, 10)
	if !ok {
		logger.Error("failed to cast price to big.Int")
		ape.RenderErr(w, problems.InternalError())
	}

	createInfo := signature.CreateInfo{
		TokenContractId:  lastTokenContractID + 1,
		TokenName:        req.TokenName,
		TokenSymbol:      req.TokenSymbol,
		PricePerOneToken: tokenPrice,
	}

	// signing
	signature, err := signature.SignCreateInfo(&createInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Debug("failed to generate eip712 create signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// update kv if operation is ended successfully
	if err = helpers.GeneratorDB(r).KeyValue().Upsert(data.KeyValue{
		Key:   tokenIdIncrementKey,
		Value: strconv.FormatInt(createInfo.TokenContractId, 10),
	}); err != nil {
		logger.WithError(err).Debug("failed to update last created token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignCreateResponse(createInfo.TokenContractId, signature))
}

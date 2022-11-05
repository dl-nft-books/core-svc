package handlers

import (
	"math/big"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/signature"
)

func SignCreate(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewSignCreateRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// TODO: generate id

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
		logger.Debug("failed to cast price to big.Int")
		ape.RenderErr(w, problems.InternalError())
	}

	createInfo := signature.CreateInfo{
		TokenContractId:  0,
		TokenName:        req.TokenName,
		TokenSymbol:      req.TokenSymbol,
		PricePerOneToken: tokenPrice,
	}

	// signing
	signature, err := signature.SignCreateInfo(&createInfo, &domainData, &mintConfig)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignCreateResponse(createInfo.TokenContractId, signature))
}

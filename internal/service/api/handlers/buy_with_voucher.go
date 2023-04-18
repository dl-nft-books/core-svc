package handlers

//
//import (
//	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
//	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
//	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
//	"github.com/dl-nft-books/core-svc/internal/signature"
//	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
//	"github.com/dl-nft-books/core-svc/solidity/generated/marketplace"
//	"github.com/ethereum/go-ethereum/common"
//	"gitlab.com/distributed_lab/ape"
//	"gitlab.com/distributed_lab/ape/problems"
//	"gitlab.com/distributed_lab/logan/v3"
//	"net/http"
//	"time"
//)
//
//func BuyWithVoucher(w http.ResponseWriter, r *http.Request) {
//	logger := helpers.Log(r)
//
//	request, err := requests.NewSignMintRequest(r)
//	if err != nil {
//		logger.WithError(err).Error("failed to fetch new sign mint request")
//		ape.RenderErr(w, problems.BadRequest(err)...)
//		return
//	}
//
//	// Getting task's mintInfo
//	task, err := helpers.DB(r).Tasks().GetById(request.TaskID)
//	if err != nil {
//		logger.WithError(err).Error("failed to get task")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//	if task == nil {
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//
//	// Getting book's mintInfo
//	book, err := helpers.Booker(r).GetBookById(task.BookId, task.ChainId)
//	if err != nil {
//		logger.WithError(err).Error("failed to get a book")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//	if book == nil {
//		logger.Warnf("Book with specified id %d in network with chain id %d was not found", task.BookId, task.ChainId)
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//	// Forming signature mintInfo
//	mintConfig := helpers.Minter(r)
//
//	network, err := helpers.Networker(r).GetNetworkDetailedByChainID(task.ChainId)
//	if err != nil {
//		logger.WithError(err).Error("failed to get network")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//	if network == nil {
//		logger.Error("network with such id doesn't exists")
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//	contractRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
//	if err != nil {
//		logger.WithError(err).Error("failed to create contract registry")
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//	marketplaceContractAddress, err := contractRegistry.GetMarketplaceContract(nil)
//	if err != nil {
//		logger.WithError(err).Error("failed to get marketplace contract name")
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//	transactor, err := marketplace.NewMarketplaceTransactor(marketplaceContractAddress, network.RpcUrl)
//	if err != nil {
//		logger.WithError(err).Error("failed to create marketplace transactor")
//		ape.RenderErr(w, problems.NotFound())
//		return
//	}
//	transaction, err := transactor.BuyTokenWithVoucher(nil)
//	domainData := signature.EIP712DomainData{
//		VerifyingAddress: marketplaceContractAddress.String(),
//		ContractName:     "Marketplace",
//		ContractVersion:  "1",
//		ChainID:          task.ChainId,
//	}
//	mintInfo := signature.MintInfo{
//		TokenContract: book.Data.Attributes.Networks[0].Attributes.ContractAddress,
//		TokenId:       task.TokenId,
//		TokenAddress:  request.TokenAddress,
//		TokenURI:      task.MetadataIpfsHash,
//		EndTimestamp:  time.Now().Add(mintConfig.Expiration).Unix(),
//	}
//
//	mintInfo.PricePerOneToken, err = getPricePerOneToken(w, r, request, *book, mintConfig.Precision)
//	if err != nil {
//		logger.WithError(err).Error("failed to get price")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//
//	// Getting promocode info
//	promocode, err := helpers.DB(r).Promocodes().FilterById(request.PromocodeID).Get()
//	if err != nil {
//		logger.WithError(err).Error("failed to get promocode")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//	mintInfo.Discount, ok = getPromocodeDiscount(w, r, promocode)
//	if !ok {
//		return
//	}
//
//	// Signing the mint transaction
//	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
//	if err != nil {
//		logger.WithError(err).Error("failed to generate eip712 mint signature")
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//
//	// Using promocode after signature is formed
//	if promocode != nil {
//		if err = helpers.DB(r).Promocodes().New().UpdateUsages(promocode.Usages + 1).FilterUpdateById(promocode.Id).Update(); err != nil {
//			logger.WithError(err).WithFields(logan.F{"promocode": promocode.Promocode}).Error("failed to update promocode")
//			ape.RenderErr(w, problems.InternalError())
//			return
//		}
//
//		logger.Info("promocode applied, discount: ", mintInfo.Discount.String())
//	}
//
//	ape.Render(w, responses.NewSignMintResponse(
//		mintInfo.PricePerOneToken.String(),
//		mintInfo.Discount.String(),
//		mintSignature,
//		mintInfo.EndTimestamp,
//		mintInfo.TokenId,
//	))
//}

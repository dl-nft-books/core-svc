package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/data/postgres"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/jsonerrors"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/core-svc/solidity/generated/erc721mintabletoken"
	"github.com/dl-nft-books/core-svc/solidity/generated/marketplace"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateTaskRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !validateCreateTaskRequest(request, w, r) {
		return
	}

	bookId := request.Data.Attributes.BookId

	// Check if book exists
	getBookResponse, err := helpers.Booker(r).GetBookById(bookId, request.Data.Attributes.ChainId)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to get book with id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if getBookResponse == nil {
		helpers.Log(r).Info("corresponding book was not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	network, err := helpers.Networker(r).GetNetworkDetailedByChainID(request.Data.Attributes.ChainId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get network")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if network == nil {
		helpers.Log(r).Error("network with such id doesn't exists")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	contractRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create contract registry")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	marketplaceContractAddress, err := contractRegistry.GetMarketplaceContract(nil)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get marketplace contract address")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	marketplaceContract, err := marketplace.NewMarketplace(marketplaceContractAddress, network.RpcUrl)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create contract registry")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	books, err := marketplaceContract.GetBaseTokenParams(nil, []common.Address{common.HexToAddress(getBookResponse.Data.Attributes.Networks[0].Attributes.ContractAddress)})
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get book from contract")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if books == nil {
		helpers.Log(r).Error("book with such contract address doesn't exists")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	erc721Contract, err := erc721mintabletoken.NewErc721mintabletoken(books[0].TokenContract, network.RpcUrl)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create new erc721 mintable token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	tokenId, err := erc721Contract.NextTokenId(nil)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get future token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	// Then creating task
	createdTaskId, err := helpers.DB(r).Tasks().Insert(data.Task{
		BookId:  bookId,
		ChainId: network.ChainId,
		//Banner:    request.Data.Attributes.Banner,
		TokenId:   tokenId.Int64(),
		Account:   request.Data.Attributes.Account,
		TokenName: books[0].TokenName,
		Status:    resources.TaskPending,
		CreatedAt: time.Now(),
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new task with book id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(createdTaskId, resources.TASKS),
	})
}

func validateCreateTaskRequest(request *requests.CreateTaskRequest, w http.ResponseWriter, r *http.Request) (ok bool) {
	var (
		database     = helpers.DB(r)
		restrictions = helpers.ApiRestrictions(r)
		statusFilter = resources.TaskFinishedGeneration
	)

	// Validating if a user has generated too many unpaid books
	tasks, err := database.Tasks().
		Sort(pgdb.Sorts{postgres.TasksCreatedAt}).
		Select(data.TaskSelector{
			Account: &request.Data.Attributes.Account,
			Status:  &statusFilter,
		})
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get select tasks from the database")
		ape.RenderErr(w, problems.InternalError())
		return false
	}

	tasksNumber := len(tasks)

	// If no tasks were found - simply return that everything is ok
	if tasksNumber == 0 {
		return true
	}

	if uint64(tasksNumber) >= restrictions.MaxFailedAttempts {
		ape.RenderErr(w, jsonerrors.WithDetails(
			problems.BadRequest(errors.New("max amount of tries exceeded"))[0],
			jsonerrors.ApiMaxTriesExceeded,
		))
		return false
	}

	// Validating how often user tries to buy a book
	lastCreatedAt := tasks[tasksNumber-1].CreatedAt

	durationAfterPreviousAttempt := time.Now().Sub(lastCreatedAt)
	if durationAfterPreviousAttempt < restrictions.RequestDelay {
		ape.RenderErr(w, jsonerrors.WithDetails(
			problems.BadRequest(errors.New("task delay not passed"))[0],
			jsonerrors.TaskDelayNotPassed,
		))
		return false
	}

	return true
}

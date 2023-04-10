package task_processor

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/data/opensea"

	"github.com/dl-nft-books/core-svc/internal/data"
	runnerHelpers "github.com/dl-nft-books/core-svc/internal/service/runners/helpers"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var bookNotFoundErr = errors.New("book was not found")

func (p *TaskProcessor) handleTask(task data.Task) error {
	p.logger.Debugf("Started processing task with id of %d", task.Id)
	p.logger.Debug("Retrieving book...")

	// Making an api request to retrieve a book
	response, err := p.booksApi.GetBookById(task.BookId, task.ChainId)

	if err != nil {
		return errors.Wrap(err, "failed to get book by id")
	}
	if response == nil {
		return errors.From(bookNotFoundErr, logan.F{"book_id": task.BookId})
	}

	p.logger.Debug("Book retrieved successfully")

	p.logger.Debug("Calculating banner IPFS Hash...")

	ipfsBannerHash, err := runnerHelpers.PrecalculateIPFSHash(task.Banner)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate IPFS hash")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsBannerHash))

	if err = p.db.Tasks().UpdateBannerIpfsHash(ipfsBannerHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update banner ipfs hash")
	}

	p.logger.Debug("Banner IPFS Hash calculated successfully")
	p.logger.Debug("Uploading banner to S3...")

	statusCode, err := p.documenter.UploadDocument(task.Banner, ipfsBannerHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload banner")
	}
	if statusCode != http.StatusOK {
		return errors.From(errors.New("failed to upload banner"), logan.F{"status code": statusCode})
	}

	p.logger.Debug("Banner downloaded successfully")
	p.logger.Debug("Retrieving document key...")

	fileKey := response.Data.Attributes.File.Attributes.Key

	p.logger.Debug("Key retrieved successfully")
	p.logger.Debug("Retrieving document link for metadata...")

	fileLink, err := p.documenter.GetDocumentLink(fileKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	p.logger.Debug("Document link retrieved successfully")
	p.logger.Debug("Calculating metadata IPFS Hash...")

	openseaData := opensea.Metadata{
		Name:        fmt.Sprintf("%s #%v", task.TokenName, task.Id),
		Description: response.Data.Attributes.Description,
		Image:       p.ipfser.BaseUri + ipfsBannerHash,
		FileURL:     fileLink.Data.Attributes.Url,
	}
	ipfsMetadataHash, err := runnerHelpers.PrecalculateMetadataIPFSHash(openseaData)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate ipfs hash for a metadata file")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsMetadataHash))

	if err = p.db.Tasks().UpdateMetadataIpfsHash(ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update metadata ipfs hash")
	}
	if err = p.db.Tasks().UpdateUri(p.ipfser.BaseUri + ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update task uri")
	}

	p.logger.Debug("Metadata IPFS Hash calculated successfully")
	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)

	return nil
}

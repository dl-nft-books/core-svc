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
	response, err := p.booksApi.GetBookById(task.BookId)

	if err != nil {
		return errors.Wrap(err, "failed to get book by id")
	}
	if response == nil {
		return errors.From(bookNotFoundErr, logan.F{"book_id": task.BookId})
	}

	p.logger.Debug("Book retrieved successfully")
	p.logger.Debug("Retrieving document key...")

	bannerKey := response.Data.Attributes.File.Attributes.Key

	p.logger.Debug("Key retrieved successfully")
	p.logger.Debugf("Retrieving document link from S3... (fileKey=%s)", bannerKey)

	bannerLink, err := p.documenter.GetDocumentLink(bannerKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link", logan.F{
			"banner_key": bannerKey,
		})
	}

	p.logger.Debug("Link retrieved successfully")
	p.logger.Debugf("Downloading document...")

	rawDocument, err := runnerHelpers.DownloadDocument(bannerLink.Data.Attributes.Url)
	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}

	p.logger.Debug("Document downloaded successfully...")

	p.logger.Debug("Calculating document IPFS Hash...")

	ipfsBannerHash, err := runnerHelpers.PrecalculateIPFSHash(rawDocument)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate IPFS hash")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsBannerHash))

	if err = p.db.Tasks().UpdateBannerIpfsHash(ipfsBannerHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}

	p.logger.Debug("Document IPFS Hash calculated successfully")
	p.logger.Debug("Uploading document to S3...")

	statusCode, err := p.documenter.UploadDocument(rawDocument, ipfsBannerHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}
	if statusCode != http.StatusOK {
		return errors.From(errors.New("failed to upload file"), logan.F{"status code": statusCode})
	}

	p.logger.Debug("Document downloaded successfully")
	p.logger.Debug("Retrieving banner key...")

	fileKey := response.Data.Attributes.File.Attributes.Key

	p.logger.Debug("Key retrieved successfully")
	p.logger.Debug("Retrieving banner link for metadata...")

	fileLink, err := p.documenter.GetDocumentLink(fileKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	p.logger.Debug("Banner link retrieved successfully")
	p.logger.Debug("Calculating metadata IPFS Hash...")

	var (
		//bookTitle       = response.Data.Attributes.Title
		bookDescription = response.Data.Attributes.Description
	)
	openseaData := opensea.Metadata{
		//Name:        fmt.Sprintf("%s #%v", bookTitle, task.Id),
		Description: bookDescription,
		Image:       p.ipfser.BaseUri + ipfsBannerHash,
		FileURL:     fileLink.Data.Attributes.Url,
	}
	ipfsMetadataHash, err := runnerHelpers.PrecalculateMetadataIPFSHash(openseaData)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate ipfs hash for a metadata file")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsMetadataHash))

	if err = p.db.Tasks().UpdateMetadataIpfsHash(ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}
	if err = p.db.Tasks().UpdateUri(p.ipfser.BaseUri + ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update task uri")
	}

	p.logger.Debug("Metadata IPFS Hash calculated successfully")
	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)

	return nil
}

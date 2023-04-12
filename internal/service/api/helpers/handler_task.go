package helpers

import (
	"fmt"
	"github.com/dl-nft-books/core-svc/internal/data/opensea"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/data"
	runnerHelpers "github.com/dl-nft-books/core-svc/internal/service/runners/helpers"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var bookNotFoundErr = errors.New("book was not found")

func HandleTask(r *http.Request, logger *logan.Entry, task data.Task, banner []byte) error {
	var (
		booker     = Booker(r)
		documenter = DocumenterConnector(r)
		db         = DB(r)
		ipfser     = Ipfser(r)
	)

	logger.Debugf("Started processing task with id of %d", task.Id)
	logger.Debug("Retrieving book...")

	// Making an api request to retrieve a book
	response, err := booker.GetBookById(task.BookId, task.ChainId)

	if err != nil {
		return errors.Wrap(err, "failed to get book by id")
	}
	if response == nil {
		return errors.From(bookNotFoundErr, logan.F{"book_id": task.BookId})
	}

	logger.Debug("Book retrieved successfully")

	logger.Debug("Calculating banner IPFS Hash...")
	ipfsBannerHash, err := runnerHelpers.PrecalculateIPFSHash(banner)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate IPFS hash")
	}

	logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsBannerHash))

	if err = db.Tasks().UpdateBannerIpfsHash(ipfsBannerHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update banner ipfs hash")
	}

	logger.Debug("Banner IPFS Hash calculated successfully")
	logger.Debug("Uploading banner to S3...")

	statusCode, err := documenter.UploadDocument(banner, ipfsBannerHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload banner")
	}
	if statusCode != http.StatusOK {
		return errors.From(errors.New("failed to upload banner"), logan.F{"status code": statusCode})
	}

	logger.Debug("Banner downloaded successfully")
	logger.Debug("Retrieving document key...")

	fileKey := response.Data.Attributes.File.Attributes.Key

	logger.Debug("Key retrieved successfully")
	logger.Debug("Retrieving document link for metadata...")

	fileLink, err := documenter.GetDocumentLink(fileKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	logger.Debug("Document link retrieved successfully")
	logger.Debug("Calculating metadata IPFS Hash...")

	openseaData := opensea.Metadata{
		Name:        fmt.Sprintf("%s #%d", task.TokenName, task.TokenId),
		Description: response.Data.Attributes.Description,
		Image:       ipfser.BaseUri + ipfsBannerHash,
		FileURL:     fileLink.Data.Attributes.Url,
	}
	ipfsMetadataHash, err := runnerHelpers.PrecalculateMetadataIPFSHash(openseaData)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate ipfs hash for a metadata file")
	}

	logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsMetadataHash))

	if err = db.Tasks().UpdateMetadataIpfsHash(ipfsMetadataHash).UpdateUri(ipfser.BaseUri + ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update metadata ipfs hash")
	}

	logger.Debug("Metadata IPFS Hash calculated successfully")
	logger.Debugf("Successfully finished processing task with id of %d", task.Id)

	return nil
}

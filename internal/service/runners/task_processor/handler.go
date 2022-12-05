package task_processor

import (
	"bytes"
	"fmt"
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data/opensea"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/pdf"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/helpers"
)

const baseURI = "https://ipfs.io/ipfs/"

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

	fileKey := response.Data.Attributes.File.Attributes.Key

	p.logger.Debug("Key retrieved successfully")
	p.logger.Debug("Retrieving document link from S3...")

	fileLink, err := p.documenter.GetDocumentLink(fileKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	p.logger.Debug("Link retrieved successfully")
	p.logger.Debug("Downloading document...")

	rawDocument, err := helpers.DownloadDocument(fileLink.Data.Attributes.Url)
	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}

	p.logger.Debug("Document downloaded successfully...")
	p.logger.Debug("Generating signature...")

	reader := bytes.NewReader(rawDocument)
	pdfSignatureGenerator := pdf.New(p.signatureParams)
	rawDocumentWithSignature, err := pdfSignatureGenerator.GenerateSignature(reader, task.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to generate signature")
	}

	p.logger.Debug("Signature generated successfully")
	p.logger.Debug("Calculating document IPFS Hash...")

	ipfsFileHash, err := helpers.PrecalculateIPFSHash(rawDocumentWithSignature)
	if err != nil {
		return errors.Wrap(err, "failed to precalculate IPFS hash")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsFileHash))

	if err = p.db.Tasks().UpdateFileIpfsHash(ipfsFileHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}

	p.logger.Debug("Document IPFS Hash calculated successfully")
	p.logger.Debug("Uploading document to S3...")

	statusCode, err := p.documenter.UploadDocument(rawDocumentWithSignature, ipfsFileHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}
	if statusCode != http.StatusOK {
		return errors.From(errors.New("failed to upload file"), logan.F{"status code": statusCode})
	}

	p.logger.Debug("Document downloaded successfully")
	p.logger.Debug("Retrieving banner key...")

	bannerKey := response.Data.Attributes.Banner.Attributes.Key

	p.logger.Debug("Key retrieved successfully")
	p.logger.Debug("Retrieving banner link for metadata...")

	bannerLink, err := p.documenter.GetDocumentLink(bannerKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	p.logger.Debug("Banner link retrieved successfully")
	p.logger.Debug("Calculating metadata IPFS Hash...")

	var (
		bookTitle       = response.Data.Attributes.Title
		bookDescription = response.Data.Attributes.Description
	)

	ipfsMetadataHash, err := helpers.PrecalculateMetadataIPFSHash(opensea.Metadata{
		Name:        fmt.Sprintf("%s #%v", bookTitle, task.Id),
		Description: bookDescription,
		Image:       bannerLink.Data.Attributes.Url,
		FileURL:     baseURI + ipfsFileHash,
	})
	if err != nil {
		return errors.Wrap(err, "failed to precalculate ipfs hash for a metadata file")
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsMetadataHash))

	if err = p.db.Tasks().UpdateMetadataIpfsHash(ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}
	if err = p.db.Tasks().UpdateUri(baseURI + ipfsMetadataHash).Update(task.Id); err != nil {
		return errors.Wrap(err, "failed to update task uri")
	}

	p.logger.Debug("Metadata IPFS Hash calculated successfully")
	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)

	return nil
}

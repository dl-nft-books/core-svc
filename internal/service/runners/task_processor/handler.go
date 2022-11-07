package task_processor

import (
	"bytes"
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/pdf_signature_generator"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/models"
)

const baseURI = "https://ipfs.io/ipfs/"

func (p *TaskProcessor) handleTask(task data.Task) error {
	// updating db
	p.booksDB = p.booksDB.New()

	p.logger.Debugf("Started processing task with id of %d", task.Id)

	p.logger.Debug("Retrieving book...")

	book, err := p.booksDB.FilterActual().FilterByID(task.BookId).Get()

	if err != nil {
		return err
	}
	if book == nil {
		return errors.From(errors.New("book not found"), logan.F{"book_id": task.BookId})
	}

	p.logger.Debug("Book retrieved successfully")

	p.logger.Debug("Retrieving document key...")

	fileKey, err := helpers.GetDocumentKey(book.File)
	if err != nil {
		return errors.Wrap(err, "failed to get document key")
	}
	if fileKey == nil {
		return errors.New("failed to get document key")
	}

	p.logger.Debug("Key retrieved successfully")

	p.logger.Debug("Retrieving document link from S3...")

	fileLink, err := p.documenterConnector.GetDocumentLink(*fileKey)
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
	pdfSignatureGenerator := pdf_signature_generator.New(p.signatureParams)
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

	if err = p.generatorDB.Tasks().UpdateFileIpfsHash(ipfsFileHash, task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}

	p.logger.Debug("Document IPFS Hash calculated successfully")

	p.logger.Debug("Uploading document to S3...")

	statusCode, err := p.documenterConnector.UploadDocument(rawDocumentWithSignature, ipfsFileHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}
	if statusCode != http.StatusOK {
		return errors.From(errors.New("failed to upload file"), logan.F{"status code": statusCode})
	}

	p.logger.Debug("Document downloaded successfully")

	p.logger.Debug("Retrieving banner key...")

	bannerKey, err := helpers.GetDocumentKey(book.Banner)
	if err != nil {
		return errors.Wrap(err, "failed to get document key")
	}
	if bannerKey == nil {
		return errors.New("failed to get document key")
	}

	p.logger.Debug("Key retrieved successfully")

	p.logger.Debug("Retrieving banner link for metadata...")

	bannerLink, err := p.documenterConnector.GetDocumentLink(*bannerKey)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	p.logger.Debug("Banner link retrieved successfully")

	p.logger.Debug("Calculating metadata IPFS Hash...")

	ipfsMetadataHash, err := helpers.PrecalculateMetadataIPFSHash(models.Metadata{
		Name:        fmt.Sprintf("%s #%v", book.Title, task.Id),
		Description: book.Description,
		Image:       bannerLink.Data.Attributes.Url,
		FileURL:     baseURI + ipfsFileHash,
	})
	if err != nil {
		return err
	}

	p.logger.Debug(fmt.Sprintf("Precalculated IPFS hash: %s", ipfsMetadataHash))

	if err = p.generatorDB.Tasks().UpdateMetadataIpfsHash(ipfsMetadataHash, task.Id); err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}

	if err = p.generatorDB.Tasks().UpdateUri(baseURI+ipfsMetadataHash, task.Id); err != nil {
		return errors.Wrap(err, "failed to update task uri")
	}

	p.logger.Debug("Metadata IPFS Hash calculated successfully")

	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)

	return nil
}

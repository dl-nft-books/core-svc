package task_processor

import (
	"bytes"
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/pdf_signature_generator"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/helpers"
)

func (p *TaskProcessor) handleTask(task data.Task) error {
	p.logger.Debugf("Started processing task with id of %d", task.Id)

	// Getting book
	p.logger.Debug("Retrieving book...")
	book, err := p.booksDB.FilterActual().FilterByID(task.BookId).Get()
	if err != nil {
		return err
	}
	if book == nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get book with id %v", task.Id))
	}
	p.logger.Debug("Book retrieved successfully")

	// Getting s3key
	p.logger.Debug("Retrieving document key...")
	key, err := helpers.GetDocumentKey(book.File)
	if err != nil {
		return err
	}
	if key == nil {
		return errors.Wrap(err, "failed to get document key")
	}
	p.logger.Debug("Key retrieved successfully")

	// Getting link to download document
	p.logger.Debug("Retrieving document link from S3...")
	linkResponse, err := p.documenterConnector.GetDocumentLink(*key)
	if err != nil {
		return err
	}
	p.logger.Debug("Link retrieved successfully")

	// Downloading document
	p.logger.Debug("Downloading document...")
	rawDocument, err := helpers.DownloadDocument(linkResponse.Data.Attributes.Url)
	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}
	p.logger.Debug("Document downloaded successfully...")

	// Generating user signature
	p.logger.Debug("Generating signature...")
	reader := bytes.NewReader(rawDocument)
	pdfSignatureGenerator := pdf_signature_generator.New(p.signatureParams)
	raw, err := pdfSignatureGenerator.GenerateSignature(reader, task.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to generate signature")
	}
	p.logger.Debug("Signature generated successfully")

	// Calculating IPFS HASH
	p.logger.Debug("Calculating IPFS Hash...")
	// TODO: Precalculate IPFS HASH
	ipfsHash := "mocked12344"

	err = p.generatorDB.Tasks().UpdateIpfsHash(ipfsHash, task.Id)
	if err != nil {
		return errors.Wrap(err, "failed to update ipfs hash")
	}
	p.logger.Debug("IPFS Hash calculated successfully")

	// Uploading File
	p.logger.Debug("Uploading document to S3...")
	statusCode, err := p.documenterConnector.UploadDocument(raw, ipfsHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to upload file: status %v", statusCode))
	}
	p.logger.Debug("Document downloaded successfully")

	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)
	return nil
}

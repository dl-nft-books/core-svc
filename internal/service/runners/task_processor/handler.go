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
	book, err := p.booksDB.FilterActual().FilterByID(task.BookId).Get()
	if err != nil {
		return err
	}
	if book == nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get book with id %v", task.Id))
	}

	// Getting s3key
	key, err := helpers.GetDocumentKey(book.File)
	if err != nil {
		return err
	}
	if key == nil {
		return errors.Wrap(err, "failed to get document key")
	}

	// Getting link
	linkResponse, err := p.documenerConnector.GetDocumentLink(*key)
	if err != nil {
		return err
	}

	// Downloading pdf
	rawDocument, err := helpers.DownloadDocument(linkResponse.Data.Attributes.Url)
	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}

	reader := bytes.NewReader(rawDocument)

	// Generating user signature
	pdfSignatureGenerator := pdf_signature_generator.New(p.signatureParams)
	newReader, err := pdfSignatureGenerator.GenerateSignature(reader, task.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to generate signature")
	}

	// Calculating IPFS HASH
	// TODO: Precalculate IPFS HASH
	ipfsHash := "mocked"

	// Uploading File
	statusCode, err := p.documenerConnector.UploadDocument(newReader, ipfsHash)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to upload file: status %v", statusCode))
	}

	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)
	return nil
}

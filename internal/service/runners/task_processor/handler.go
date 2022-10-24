package task_processor

import (
	"os"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/pdf_signature_generator"
)

func (p *TaskProcessor) handleTask(task data.Task) error {
	p.logger.Debugf("Started processing task with id of %d", task.Id)
	// TODO: Get pdf file

	// MOCKED
	f, _ := os.Open("mocked")

	pdfSignatureGenerator := pdf_signature_generator.New(p.signatureParams)
	pdfSignatureGenerator.GenerateSignature(
		f, // MOCKED
		task.Signature,
	)

	// TODO: Precalculate IPFS HASH

	// TODO: Save pdf on S3
	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)
	return nil
}

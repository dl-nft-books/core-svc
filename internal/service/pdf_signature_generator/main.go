package pdf_signature_generator

import (
	"io"

	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
)

type PdfSignatureGenerator struct {
	signatureParams *config.SignatureParams
}

func New(sigParams *config.SignatureParams) *PdfSignatureGenerator {
	return &PdfSignatureGenerator{
		signatureParams: sigParams,
	}
}

func (g *PdfSignatureGenerator) GenerateSignature(document io.ReadSeeker, signature string) error {
	pdfReader, err := model.NewPdfReader(document)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	c := creator.New()

	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return err
		}

		err = c.AddPage(page)
		if err != nil {
			return err
		}

		// Setting signature only to first page
		if i == 0 {
			pageWidth, _, err := page.Size()
			if err != nil {
				return err
			}

			err = g.addSignature(c, pageWidth, signature)
			if err != nil {
				return err
			}
		}
	}

	// TODO: SAVE RESULT

	return nil
}

func (g *PdfSignatureGenerator) addSignature(cr *creator.Creator, pageWidth float64, signature string) error {
	upperLineYPos := g.signatureParams.MainHeightIndent - g.signatureParams.UpperLineHeightIdent
	upperLine := g.createLine(
		cr,
		g.signatureParams.LineWidthIndent,
		upperLineYPos,
		pageWidth-g.signatureParams.LineWidthIndent,
		upperLineYPos,
	)

	paragraph := g.createSignatureParagraph(
		cr,
		g.signatureParams.ParagraphWidthIndent,
		g.signatureParams.MainHeightIndent,
		pageWidth,
		signature,
	)

	lowerLineYPos := g.signatureParams.MainHeightIndent + paragraph.Height() + g.signatureParams.LowerLineHeightIdent
	lowerLine := g.createLine(
		cr,
		g.signatureParams.LineWidthIndent,
		lowerLineYPos,
		pageWidth-g.signatureParams.LineWidthIndent,
		lowerLineYPos,
	)

	err := cr.Draw(upperLine)
	if err != nil {
		return err
	}

	err = cr.Draw(paragraph)
	if err != nil {
		return err
	}

	return cr.Draw(lowerLine)
}

func (g *PdfSignatureGenerator) createLine(cr *creator.Creator, x1, y1, x2, y2 float64) *creator.Line {
	lineStyle := draw.LineStyleSolid
	lineColor := creator.ColorRGBFromHex("#BCBCBC")

	line := cr.NewLine(x1, y1, x2, y2)
	line.SetStyle(lineStyle)
	line.SetColor(lineColor)

	return line
}

func (g *PdfSignatureGenerator) createSignatureParagraph(cr *creator.Creator, xpos, ypos, pageWidth float64, signature string) *creator.Paragraph {
	paragraphFont, _ := model.NewStandard14Font("Times-Italic")
	paragraphColor := creator.ColorRGBFromHex("#2E2E2E")

	paragraph := cr.NewParagraph(signature)
	paragraph.SetWidth(pageWidth - xpos*2)
	// Setting position
	paragraph.SetPos(xpos, ypos)
	// Setting font
	paragraph.SetFont(paragraphFont)
	paragraph.SetFontSize(15)
	// Setting text aligment
	paragraph.SetTextAlignment(creator.TextAlignmentCenter)
	// Setting text color
	paragraph.SetColor(paragraphColor)

	return paragraph
}

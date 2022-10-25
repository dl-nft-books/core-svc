package pdf_signature_generator

import (
	"bytes"
	"io"

	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
)

const (
	baseFont       = "Times-Italic"
	paragraphColor = "#2E2E2E"
	lineColor      = "#BCBCBC"
)

type PdfSignatureGenerator struct {
	signatureParams *config.SignatureParams
}

func New(sigParams *config.SignatureParams) *PdfSignatureGenerator {
	return &PdfSignatureGenerator{
		signatureParams: sigParams,
	}
}

func (g *PdfSignatureGenerator) GenerateSignature(document io.ReadSeeker, signature string) (io.Reader, error) {
	pdfReader, err := model.NewPdfReader(document)
	if err != nil {
		return nil, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}

	c := creator.New()

	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return nil, err
		}

		err = c.AddPage(page)
		if err != nil {
			return nil, err
		}

		// Setting signature only to first page
		if i == 0 {
			pageWidth, _, err := page.Size()
			if err != nil {
				return nil, err
			}

			err = g.addSignature(c, pageWidth, signature)
			if err != nil {
				return nil, err
			}
		}
	}

	buff := bytes.Buffer{}
	err = c.Write(&buff)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buff.Bytes()), nil
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
	lineColor := creator.ColorRGBFromHex(lineColor)

	line := cr.NewLine(x1, y1, x2, y2)
	line.SetStyle(lineStyle)
	line.SetColor(lineColor)

	return line
}

func (g *PdfSignatureGenerator) createSignatureParagraph(cr *creator.Creator, xpos, ypos, pageWidth float64, signature string) *creator.Paragraph {
	paragraphFont, _ := model.NewStandard14Font(baseFont)
	paragraphColor := creator.ColorRGBFromHex(paragraphColor)

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

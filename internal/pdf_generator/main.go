package pdf_generator

import (
	"bytes"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"math/rand"

	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
)

const (
	baseFont        = "Times-Italic"
	fontSize        = 14
	paragraphColor  = "#2E2E2E"
	lineColor       = "#BCBCBC"
	maxBoxDeviation = 1.0
)

type SignatureGenerator struct {
	signatureParams *config.SignatureParams
}

func New(signatureParams *config.SignatureParams) *SignatureGenerator {
	return &SignatureGenerator{
		signatureParams: signatureParams,
	}
}

func (g *SignatureGenerator) GenerateSignature(document io.ReadSeeker, signature string) ([]byte, error) {
	pdfReader, err := model.NewPdfReader(document)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize a new reader")
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get number of pages")
	}

	pdfCreator := creator.New()

	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get page", logan.F{
				"page_number": i + 1,
			})
		}

		if err = pdfCreator.AddPage(page); err != nil {
			return nil, errors.Wrap(err, "failed to add page")
		}

		// Setting signature only to the first page
		if i == 0 {
			pageWidth, _, err := page.Size()
			if err != nil {
				return nil, errors.Wrap(err, "failed to get page's size")
			}

			err = g.addSignature(pdfCreator, pageWidth, signature)
			if err != nil {
				return nil, errors.Wrap(err, "failed to add signature to the page")
			}
		}
	}

	buff := bytes.Buffer{}
	err = pdfCreator.Write(&buff)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write output of creator")
	}

	return buff.Bytes(), nil
}

func (g *SignatureGenerator) addSignature(pdfCreator *creator.Creator, pageWidth float64, signature string) error {
	upperLineYPos := g.signatureParams.MainHeightIndent - g.signatureParams.UpperLineHeightIndent
	upperLine := g.createLine(
		pdfCreator,
		g.signatureParams.LineWidthIndent,
		upperLineYPos,
		pageWidth-g.signatureParams.LineWidthIndent,
		upperLineYPos,
	)

	paragraph := g.createSignatureParagraph(
		pdfCreator,
		g.signatureParams.ParagraphWidthIndent,
		g.signatureParams.MainHeightIndent,
		pageWidth,
		signature,
	)

	lowerLineYPos := g.signatureParams.MainHeightIndent + paragraph.Height() + g.signatureParams.LowerLineHeightIndent
	lowerLine := g.createLine(
		pdfCreator,
		g.signatureParams.LineWidthIndent,
		lowerLineYPos,
		pageWidth-g.signatureParams.LineWidthIndent,
		lowerLineYPos,
	)

	if err := pdfCreator.Draw(upperLine); err != nil {
		return errors.Wrap(err, "failed to draw a line")
	}

	if err := pdfCreator.Draw(paragraph); err != nil {
		return errors.Wrap(err, "failed to draw a paragraph")
	}

	return pdfCreator.Draw(lowerLine)
}

func (g *SignatureGenerator) createLine(cr *creator.Creator, x1, y1, x2, y2 float64) *creator.Line {
	var (
		lineStyle    = draw.LineStyleSolid
		lineColorRGB = creator.ColorRGBFromHex(lineColor)
	)

	line := cr.NewLine(x1, y1, x2, y2)
	line.SetStyle(lineStyle)
	line.SetColor(lineColorRGB)

	return line
}

func (g *SignatureGenerator) createSignatureParagraph(cr *creator.Creator, xPosition, yPosition, pageWidth float64, signature string) *creator.Paragraph {
	paragraphFont, _ := model.NewStandard14Font(baseFont)
	paragraphColorRGB := creator.ColorRGBFromHex(paragraphColor)

	// pseudo-randomization to avoid ipfs hash collision
	xPosition += randomDeviation(maxBoxDeviation)

	paragraph := cr.NewParagraph(signature)
	paragraph.SetWidth(pageWidth - xPosition*2)
	// Setting position
	paragraph.SetPos(xPosition, yPosition)
	// Setting font
	paragraph.SetFont(paragraphFont)
	paragraph.SetFontSize(fontSize)
	// Setting text alignment
	paragraph.SetTextAlignment(creator.TextAlignmentCenter)
	// Setting text color
	paragraph.SetColor(paragraphColorRGB)

	return paragraph
}

func randomDeviation(maxDeviation float64) float64 {
	var dx float64

	// Each time in the loop we will have different values
	for i := 0; i < 3; i++ {
		dx += maxDeviation * float64(rand.Int31n(100)) / 300 // max is maxDeviation/3
	}

	// Max offset is maxOffset which is still barely noticeable for sufficiently small values
	return dx
}

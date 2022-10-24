package config

import (
	"github.com/unidoc/unipdf/v3/common/license"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type PdfSignatureConfig struct {
	ApiKey          string          `json:"api_key"`
	SignatureParams SignatureParams `json:"params"`
}

type SignatureParams struct {
	MainHeightIndent     float64 `json:"main_height_indent"`
	LineWidthIndent      float64 `json:"line_width_indent"`
	ParagraphWidthIndent float64 `json:"paragraph_width_indent"`
	UpperLineHeightIdent float64 `json:"upper_line_height_ident"`
	LowerLineHeightIdent float64 `json:"lower_line_height_Ident"`
}

func (c *config) PdfSignatureParams() *SignatureParams {
	return c.pdfSignatureParams.Do(func() interface{} {
		var cfg PdfSignatureConfig

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "pdf_signature_params")).
			Please()
		if err != nil {
			panic(err)
		}

		err = license.SetMeteredKey(cfg.ApiKey)
		if err != nil {
			panic(err)
		}

		return &cfg.SignatureParams
	}).(*SignatureParams)
}

package config

import (
	"github.com/unidoc/unipdf/v3/common/license"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type PdfSignatureConfig struct {
	ApiKey          string          `fig:"api_key,required"`
	SignatureParams SignatureParams `fig:"params,required"`
}

type SignatureParams struct {
	MainHeightIndent      float64 `fig:"main_height_indent,required"`
	LineWidthIndent       float64 `fig:"line_width_indent,required"`
	ParagraphWidthIndent  float64 `fig:"paragraph_width_indent,required"`
	UpperLineHeightIndent float64 `fig:"upper_line_height_indent,required"`
	LowerLineHeightIndent float64 `fig:"lower_line_height_indent,required"`
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

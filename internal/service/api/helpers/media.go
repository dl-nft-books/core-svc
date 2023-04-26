package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func MarshalMedia(media ...*resources.Media) []string {
	var res []string

	for _, v := range media {
		raw, err := json.Marshal(v)
		if err != nil {
			return nil
		}

		res = append(res, string(raw))
	}

	return res
}

func UnmarshalMedia(media ...string) ([]resources.Media, error) {
	var res []resources.Media
	var unmarshalledMedia *resources.Media

	for _, value := range media {
		err := json.Unmarshal([]byte(value), &unmarshalledMedia)
		if err != nil {
			return nil, err
		}

		res = append(res, *unmarshalledMedia)
		unmarshalledMedia = nil

	}
	return res, nil
}

func SetMediaLink(r *http.Request, media *resources.Media) error {
	dconnector := DocumenterConnector(r)

	link, err := dconnector.GetDocumentLink(media.Attributes.Key)
	if err != nil {
		return err
	}

	media.Attributes.Url = &link.Data.Attributes.Url

	return nil
}

func CheckBannerMimeType(ext string, r *http.Request) error {
	for _, el := range MimeTypes(r).AllowedBannerMimeTypes {
		if el == ext {
			return nil
		}
	}
	return errors.New("invalid banner extension")
}

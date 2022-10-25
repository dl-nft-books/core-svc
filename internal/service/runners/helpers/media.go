package helpers

import (
	"encoding/json"
)

func GetDocumentKey(rawMedia string) (*string, error) {
	media, err := UnmarshalMedia(rawMedia)
	if err != nil {
		return nil, err
	}

	return &media.Attributes.Key, nil
}

func UnmarshalMedia(media string) (Media, error) {
	var res Media

	err := json.Unmarshal([]byte(media), &res)
	if err != nil {
		return Media{}, err
	}

	return res, nil
}

type Media struct {
	Attributes MediaAttributes `json:"attributes"`
}

type MediaAttributes struct {
	// S3 key
	Key string `json:"key"`
	// Media mime type
	MimeType string `json:"mime_type"`
	// Media title
	Name string `json:"name"`
}

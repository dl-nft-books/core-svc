package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *Connector) UploadDocument(raw []byte, key string) (int, error) {
	// forming endpoint
	endpoint := fmt.Sprintf("%s/%s", c.baseUrl, DocumentEndpoint)

	// forming multipart/form-data request : setting headers
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="Document"; filename="document.pdf"`))
	h.Set("Content-Type", "application/pdf")

	// forming multipart/form-data request : adding file
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreatePart(h)
	if err != nil {
		return http.StatusBadRequest, err
	}

	_, err = part.Write(raw)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// forming multipart/form-data request : adding `Key` field
	fw, err := writer.CreateFormField("Key")
	if err != nil {
		return http.StatusBadRequest, err
	}
	_, err = io.Copy(fw, strings.NewReader(key))
	if err != nil {
		return http.StatusBadRequest, err
	}

	writer.Close()

	// creating request
	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return http.StatusBadRequest, errors.Wrap(err, "failed to build request")
	}

	// setting headers
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// sending request
	resp, err := c.client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return http.StatusBadRequest, err
	}

	return resp.StatusCode, nil
}

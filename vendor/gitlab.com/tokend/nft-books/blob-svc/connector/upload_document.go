package connector

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *Connector) UploadDocument(file io.Reader, key string) (int, error) {
	parsedUrl, err := url.Parse(DocumentEndpoint)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse document url")
	}

	fullEndpoint, err := c.client.Resolve(parsedUrl)
	if err != nil {
		return 0, err
	}

	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="Document"; filename="1.pdf"`))
	h.Set("Content-Type", "application/pdf")

	part, err := mp.CreatePart(h)
	if err != nil {
		return http.StatusBadRequest, err
	}

	bd, err := ioutil.ReadAll(file)
	if err != nil {
		return http.StatusBadRequest, err
	}

	part.Write(bd)
	mp.Close()

	req, err := http.NewRequest(http.MethodPost, fullEndpoint, body)
	req.Header.Add("Content-Type", mp.FormDataContentType())
	req.Form.Set("Key", key)

	resp, err := c.client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

package ga4datatest

import (
	"bytes"
	"io"
	"net/http"
)

type customTransport struct {
	body      []byte
	transport http.RoundTripper
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return res, err
	}
	res.Status = "200 OK"
	res.StatusCode = 200
	res.Body = io.NopCloser(bytes.NewReader(t.body))

	return res, nil
}

// NewTestClient :
func NewTestClient(body []byte) *http.Client {
	tr := &customTransport{
		body:      body,
		transport: http.DefaultTransport,
	}
	return &http.Client{
		Transport: tr,
	}
}

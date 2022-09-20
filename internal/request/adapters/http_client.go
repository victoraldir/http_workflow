package adapters

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/victoraldir/http-follower/internal/request/core/domain"
)

type (
	// HTTPClient is an adapter that implements the Client interface
	// and uses the standard library http.Client to make requests.
	httpClient struct {
		client *http.Client
	}
)

// NewHTTPClient returns a new instance of HTTPClient.
func NewHTTPClient() *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

// Do implements the Client interface.
func (c *httpClient) Do(req *domain.Request) (*domain.Response, error) {

	httpReq, err := http.NewRequest(req.Method, req.URL, bytes.NewBuffer([]byte(req.Body)))
	if err != nil {
		return nil, err
	}

	if req.Headers != nil {
		for key, element := range req.Headers {
			if strings.ToLower(key) == "host" {
				httpReq.Host = element
			} else {
				httpReq.Header.Set(key, element)
			}
		}
	}

	resp, err := c.client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &domain.Response{
		StatusCode: resp.StatusCode,
		Body:       string(bodyBytes),
	}, nil
}

package bitly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type client struct {
	ApiKey     string
	userAgent  string
	httpClient *http.Client
}

func (c *client) do(method, path string, body interface{}) (*http.Response, error) {
	u := formatUrl(path)

	buf, err := json.Marshal(body)

	if err != nil {
		return nil, ErrEncodingError
	}

	req, err := http.NewRequest(method, u, bytes.NewBuffer(buf))

	if err != nil {
		return nil, ErrRequestError
	}

	// set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	return c.httpClient.Do(req)
}

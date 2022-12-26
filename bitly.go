package bitly

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bitly struct {
	Client client
}

func New(token string) *Bitly {
	client := client{
		ApiKey:     token,
		userAgent:  "go-shorten/Bitly",
		httpClient: http.DefaultClient,
	}

	return &Bitly{
		client,
	}
}

// S interface implementation to be used with shorten aggregator. Use other functions instead for provider specific functionality.
func (b *Bitly) S(url string) (string, error) {
	res, err := b.ShortenBitLink(&ShortenUrlRequest{
		LongUrl: url,
	})

	if err != nil {
		return "", err
	}

	return res.Link, nil
}

// ShortenBitLink - /v4/shorten Shorten Url.
func (b *Bitly) ShortenBitLink(u *ShortenUrlRequest) (*ShortenUrlResponse, error) {
	resp, err := b.Client.do(http.MethodPost, "shorten", u)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var res ShortenUrlResponse

	err = json.Unmarshal(body, &res)

	if err != nil {
		fmt.Println("yehan hai error")
		fmt.Println(string(body))
		return nil, err
	}

	return &res, err
}

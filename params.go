package bitly

import "fmt"

const (
	baseUrl = "https://api-ssl.bitly.com"
	version = "v4"
)

type ShortenUrlRequest struct {
	LongUrl   string `json:"long_url,omitempty"`
	Domain    string `json:"domain,omitempty"`
	GroupGuid string `json:"group_guid,omitempty"`
}

type ShortenUrlResponse struct {
	Link           string   `json:"link"`
	Id             string   `json:"id"`
	LongUrl        string   `json:"long_url"`
	Archived       bool     `json:"archived"`
	CreatedAt      string   `json:"created_at"`
	CustomBitlinks []string `json:"custom_bitlinks"`
	Tags           []string `json:"tags"`
	Deeplinks      []struct {
		Guid        string `json:"guid"`
		Bitlink     string `json:"bitlink"`
		AppUriPath  string `json:"app_uri_path"`
		InstallUrl  string `json:"install_url"`
		AppGuid     string `json:"app_guid"`
		Os          string `json:"os"`
		InstallType string `json:"install_type"`
		Created     string `json:"created"`
		Modified    string `json:"modified"`
		BrandGuid   string `json:"brand_guid"`
	} `json:"deeplinks"`
}

func formatUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", baseUrl, version, endpoint)
}

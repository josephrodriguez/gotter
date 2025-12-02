package http

import (
	"fmt"
	"net/http"
	"time"
)

type WebhookResponse struct {
	Id       string           `json:"webhook_id"`
	Name     string           `json:"id"`
	URL      string           `json:"url"`
	Active   bool             `json:"active"`
	Packages []WebhookPackage `json:"packages"`
}

type WebhookPackage struct {
	Id             string            `json:"package_id"`
	Name           string            `json:"name"`
	NormalizedName string            `json:"normalized_name"`
	Description    string            `json:"description"`
	Version        string            `json:"version"`
	AppVersion     string            `json:"app_version"`
	LogoImageId    string            `json:"logo_image_id"`
	License        string            `json:"license"`
	Deprecated     bool              `json:"deprecated"`
	Signed         bool              `json:"signed"`
	Stars          int               `json:"stars"`
	Repository     WebhookRepository `json:"repository"`
}

type WebhookRepository struct {
	Id                      string `json:"repository_id"`
	Name                    string `json:"name"`
	URL                     string `json:"url"`
	Private                 bool   `json:"private"`
	Kind                    int    `json:"kind"`
	VerifiedPublisher       bool   `json:"verified_publisher"`
	Official                bool   `json:"official"`
	ScannerDisabled         bool   `json:"scanner_disabled"`
	OrganizationName        string `json:"organization_name"`
	OrganizationDisplayName string `json:"organization_display_name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	ApiKey     string
	ApiSecret  string
}

func NewClient(baseUrl, apiKey, apiSecret string) Client {
	return Client{
		BaseURL:   baseUrl,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", c.ApiKey)
	req.Header.Set("X-API-SECRET", c.ApiSecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

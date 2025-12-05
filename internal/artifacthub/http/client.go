package http

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) handleResponse(resp *http.Response, out interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return fmt.Errorf("request failed with status %d:%s", resp.StatusCode, errResp.Message)
		}
		return fmt.Errorf("request failed with status %d:%s", resp.StatusCode, string(body))
	}

	if len(body) > 0 && out != nil {
		if err := json.Unmarshal(body, &out); err != nil {
			return fmt.Errorf("failed to unmarshal response body: %w", err)
		}
	}

	return nil
}

func (c *Client) getWebhooks(path string) ([]WebhookResponse, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var out []WebhookResponse
	if err := c.handleResponse(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) getWebhook(path string) (WebhookResponse, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return WebhookResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return WebhookResponse{}, err
	}

	var out WebhookResponse
	if err := c.handleResponse(resp, &out); err != nil {
		return WebhookResponse{}, err
	}

	return out, nil
}

func (c *Client) GetUserWebhooks() ([]WebhookResponse, error) {
	return c.getWebhooks("api/v1/webhooks/user")
}

func (c *Client) GetOrganizationWebhooks(name string) ([]WebhookResponse, error) {
	path := fmt.Sprintf("api/v1/webhooks/org/%s", name)
	return c.getWebhooks(path)
}

func (c *Client) GetUserWebhook(webhookId string) (WebhookResponse, error) {
	path := fmt.Sprintf("/api/v1/webhooks/user/%s", webhookId)
	return c.getWebhook(path)
}

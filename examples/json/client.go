package json

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AntonLuning/tiny-url/models"
)

const (
	URL_API_PATH = "/api/v1/url"
)

type Client struct {
	endpointAddr string
}

func New(endpointAddr string) *Client {
	return &Client{
		endpointAddr: endpointAddr,
	}
}

func (c *Client) CreateShortenURL(originalURL string) (*models.CreateShortenURLResponse, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s?original=%s", c.endpointAddr, URL_API_PATH, originalURL), nil)
	if err != nil {
		return nil, err
	}

	httpResp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	resp := models.CreateShortenURLResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetOriginalURL(shortenURL string) (*models.GetOriginalURLResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s%s?shorten=%s", c.endpointAddr, URL_API_PATH, shortenURL), nil)
	if err != nil {
		return nil, err
	}

	httpResp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	resp := models.GetOriginalURLResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func executeRequest(r *http.Request) (*http.Response, error) {
	httpResp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		errMsg := map[string]any{}
		if err := json.NewDecoder(httpResp.Body).Decode(&errMsg); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with non OK status code: \"%d\" msg: \"%s\"", httpResp.StatusCode, errMsg)
	}

	return httpResp, nil
}

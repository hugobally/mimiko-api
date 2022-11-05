package spotify

import (
	"github.com/hugobally/mimiko/backend/config"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Config     *config.Config
}

func New(cfg *config.Config, client *http.Client) *Client {
	return &Client{
		HttpClient: client,
		Config:     cfg,
	}
}

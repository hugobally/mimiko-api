package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`
}

func (c *Client) CreateClientCredentialsToken() (*TokenResponse, error) {
	v := url.Values{}
	v.Add("grant_type", "client_credentials")

	return c.performTokenRequest(v)
}

func (c *Client) CreateAuthCodeToken(authCode string) (*TokenResponse, error) {
	v := url.Values{}
	v.Add("grant_type", "authorization_code")
	v.Add("code", authCode)
	v.Add("redirect_uri", c.Config.Spotify.RedirectUri)

	return c.performTokenRequest(v)
}

func (c *Client) RefreshToken(token string) (*TokenResponse, error) {
	v := url.Values{}
	v.Add("grant_type", "refresh_token")
	v.Add("refresh_token", token)

	return c.performTokenRequest(v)
}

func (c *Client) performTokenRequest(content url.Values) (*TokenResponse, error) {
	req, err := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token",
		bytes.NewBuffer([]byte(content.Encode())))
	if err != nil {
		return nil, errors.New("internal server error")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	credentials := base64.URLEncoding.EncodeToString(
		[]byte(c.Config.Spotify.ClientId + ":" + c.Config.Spotify.ClientSecret))

	req.Header.Set("Authorization", "Basic "+credentials)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	token := &TokenResponse{}
	_ = json.Unmarshal(body, token)

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("%v", token.Error)
		return nil, errors.New(msg)
	}

	return token, nil
}

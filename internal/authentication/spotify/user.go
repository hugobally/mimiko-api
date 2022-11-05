package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserResponse struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Error       struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *Client) GetUser(token string) (*UserResponse, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := &UserResponse{}
	_ = json.Unmarshal(body, &user)
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("error during spotify authentication : %v %v", user.Error.Status, user.Error.Message)
		return nil, errors.New(msg)
	}

	return user, nil
}

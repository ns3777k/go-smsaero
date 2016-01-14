package smsaero

import (
	"net/http"
	"strings"
	"net/url"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	baseURL = "https://gate.smsaero.ru"
)

type ErrorResponse struct {
	Result string `json:"result,omitempty"`
	Reason string `json:"reason,omitempty"`
}

func (e ErrorResponse) IsErrorResponse() bool {
	return (len(e.Result) > 0 && e.Result == "reject") || len(e.Reason) > 0
}

func (e ErrorResponse) GetError() error {
	return errors.New(fmt.Sprintf("Result: %s. Reason: %s", e.Result, e.Reason))
}

type ErrorableResponse interface {
	IsErrorResponse() bool
	GetError() error
}

type Client struct {
	username string
	password string
	client   *http.Client
}

func (c *Client) executeRequest(path string, destination ErrorableResponse, params url.Values) error {
	fullURL := baseURL + path

	if params == nil {
		params = make(url.Values, 0)
	}

	params.Add("user", c.username)
	params.Add("password", c.password)
	params.Add("answer", "json")

	resp, err := c.client.Post(fullURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(destination); err != nil {
		return err
	}

	return nil
}

func NewClient(httpClient *http.Client, username, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		username: username,
		password: password,
		client:   httpClient,
	}
}

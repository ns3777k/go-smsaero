package smsaero

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://gate.smsaero.ru"
)

// ErrorResponse описывает поля при ответе с ошибкой
type ErrorResponse struct {
	Result string `json:"result,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// IsErrorResponse проверяет ответ на наличие ошибок
func (e ErrorResponse) IsErrorResponse() bool {
	return (len(e.Result) > 0 && e.Result == "reject") || len(e.Reason) > 0
}

// GetError возвращает строковое представление ошибки в ответе
func (e ErrorResponse) GetError() error {
	return fmt.Errorf("Result: %s. Reason: %s", e.Result, e.Reason)
}

// ErrorableResponse определяет интерфейс ответа, который может содержать ошибку
type ErrorableResponse interface {
	IsErrorResponse() bool
	GetError() error
}

// Client определяем структуру клиента
type Client struct {
	username string
	password string
	client   *http.Client
}

// executeRequest отсылает POST-запрос сервису и декодирует JSON-ответ в
// структуру реализующую ErrorableResponse
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

// NewClient создает новый клиент и возвращает его
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

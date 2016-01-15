package smsaero

import (
	"net/url"
	"strconv"
)

const (
	sendPath   = "/send/"
	statusPath = "/status/"
)

// SendReport описывает ответ API по отправке SMS
type SendReport struct {
	ID int `json:"id,omitempty"`
	ErrorResponse
}

// StatusReport описывает ответ API по запросу статуса отправки
type StatusReport struct {
	ID int `json:"id,omitempty"`
	ErrorResponse
}

// Send отправляет SMS
// TODO: добавить опциональные параметры - дата отправки и цифровой канал
func (c *Client) Send(to int, text string, from string) (int, error) {
	sendReportResponse := new(SendReport)
	params := make(url.Values, 0)
	params.Add("to", strconv.Itoa(to))
	params.Add("text", text)
	params.Add("from", from)

	if err := c.executeRequest(sendPath, sendReportResponse, params); err != nil {
		return 0, err
	}

	if sendReportResponse.IsErrorResponse() {
		return 0, sendReportResponse.GetError()
	}

	return sendReportResponse.ID, nil
}

// GetStatus запрашивает статус отправки SMS
func (c *Client) GetStatus(smsID int) (string, error) {
	statusReportResponse := new(StatusReport)
	params := make(url.Values, 0)
	params.Add("id", strconv.Itoa(smsID))

	if err := c.executeRequest(statusPath, statusReportResponse, params); err != nil {
		return "", err
	}

	if statusReportResponse.IsErrorResponse() {
		return "", statusReportResponse.GetError()
	}

	return statusReportResponse.Result, nil
}

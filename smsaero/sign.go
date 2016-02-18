package smsaero

import (
	"encoding/json"
	"net/url"
)

const (
	signsListPath = "/senders/"
	signAddPath   = "/sign/"
)

// SignsList описывает ответ API на получение списка подписей
type SignsList struct {
	Senders []string
	ErrorResponse
}

// SignStatus описывает ответ API на запрос добавления/получение информации
// по статусу подписи
type SignStatus struct {
	Accepted string `json:"accepted,omitempty"`
	ErrorResponse
}

// UnmarshalJSON является необходимым для реализации интерфейса Unmarshaler,
// поскольку JSON-ответ может быть либо массив либо объект
func (s *SignsList) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &s.Senders); err == nil {
		return nil
	}

	if err := json.Unmarshal(b, &s.ErrorResponse); err != nil {
		return err
	}

	return nil
}

// GetSigns запрашивает список подписей отправителя
func (c *Client) GetSigns() ([]string, error) {
	signListResponse := new(SignsList)

	if err := c.executeRequest(signsListPath, signListResponse, nil); err != nil {
		return signListResponse.Senders, err
	}

	if signListResponse.IsErrorResponse() {
		return signListResponse.Senders, signListResponse.GetError()
	}

	return signListResponse.Senders, nil
}

// AddSign добавляет подпись отправителя или если такая подпись уже есть,
// возвращает статус ее модерации
func (c *Client) AddSign(sign string) (string, error) {
	signStatusResponse := new(SignStatus)
	params := make(url.Values, 0)
	params.Add("sign", sign)

	if err := c.executeRequest(signAddPath, signStatusResponse, params); err != nil {
		return signStatusResponse.Accepted, err
	}

	if signStatusResponse.IsErrorResponse() {
		return signStatusResponse.Accepted, signStatusResponse.GetError()
	}

	return signStatusResponse.Accepted, nil
}

package smsaero

import (
	"encoding/json"
	"net/url"
)

const (
	singsListPath = "/senders/"
	signAddPath = "/sign/"
)

type SignsList struct {
	Senders []string
	ErrorResponse
}

type SignStatus struct {
	Accepted string `json:"accepted,omitempty"`
	ErrorResponse
}

func (s *SignsList) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &s.Senders); err == nil {
		return nil
	}

	if err := json.Unmarshal(b, &s.ErrorResponse); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetSigns() ([]string, error) {
	signListResponse := new(SignsList)

	if err := c.executeRequest(singsListPath, signListResponse, nil); err != nil {
		return signListResponse.Senders, err
	}

	if signListResponse.IsErrorResponse() {
		return signListResponse.Senders, signListResponse.GetError()
	}

	return signListResponse.Senders, nil
}

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

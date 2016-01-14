package smsaero

const (
	balancePath = "/balance/"
)

type Balance struct {
	Balance float64 `json:"balance,string,omitempty"`
	ErrorResponse
}

func (c *Client) CheckBalance() (float64, error) {
	balanceResponse := new(Balance)

	if err := c.executeRequest(balancePath, balanceResponse, nil); err != nil {
		return 0, err
	}

	if balanceResponse.IsErrorResponse() {
		return 0, balanceResponse.GetError()
	}

	return balanceResponse.Balance, nil
}

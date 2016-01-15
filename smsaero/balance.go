package smsaero

const (
	balancePath = "/balance/"
)

// Balance описывает ответ API по запросу баланса
type Balance struct {
	Balance float64 `json:"balance,string,omitempty"`
	ErrorResponse
}

// CheckBalance запрашивает баланс аккаунта
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

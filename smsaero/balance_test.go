package smsaero

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_CheckBalance(t *testing.T) {
	setUpTestServe()
	defer tearDownTestServe()

	mux.HandleFunc(balancePath, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		assertRequiredFields(t, r.Form)
		assert.Equal(t, "POST", r.Method)
		w.Write(getStub(t, "balance"))
	})

	balance, err := client.CheckBalance()
	assert.Nil(t, err)
	assert.Equal(t, 42.85, balance)
}

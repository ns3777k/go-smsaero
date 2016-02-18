package smsaero

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetSigns(t *testing.T) {
	setUpTestServe()
	defer tearDownTestServe()

	mux.HandleFunc(signsListPath, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		assertRequiredFields(t, r.Form)
		assert.Equal(t, "POST", r.Method)
		w.Write(getStub(t, "signs"))
	})

	signs, err := client.GetSigns()
	assert.Nil(t, err)

	expected := []string{"NEWS", "TEST"}
	assert.Equal(t, expected, signs)
}

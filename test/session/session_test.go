package session

import (
	"bytes"
	"net/http"
	"testing"

	. "github.com/MoonSHRD/matrix-api-tests/configs"
)

func TestLogin(t *testing.T) {
	requestBody := []byte(`{"identifier": {"type": "m.id.user","user": "test"},"password": "testing","type": "m.login.password"}`)

	req, err := http.NewRequest("POST", Config.BaseURI+"/_matrix/client/r0/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

package user

import (
	"bytes"
	"net/http"
	"testing"

	. "github.com/MoonSHRD/matrix-api-tests/configs"
)

func TestRegister(t *testing.T) {
	requestBody := []byte(`{"username": "test", "password": "testing", "auth": {"type": "m.login.dummy"}}`)

	req, err := http.NewRequest("POST", Config.BaseURI+"/_matrix/client/r0/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("kind", "user")

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		// Get 400 status code if user already exist
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRegisterAvailable(t *testing.T) {
	req, err := http.NewRequest("GET", Config.BaseURI+"/_matrix/client/r0/register/available", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	query := req.URL.Query()
	query.Add("username", "test")

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		// Get 400 status code if user already exist
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

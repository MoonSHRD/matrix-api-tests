package session

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/MoonSHRD/matrix-api-tests/configs"
	. "github.com/MoonSHRD/matrix-api-tests/test/user"
)

var (
	user User
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
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	if err = fillUserData(res.Body, &user); err != nil {
		t.Fatal(err.Error())
	}
}

func fillUserData(body io.ReadCloser, user *User) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	defer body.Close()

	if err = json.Unmarshal(data, user); err != nil {
		return err
	}

	return nil
}

func TestLogout(t *testing.T) {
	req, err := http.NewRequest("POST", Config.BaseURI+"/_matrix/client/r0/logout", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	query := req.URL.Query()
	query.Add("access_token", user.AccessToken)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

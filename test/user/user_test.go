package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"

	. "github.com/MoonSHRD/matrix-api-tests/configs"
)

var (
	user *User
)

func TestRegister(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	requestBody := []byte(fmt.Sprintf(`{"username": "test%v", "password": "testing", "auth": {"type": "m.login.dummy"}}`, rand.Int()))

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
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err.Error())
		}

		// Checks if user already exists
		expected := []byte(`{"errcode": "M_USER_IN_USE","error": "User ID already taken."}`)
		if !bytes.Equal(resBody, expected) {
			t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
		}
	}

	if err = json.NewDecoder(res.Body).Decode(user); err != nil {
		t.Fatal(err.Error())
	}
}

func TestRegisterAvailable(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	req, err := http.NewRequest("GET", Config.BaseURI+"/_matrix/client/r0/register/available", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	query := req.URL.Query()
	query.Add("username", fmt.Sprintf("test%v", rand.Int()))

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err.Error())
		}

		// Checks if user already exists
		expected := []byte(`{"errcode": "M_USER_IN_USE","error": "User ID already taken."}`)
		if !bytes.Equal(resBody, expected) {
			t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
		}
	}
}

func TestWhoAmI(t *testing.T) {
	req, err := http.NewRequest("GET", Config.BaseURI+"/_matrix/client/r0/account/whoami", nil)
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

func TestGetUserInfoByID(t *testing.T) {
	req, err := http.NewRequest("GET", Config.BaseURI+"/_matrix/client/r0/profile/"+user.UserID, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGet3PID(t *testing.T) {
	req, err := http.NewRequest("GET", Config.BaseURI+"/_matrix/client/r0/account/3pid", nil)
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

package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"rest-get-api/pkg/model"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pong)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `pong`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedUsers := model.Users

	r := []model.User{}
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Error(err, "ioutil.ReadAll")
		t.FailNow()
	}
	if err := json.Unmarshal(body, &r); err != nil {
		t.Error(err, "json.Unmarshal")
		t.FailNow()
	}
	responseUsers := map[string]model.User{}
	for i := range r {
		responseUsers[r[i].Name] = r[i]
	}

	if !reflect.DeepEqual(expectedUsers, responseUsers) {
		t.Error("expected", expectedUsers)
		t.Error("response", responseUsers)
		t.FailNow()
	}
}

func TestGetKnownUser(t *testing.T) {
	const username = "John"
	req, err := http.NewRequest("GET", "/users?name="+username, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedUser := model.Users[username]

	r := []model.User{}
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Error(err, "ioutil.ReadAll")
		t.FailNow()
	}
	if err := json.Unmarshal(body, &r); err != nil {
		t.Error(err, "json.Unmarshal")
		t.FailNow()
	}
	responseUsers := map[string]model.User{}
	for i := range r {
		responseUsers[r[i].Name] = r[i]
	}

	responseUser, ok := responseUsers[username]
	if !ok {
		t.Error("user not found", username)
		t.Error("response", responseUser)
		t.FailNow()
	}

	if !reflect.DeepEqual(expectedUser, responseUser) {
		t.Error("expected", expectedUser)
		t.Error("response", responseUsers)
		t.FailNow()
	}
}

func TestGetUnknownUser(t *testing.T) {
	const username = "Doe"
	req, err := http.NewRequest("GET", "/users?name="+username, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Error(err, "ioutil.ReadAll")
		t.FailNow()
	}

	responseUser := []model.User{}
	if err := json.Unmarshal(body, &responseUser); err != nil {
		t.Error(err, "json.Unmarshal")
		t.FailNow()
	}

	if len(responseUser) != 0 {
		t.Error("expected no users to be found")
		t.Error("response", responseUser)
		t.FailNow()
	}
}

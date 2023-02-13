package route

import (
	"bytes"
	"car-rental/internal/model"
	"car-rental/utils"
	"net/http"
	"net/http/httptest"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

var (
	testUser = &model.User{
		Name:    "testName",
		Contact: "testContact",
	}
	byteUser = []byte{}
)

func (ts *testSuite) Test10AddUser() {
	var err error
	byteUser, err = json.Marshal(testUser)
	assert.NoError(ts.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(byteUser))
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	outputUser := &model.User{}
	err = json.Unmarshal(rec.Body.Bytes(), outputUser)
	assert.NoError(ts.T(), err)
	assert.False(ts.T(), outputUser.UUID.IsNil())
	assert.Equal(ts.T(), testUser.Contact, outputUser.Contact)
	assert.Equal(ts.T(), testUser.Name, outputUser.Name)
	testUser = outputUser
}

func (ts *testSuite) Test20GetUser() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/"+testUser.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	outputUser := model.User{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputUser)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), outputUser, *testUser)
}

func (ts *testSuite) Test30GetUsers() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	outputUsers := []model.User{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputUsers)
	assert.NoError(ts.T(), err)
	assert.True(ts.T(), len(outputUsers) > 0)
	found := false
find:
	for i := range outputUsers {
		if outputUsers[i].UUID == testUser.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)
}

func (ts *testSuite) Test40UpdateUser() {
	outputUser := model.User{
		Name:    "updateName",
		Contact: "updateContact",
	}
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/"+testUser.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	err = json.Unmarshal(rec.Body.Bytes(), &outputUser)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputUser, *testUser)
}

func (ts *testSuite) Test99DeleteUser() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/"+testUser.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/users/"+testUser.UUID.String(), nil)
	rec = httptest.NewRecorder()
	assert.NoError(ts.T(), err)
	ts.Serve(req, rec)

	outputMsg := utils.APIError{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputMsg)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), outputMsg.ErrorCode, http.StatusNotFound)
	assert.Equal(ts.T(), outputMsg.ErrorMessage, "record not found")
}

package repository

import (
	"car-rental/internal/model"

	"github.com/stretchr/testify/assert"
)

var (
	originalLenUsers = 0

	testUser = &model.User{}
)

func (ts testSuite) Test20CreateUser() {
	users, err := ts.GetAllUsers()
	assert.NoError(ts.T(), err)
	originalLenUsers = len(users)
	newUser := &model.User{
		Name:        "testName",
		Contact:     "testContact",
		PhoneNumber: "testPhoneNumber",
	}
	testUser, err = ts.AddUser(newUser)
	assert.NotNil(ts.T(), testUser.UUID.IsNil())
	assert.NoError(ts.T(), err)
}

func (ts testSuite) Test3GetUser() {
	User, err := ts.GetUserByUUID(testUser.UUID)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), User, testUser)
}

func (ts testSuite) Test4GetAllUser() {
	users, err := ts.GetAllUsers()
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), len(users), originalLenUsers+1)
	found := false
find:
	for i := range users {
		if users[i].UUID == testUser.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)

}

func (ts testSuite) Test5UpdateUser() {

	newUser := testUser.Copy()
	newUser.Contact = "newContact"
	err := ts.UpdateUser(newUser, newUser.UUID)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), newUser.Contact, testUser.Contact)
}

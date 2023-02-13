package repository

import (
	"car-rental/internal/model"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (tc testSuite) GetUserByUUID(uuid uuid.UUID) (*model.User, error) {
	user, ok := mockUsers[uuid]
	if ok {
		return user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (tc testSuite) GetAllUsers() ([]*model.User, error) {
	output := []*model.User{}
	for key := range mockUsers {
		output = append(output, mockUsers[key])
	}
	return output, nil
}

func (tc testSuite) DeleteUser(uuid uuid.UUID) error {
	_, ok := mockUsers[uuid]
	if ok {
		delete(mockUsers, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (tc testSuite) AddUser(user *model.User) (*model.User, error) {
	id, _ := uuid.NewV4()
	user.UUID = id
	mockUsers[user.UUID] = user
	return user, nil
}

func (tc testSuite) UpdateUser(user *model.User, uuid uuid.UUID) error {
	_, ok := mockUsers[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	user.UUID = uuid
	mockUsers[uuid] = user
	return nil
}

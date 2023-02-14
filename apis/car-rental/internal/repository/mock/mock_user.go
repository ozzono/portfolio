package mock

import (
	"car-rental/internal/model"
	"car-rental/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (m Repo) GetUserByUUID(uuid utils.UUID) (*model.User, error) {
	user, ok := mockUsers[uuid]
	if ok {
		return user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m Repo) GetAllUsers() ([]*model.User, error) {
	output := []*model.User{}
	for key := range mockUsers {
		output = append(output, mockUsers[key])
	}
	return output, nil
}

func (m Repo) DeleteUser(uuid utils.UUID) error {
	_, ok := mockUsers[uuid]
	if ok {
		delete(mockUsers, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (m Repo) AddUser(user *model.User) (*model.User, error) {
	id, _ := uuid.NewV4()
	user.UUID = utils.UUID{UUID: id}
	mockUsers[user.UUID] = user
	return user, nil
}

func (m Repo) UpdateUser(user *model.User, uuid utils.UUID) error {
	_, ok := mockUsers[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	user.UUID = uuid
	mockUsers[uuid] = user
	return nil
}

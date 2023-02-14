package repository

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"car-rental/internal/model"
	"car-rental/utils"
)

// GetUserByUUID returns User with given uuid
func (c client) GetUserByUUID(uuid utils.UUID) (*model.User, error) {
	user := &model.User{}
	if result := c.First(user, "uuid = ?", uuid); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.First")
	}
	return user, nil
}

// GetAllUsers returns all Users
func (c client) GetAllUsers() ([]*model.User, error) {
	users := []*model.User{}
	if result := c.Find(&users); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Find")
	}
	return users, nil
}

// DeleteUser returns User with given uuid
func (c client) DeleteUser(uuid utils.UUID) error {
	if uuid.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}

	user := &model.User{}
	if result := c.Where("uuid = ?", uuid.String()).Delete(user); result.Error != nil {
		return errors.Wrap(result.Error, "c.Delete")
	}
	return nil
}

// AddUser adds a new User and return it with new uuid
func (c client) AddUser(user *model.User) (*model.User, error) {
	newUUID, _ := uuid.NewV4()
	user.UUID = utils.UUID{UUID: newUUID}
	if result := c.Create(user); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Create")
	}
	return user, nil
}

// UpdateUser updates a User
func (c client) UpdateUser(user *model.User, uuid utils.UUID) error {
	if user.UUID.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}

	result := c.Save(user)
	if result.RowsAffected == 0 {
		return NoRowsAffectedErr
	}

	if result.Error != nil {
		return errors.Wrap(result.Error, "c.Save")
	}
	return nil
}

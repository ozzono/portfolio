package repository

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"car-rental/internal/model"
)

// GetRentByUUID returns Rent with given uuid
func (c client) GetRentByUUID(uuid uuid.UUID) (*model.Rent, error) {
	rent := &model.Rent{}
	if result := c.First(rent, "uuid = ?", uuid); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.First")
	}
	return rent, nil
}

// GetRentByUUID returns Rent with given vehicle uuid
func (c client) GetRentByVehicleUUID(uuid uuid.UUID) (*model.Rent, error) {
	rent := &model.Rent{}
	if result := c.Where("vehicle_uuid = ?", uuid).First(rent); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.First")
	}
	return rent, nil
}

// GetRentByUUID returns Rent with given vehicle and user uuid
func (c client) GetRentByVehicleNUserUUID(vehicleID, userID uuid.UUID) (*model.Rent, error) {
	rent := &model.Rent{}
	if result := c.
		Where("vehicle_uuid = ?", vehicleID).
		Where("user_uuid = ?", userID).
		First(rent); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.First")
	}
	return rent, nil
}

// GetAllRents returns all Rents
func (c client) GetAllRents() ([]*model.Rent, error) {
	rents := []*model.Rent{}
	if result := c.Find(&rents); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Find")
	}
	return rents, nil
}

// GetAllRents returns all Rents
func (c client) GetAllVehicleRents(vehicle *model.Vehicle) ([]*model.Rent, error) {
	rents := []*model.Rent{}
	if result := c.Where("vehicle_uuid = ?", vehicle.UUID).Find(&rents); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Find")
	}
	return rents, nil
}

// GetAllRents returns all Rents
func (c client) GetAllUserRents(user *model.User) ([]*model.Rent, error) {
	rents := []*model.Rent{}
	if result := c.Where("user_uuid = ?", user.UUID).Find(&rents); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Find")
	}
	return rents, nil
}

// DeleteRent returns Rent with given uuid
func (c client) DeleteRent(uuid uuid.UUID) error {
	if uuid.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}

	rent := &model.Rent{}
	if result := c.Where("uuid = ?", uuid.String()).Delete(rent); result.Error != nil {
		return errors.Wrap(result.Error, "c.Delete")
	}
	return nil
}

// AddRent adds a new Rent and return it with new uuid
func (c client) AddRent(rent *model.Rent) (*model.Rent, error) {
	newUUID, _ := uuid.NewV4()
	rent.UUID = newUUID
	if rent.Status == "" {
		rent.Status = "created"
	}
	rent.CreatedAt = time.Now()
	if result := c.Create(rent); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Create")
	}
	return rent, nil
}

// UpdateRent updates a Rent
func (c client) UpdateRent(rent *model.Rent, uuid uuid.UUID) error {
	if rent.UUID.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}

	result := c.Save(rent)

	if result.RowsAffected == 0 {
		return NoRowsAffectedErr
	}

	if result.Error != nil {
		return errors.Wrap(result.Error, "c.Save")
	}
	return nil
}

package repository

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"car-rental/internal/model"
)

// GetVehicleByUUID returns Vehicle with given uuid
func (c client) GetVehicleByUUID(uuid uuid.UUID) (*model.Vehicle, error) {
	vehicle := &model.Vehicle{}
	if result := c.First(vehicle, "uuid = ?", uuid); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.First")
	}
	return vehicle, nil
}

// GetAllVehicles returns all Vehicles
func (c client) GetAllVehicles() ([]*model.Vehicle, error) {
	vehicles := []*model.Vehicle{}
	if result := c.Find(&vehicles); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Find")
	}
	return vehicles, nil
}

// DeleteVehicle returns Vehicle with given uuid
func (c client) DeleteVehicle(uuid uuid.UUID) error {
	if uuid.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}

	vehicle := &model.Vehicle{}
	if result := c.Where("uuid = ?", uuid.String()).Delete(vehicle); result.Error != nil {
		return errors.Wrap(result.Error, "c.Delete")
	}
	return nil
}

// AddVehicle adds a new Vehicle and return it with new uuid
func (c client) AddVehicle(vehicle *model.Vehicle) (*model.Vehicle, error) {
	newUUID, _ := uuid.NewV4()
	vehicle.UUID = newUUID
	if result := c.Create(vehicle); result.Error != nil {
		return nil, errors.Wrap(result.Error, "c.Create")
	}
	return vehicle, nil
}

// UpdateVehicle updates a Vehicle
func (c client) UpdateVehicle(vehicle *model.Vehicle, uuid uuid.UUID) error {
	if vehicle.UUID.IsNil() {
		return errors.Wrap(ErrInvalidUUID, "cannot be nil")
	}
	result := c.Save(vehicle)
	if result.RowsAffected == 0 {
		return NoRowsAffectedErr
	}
	if result.Error != nil {
		return errors.Wrap(result.Error, "c.Save")
	}
	return nil
}

package repository

import (
	"car-rental/internal/model"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (tc testSuite) GetVehicleByUUID(uuid uuid.UUID) (*model.Vehicle, error) {
	vehicle, ok := mockVehicles[uuid]
	if ok {
		return vehicle, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (tc testSuite) GetAllVehicles() ([]*model.Vehicle, error) {
	output := []*model.Vehicle{}
	for key := range mockVehicles {
		output = append(output, mockVehicles[key])
	}
	return output, nil
}

func (tc testSuite) DeleteVehicle(uuid uuid.UUID) error {
	_, ok := mockVehicles[uuid]
	if ok {
		delete(mockVehicles, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (tc testSuite) AddVehicle(vehicle *model.Vehicle) (*model.Vehicle, error) {
	id, _ := uuid.NewV4()
	vehicle.UUID = id
	vehicle.CreatedAt = time.Now()
	mockVehicles[vehicle.UUID] = vehicle
	return vehicle, nil
}

func (tc testSuite) UpdateVehicle(vehicle *model.Vehicle, uuid uuid.UUID) error {
	_, ok := mockVehicles[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	vehicle.UUID = uuid
	vehicle.UpdatedAt = time.Now()
	mockVehicles[uuid] = vehicle
	return nil
}

package mock

import (
	"car-rental/internal/model"
	"car-rental/utils"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (m Repo) GetVehicleByUUID(uuid utils.UUID) (*model.Vehicle, error) {
	vehicle, ok := mockVehicles[uuid]
	if ok {
		return vehicle, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m Repo) GetAllVehicles() ([]*model.Vehicle, error) {
	output := []*model.Vehicle{}
	for key := range mockVehicles {
		output = append(output, mockVehicles[key])
	}
	return output, nil
}

func (m Repo) DeleteVehicle(uuid utils.UUID) error {
	_, ok := mockVehicles[uuid]
	if ok {
		delete(mockVehicles, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (m Repo) AddVehicle(vehicle *model.Vehicle) (*model.Vehicle, error) {
	id, _ := uuid.NewV4()
	vehicle.UUID = utils.UUID{UUID: id}
	vehicle.CreatedAt = time.Now()
	mockVehicles[vehicle.UUID] = vehicle
	return vehicle, nil
}

func (m Repo) UpdateVehicle(vehicle *model.Vehicle, uuid utils.UUID) error {
	_, ok := mockVehicles[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	vehicle.UUID = uuid
	vehicle.UpdatedAt = time.Now()
	mockVehicles[uuid] = vehicle
	return nil
}

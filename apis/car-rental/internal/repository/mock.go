package repository

import (
	"car-rental/internal/model"
	"time"

	"github.com/gofrs/uuid"
)

var (
	mockUser     = &model.User{}
	mockUsers    = map[uuid.UUID]*model.User{}
	mockVehicle  = &model.Vehicle{}
	mockVehicles = map[uuid.UUID]*model.Vehicle{}
	mockRent     = &model.Rent{}
	mockRents    = map[uuid.UUID]*model.Rent{}
)

func init() {
	mockUser = &model.User{
		Name:        "defaultName",
		PhoneNumber: "defaultNumber",
		Contact:     "defaultContact",
	}
	mockUsers[mockUser.UUID] = mockUser

	mockVehicle = &model.Vehicle{
		Model:        "defaultModel",
		LicensePlate: "defaultLicensePlate",
		State:        "defaultState",
		Year:         2000,
		CreatedAt:    time.Now(),
	}
	mockVehicles[mockVehicle.UUID] = mockVehicle
	mockRent = &model.Rent{
		UserUUID:    mockUser.UUID,
		VehicleUUID: mockVehicle.UUID,
		CreatedAt:   time.Now(),
	}
	mockRents[mockRent.UUID] = mockRent
}

func NewMockClient() Client {
	return testSuite{}
}

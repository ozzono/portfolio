package mock

import (
	"car-rental/internal/model"
	"car-rental/utils"
	"time"
)

type Repo struct{}

var (
	mockUser     = &model.User{}
	mockUsers    = map[utils.UUID]*model.User{}
	mockVehicle  = &model.Vehicle{}
	mockVehicles = map[utils.UUID]*model.Vehicle{}
	mockRent     = &model.Rent{}
	mockRents    = map[utils.UUID]*model.Rent{}
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

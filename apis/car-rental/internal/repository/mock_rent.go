package repository

import (
	"car-rental/internal/model"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (ts testSuite) GetRentByUUID(uuid uuid.UUID) (*model.Rent, error) {
	rent, ok := mockRents[uuid]
	if ok {
		return rent, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (ts testSuite) GetAllRents() ([]*model.Rent, error) {
	output := []*model.Rent{}
	for key := range mockRents {
		output = append(output, mockRents[key])
	}
	return output, nil
}

func (ts testSuite) DeleteRent(uuid uuid.UUID) error {
	_, ok := mockRents[uuid]
	if ok {
		delete(mockRents, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (ts testSuite) AddRent(rent *model.Rent) (*model.Rent, error) {
	id, _ := uuid.NewV4()
	rent.UUID = id
	rent.CreatedAt = time.Now()
	mockRents[rent.UUID] = rent
	return rent, nil
}

func (ts testSuite) UpdateRent(rent *model.Rent, uuid uuid.UUID) error {
	_, ok := mockRents[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	rent.UUID = uuid
	mockRents[uuid] = rent
	return nil
}

func (ts testSuite) GetAllVehicleRents(vehicle *model.Vehicle) ([]*model.Rent, error) {
	output := []*model.Rent{}
	for key := range mockRents {
		if mockRents[key].UUID == vehicle.UUID {
			output = append(output, mockRents[key])
		}
	}
	if len(output) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return output, nil
}
func (ts testSuite) GetAllUserRents(user *model.User) ([]*model.Rent, error) {
	output := []*model.Rent{}
	for key := range mockRents {
		if mockRents[key].UUID == user.UUID {
			output = append(output, mockRents[key])
		}
	}
	if len(output) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return output, nil
}
func (ts testSuite) GetRentByVehicleNUserUUID(userID, vehicleID uuid.UUID) (*model.Rent, error) {
	if vehicleID.IsNil() {
		return nil, ErrInvalidUUID
	}
	if userID.IsNil() {
		return nil, ErrInvalidUUID
	}
	for _, rent := range mockRents {
		if rent.VehicleUUID == vehicleID && rent.UserUUID == userID {
			return rent, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

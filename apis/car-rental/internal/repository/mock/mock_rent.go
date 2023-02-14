package mock

import (
	"car-rental/internal/model"
	"car-rental/utils"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (m Repo) GetRentByUUID(uuid utils.UUID) (*model.Rent, error) {
	rent, ok := mockRents[uuid]
	if ok {
		return rent, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m Repo) GetAllRents() ([]*model.Rent, error) {
	output := []*model.Rent{}
	for key := range mockRents {
		output = append(output, mockRents[key])
	}
	return output, nil
}

func (m Repo) DeleteRent(uuid utils.UUID) error {
	_, ok := mockRents[uuid]
	if ok {
		delete(mockRents, uuid)
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (m Repo) AddRent(rent *model.Rent) (*model.Rent, error) {
	id, _ := uuid.NewV4()
	rent.UUID = utils.UUID{UUID: id}
	rent.CreatedAt = time.Now()
	mockRents[rent.UUID] = rent
	return rent, nil
}

func (m Repo) UpdateRent(rent *model.Rent, uuid utils.UUID) error {
	_, ok := mockRents[uuid]
	if !ok {
		return gorm.ErrRecordNotFound
	}
	rent.UUID = uuid
	mockRents[uuid] = rent
	return nil
}

func (m Repo) GetAllVehicleRents(vehicle *model.Vehicle) ([]*model.Rent, error) {
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
func (m Repo) GetAllUserRents(user *model.User) ([]*model.Rent, error) {
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
func (m Repo) GetRentByVehicleNUserUUID(userID, vehicleID utils.UUID) (*model.Rent, error) {
	if err := vehicleID.Valid(); err != nil {
		return nil, err
	}
	if err := userID.Valid(); err != nil {
		return nil, err
	}
	for _, rent := range mockRents {
		if rent.VehicleUUID == vehicleID && rent.UserUUID == userID {
			return rent, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

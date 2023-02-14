package repository

import (
	"car-rental/internal/model"
	"car-rental/internal/repository/mock"
	"car-rental/utils"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const rentalDB = "car-rental.db"

var (
	ErrInvalidUUID    = errors.New("invalid uuid")
	NoRowsAffectedErr = errors.New("no rows affected")
)

type client struct {
	*gorm.DB
}

type Client interface {
	// GetVehicleByUUID returns Vehicle with given uuid
	GetVehicleByUUID(uuid utils.UUID) (*model.Vehicle, error)
	// GetAllVehicles returns all Vehicles
	GetAllVehicles() ([]*model.Vehicle, error)
	// DeleteVehicle returns Vehicle with given uuid
	DeleteVehicle(uuid utils.UUID) error
	// AddVehicle adds a new Vehicle and return it with new uuid
	AddVehicle(vehicle *model.Vehicle) (*model.Vehicle, error)
	// UpdateVehicle updates a Vehicle
	UpdateVehicle(vehicle *model.Vehicle, uuid utils.UUID) error

	// GetUserByUUID returns User with given uuid
	GetUserByUUID(uuid utils.UUID) (*model.User, error)
	// GetAllUsers returns all Users
	GetAllUsers() ([]*model.User, error)
	// DeleteUser returns User with given uuid
	DeleteUser(uuid utils.UUID) error
	// AddUser adds a new User and return it with new uuid
	AddUser(user *model.User) (*model.User, error)
	// UpdateUser updates a User
	UpdateUser(user *model.User, uuid utils.UUID) error

	// GetRentByUUID returns Rent with given uuid
	GetRentByUUID(uuid utils.UUID) (*model.Rent, error)
	// GetRentByVehicleNUserUUID returns Rent with given vehicle and user uuid
	GetRentByVehicleNUserUUID(vehicleID, userID utils.UUID) (*model.Rent, error)
	// GetAllRents returns all Rents
	GetAllRents() ([]*model.Rent, error)
	// GetAllVehicleRents returns all Rents
	GetAllVehicleRents(vehicle *model.Vehicle) ([]*model.Rent, error)
	// GetAllUserRents returns all Rents
	GetAllUserRents(user *model.User) ([]*model.Rent, error)
	// DeleteRent returns Rent with given uuid
	DeleteRent(uuid utils.UUID) error
	// AddRent adds a new Rent and return it with new uuid
	AddRent(rent *model.Rent) (*model.Rent, error)
	// UpdateRent updates a Rent
	UpdateRent(rent *model.Rent, uuid utils.UUID) error
}

func NewDBClient(debug bool) (Client, error) {

	db, err := gorm.Open(sqlite.Open(rentalDB), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm sqlite open")
	}

	db.AutoMigrate(&model.Vehicle{}, &model.User{}, &model.Rent{})

	if debug {
		db = db.Debug()
	}
	return client{db}, nil
}

func NewMockClient() Client {
	return mock.Repo{}
}

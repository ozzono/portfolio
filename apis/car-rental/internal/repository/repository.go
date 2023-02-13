package repository

import (
	"car-rental/internal/model"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlUser = "car_rental"
	mysqlPW   = "car_rental"
	mysqlHost = "127.0.0.1"
	mysqlDB   = "car_rental"
)

type testSuite struct {
	suite.Suite
}

var (
	ErrInvalidUUID    = errors.New("invalid uuid")
	NoRowsAffectedErr = errors.New("no rows affected")
)

type client struct {
	*gorm.DB
}

type Client interface {
	// GetVehicleByUUID returns Vehicle with given uuid
	GetVehicleByUUID(uuid uuid.UUID) (*model.Vehicle, error)
	// GetAllVehicles returns all Vehicles
	GetAllVehicles() ([]*model.Vehicle, error)
	// DeleteVehicle returns Vehicle with given uuid
	DeleteVehicle(uuid uuid.UUID) error
	// AddVehicle adds a new Vehicle and return it with new uuid
	AddVehicle(vehicle *model.Vehicle) (*model.Vehicle, error)
	// UpdateVehicle updates a Vehicle
	UpdateVehicle(vehicle *model.Vehicle, uuid uuid.UUID) error

	// GetUserByUUID returns User with given uuid
	GetUserByUUID(uuid uuid.UUID) (*model.User, error)
	// GetAllUsers returns all Users
	GetAllUsers() ([]*model.User, error)
	// DeleteUser returns User with given uuid
	DeleteUser(uuid uuid.UUID) error
	// AddUser adds a new User and return it with new uuid
	AddUser(user *model.User) (*model.User, error)
	// UpdateUser updates a User
	UpdateUser(user *model.User, uuid uuid.UUID) error

	// GetRentByUUID returns Rent with given uuid
	GetRentByUUID(uuid uuid.UUID) (*model.Rent, error)
	// GetRentByVehicleNUserUUID returns Rent with given vehicle and user uuid
	GetRentByVehicleNUserUUID(vehicleID, userID uuid.UUID) (*model.Rent, error)
	// GetAllRents returns all Rents
	GetAllRents() ([]*model.Rent, error)
	// GetAllVehicleRents returns all Rents
	GetAllVehicleRents(vehicle *model.Vehicle) ([]*model.Rent, error)
	// GetAllUserRents returns all Rents
	GetAllUserRents(user *model.User) ([]*model.Rent, error)
	// DeleteRent returns Rent with given uuid
	DeleteRent(uuid uuid.UUID) error
	// AddRent adds a new Rent and return it with new uuid
	AddRent(rent *model.Rent) (*model.Rent, error)
	// UpdateRent updates a Rent
	UpdateRent(rent *model.Rent, uuid uuid.UUID) error
}

func NewDBClient(debug bool) (Client, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPW, mysqlHost, mysqlDB)
	// "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>?options"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}

	db.AutoMigrate(&model.Vehicle{}, &model.User{}, &model.Rent{})

	if debug {
		db = db.Debug()
	}
	return client{db}, nil
}

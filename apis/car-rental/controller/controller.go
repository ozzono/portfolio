package controller

import (
	"car-rental/gateway"
	"car-rental/internal/model"
	"car-rental/internal/repository"
	"car-rental/utils"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	callback interface{}
)

type Controller struct {
	log    *zap.SugaredLogger
	client repository.Client
}

func NewController(l *zap.SugaredLogger, c repository.Client) *Controller {
	return &Controller{log: l, client: c}
}

func (c *Controller) Schedule(vehicleID uuid.UUID, userID uuid.UUID, pickUp, dropOff time.Time) (*model.Rent, error) {
	_, err := c.client.GetVehicleByUUID(vehicleID)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.GetVehicleByUUID")
	}

	_, err = c.client.GetUserByUUID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.GetUserByUUID")
	}
	now := utils.Now()
	if now.After(pickUp) {
		fmt.Println("now.After(pickUp)")
		fmt.Println("pickUp", pickUp.String())
		fmt.Println("now", now.String())
		return nil, errors.New("cannot schedule pickup at a past time")
	}

	if dropOff.Before(pickUp) {
		fmt.Println("pickUp.Before(dropOff)")
		fmt.Println("pickUp", pickUp.String())
		fmt.Println("dropOff", dropOff.String())
		return nil, errors.New("cannot schedule dropoff ealier than pickup")
	}
	cost := float64(dropOff.Sub(pickUp).Hours()) * model.RegularFee
	rent := &model.Rent{
		UserUUID:    userID,
		VehicleUUID: vehicleID,
		PickUpAt:    &pickUp,
		DropOffAt:   &dropOff,
		Status:      "scheduled",
		Cost:        cost,
		Refundable:  cost,
	}

	rent, err = c.client.AddRent(rent)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.UpdateRent")
	}

	gateway.Schedule(int32(dropOff.Sub(pickUp).Seconds()), callback)

	user, err := c.client.GetUserByUUID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "c.client.GetUserByUUID")
	}

	gateway.SendText(user.PhoneNumber, "rent scheduled successfully")

	return rent, nil
}

func (c *Controller) PickupOrDropOff(vehicleID uuid.UUID, userID uuid.UUID, status string) error {
	if status != "active" && status != "inactive" && status != "canceled" {
		return errors.Errorf("invalid status %s; must be `active` or `inactive`", status)
	}

	rent, err := c.client.GetRentByVehicleNUserUUID(vehicleID, userID)
	if err != nil {
		return errors.Wrap(err, "c.client.GetRentByVehicleNUserUUID")
	}

	rent.Status = status

	t := utils.Now()
	msg := ""
	extraMsg := ""
	if status == "active" {
		rent.PickedAt = &t
		msg = "vehicle picked successfully\n"
	}

	elapsedTime := time.Duration(t.Sub(rent.CreatedAt)).Hours()
	if status == "inactive" {
		rent.DroppedAt = &t
		if rent.DroppedAt.After(*rent.DropOffAt) {
			extra := elapsedTime * model.OvertimeFee
			extraMsg = fmt.Sprintf("applying extra fee of %.2f due to overtime duration\n", extra)
			rent.Cost += extra
		}
		msg = "vehicle dropped successfully\n"
	}

	if status == "canceled" {
		rent.CanceledAt = &t
		if elapsedTime < 24 {
			extraMsg = "rent canceled without refund"
			rent.Refundable = 0
		}
		if 24 < elapsedTime && elapsedTime < 48 {
			rent.Refundable = rent.Cost / 4
			extraMsg = fmt.Sprintf("rent canceled with 25%% refund of %f", rent.Refundable)
		}
		msg = "rent schedule canceled successfully\n"
	}

	err = c.client.UpdateRent(rent, rent.UUID)
	if err != nil {
		return errors.Wrap(err, "c.client.UpdateRent")
	}

	user, err := c.client.GetUserByUUID(userID)
	if err != nil {
		return errors.Wrap(err, "c.client.GetUserByUUID")
	}

	gateway.SendText(user.PhoneNumber, msg+extraMsg)

	return nil
}

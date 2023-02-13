package repository

import (
	"car-rental/internal/model"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	originalLenRents = 0

	testRent = &model.Rent{}
)

func (ts testSuite) Test21CreateRent() {
	rents, err := ts.GetAllRents()
	assert.NoError(ts.T(), err)
	originalLenRents = len(rents)
	newRent := &model.Rent{
		UserUUID:    testUser.UUID,
		VehicleUUID: testVehicle.UUID,
	}
	testRent, err = ts.AddRent(newRent)
	assert.NoError(ts.T(), err)
	assert.NotNil(ts.T(), testRent.UUID.IsNil())
}

func (ts testSuite) Test3GetRentByUUID() {
	rent, err := ts.GetRentByUUID(testRent.UUID)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), rent, testRent)
}

func (ts testSuite) Test4GetRentByUserNVehicle() {
	rent, err := ts.GetRentByVehicleNUserUUID(testRent.UserUUID, testRent.VehicleUUID)
	assert.NoError(ts.T(), err)
	testRent.CreatedAt = rent.CreatedAt
	assert.Equal(ts.T(), rent, testRent)
}

func (ts testSuite) Test5GetAllRent() {
	rents, err := ts.GetAllRents()
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), len(rents), originalLenRents+1)
	found := false
find:
	for i := range rents {
		if rents[i].UUID == testRent.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)
}

func (ts testSuite) Test6UpdateRent() {
	newRent := testRent.Copy()
	time.Sleep(time.Second) //delaying just 1 sec to make it a bit different
	t := time.Now()
	newRent.DroppedAt = &t
	err := ts.UpdateRent(newRent, newRent.UUID)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), newRent.DroppedAt, testRent.DroppedAt)
}

func (ts testSuite) Test7DeleteRent() {
	err := ts.DeleteRent(testRent.UUID)
	assert.NoError(ts.T(), err, "deleting rent")
	_, err = ts.GetRentByUUID(testRent.UUID)
	assert.ErrorIs(ts.T(), err, gorm.ErrRecordNotFound)

	err = ts.DeleteUser(testRent.UserUUID)
	assert.NoError(ts.T(), err, "deleting user")
	_, err = ts.GetUserByUUID(testRent.UUID)
	assert.ErrorIs(ts.T(), err, gorm.ErrRecordNotFound)

	err = ts.DeleteVehicle(testRent.VehicleUUID)
	assert.NoError(ts.T(), err, "deleting vehicle")
	_, err = ts.GetVehicleByUUID(testRent.UUID)
	assert.ErrorIs(ts.T(), err, gorm.ErrRecordNotFound)
}

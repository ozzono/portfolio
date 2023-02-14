package repository

import (
	"car-rental/internal/model"

	"github.com/stretchr/testify/assert"
)

var (
	testVehicle         = &model.Vehicle{}
	originalLenVehicles = 0
)

func (ts testSuite) Test20CreateVehicle() {
	vehicles, err := ts.client.GetAllVehicles()
	assert.NoError(ts.T(), err)
	originalLenVehicles = len(vehicles)
	newVehicle := &model.Vehicle{
		Model:        "Model         ",
		LicensePlate: "LicensePlate ",
		State:        "State         ",
		Year:         2000,
	}
	testVehicle, err = ts.client.AddVehicle(newVehicle)
	assert.NoError(ts.T(), err)
}

func (ts testSuite) Test3GetVehicle() {
	vehicle, err := ts.client.GetVehicleByUUID(testVehicle.UUID)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), vehicle, testVehicle)
	testVehicle.CreatedAt = vehicle.CreatedAt
}

func (ts testSuite) Test4GetAllVehicle() {
	vehicles, err := ts.client.GetAllVehicles()
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), len(vehicles), originalLenVehicles+1)
	found := false
find:
	for i := range vehicles {
		if vehicles[i].UUID == testVehicle.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)

}

func (ts testSuite) Test5UpdateVehicle() {

	newVehicle := testVehicle.Copy()
	newVehicle.Archived = !testVehicle.Archived
	newVehicle.Available = !testVehicle.Available

	err := ts.client.UpdateVehicle(newVehicle, newVehicle.UUID)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), newVehicle.Archived, testVehicle.Archived)
	assert.NotEqual(ts.T(), newVehicle.Available, testVehicle.Available)
	assert.False(ts.T(), newVehicle.UpdatedAt.IsZero())
}

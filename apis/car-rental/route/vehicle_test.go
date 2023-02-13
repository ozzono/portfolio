package route

import (
	"bytes"
	"car-rental/internal/model"
	"car-rental/utils"
	"net/http"
	"net/http/httptest"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

var (
	testVehicle = &model.Vehicle{
		Model:        "Benz Patent Motor Car",
		LicensePlate: "37435",
		State:        "Germany",
		Archived:     true,
		Available:    false,
		Year:         1886,
	}
)

func (ts *testSuite) Test10AddVehicle() {
	byteVehicle, err := json.Marshal(testVehicle)
	assert.NoError(ts.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/vehicles", bytes.NewBuffer(byteVehicle))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)

	outputVehicle := &model.Vehicle{}
	err = json.Unmarshal(rec.Body.Bytes(), outputVehicle)
	assert.NoError(ts.T(), err)

	assert.False(ts.T(), outputVehicle.UUID.IsNil())

	testVehicle.UUID = outputVehicle.UUID
	testVehicle.CreatedAt = outputVehicle.CreatedAt
	testVehicle.UpdatedAt = outputVehicle.UpdatedAt

	assert.Equal(ts.T(), testVehicle, outputVehicle)
	testVehicle = outputVehicle
}

func (ts *testSuite) Test20GetVehicle() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/vehicles/"+testVehicle.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	outputVehicle := model.Vehicle{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputVehicle)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), outputVehicle, *testVehicle)
}

func (ts *testSuite) Test30GetVehicles() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/vehicles", nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	outputVehicles := []model.Vehicle{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputVehicles)
	assert.NoError(ts.T(), err)
	assert.True(ts.T(), len(outputVehicles) > 0)
	found := false
find:
	for i := range outputVehicles {
		if outputVehicles[i].UUID == testVehicle.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)

}

func (ts *testSuite) Test40UpdateVehicle() {
	outputVehicle := model.Vehicle{
		Model:        "Apollo 11",
		LicensePlate: "11",
		State:        "Florida",
		Archived:     true,
		Available:    false,
		Year:         1969,
	}
	req, err := http.NewRequest(http.MethodPut, "/api/v1/vehicles/"+testVehicle.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	err = json.Unmarshal(rec.Body.Bytes(), &outputVehicle)
	assert.NoError(ts.T(), err)
	testVehicle.UpdatedAt = outputVehicle.UpdatedAt
	assert.NotEqual(ts.T(), outputVehicle, *testVehicle)
}

func (ts *testSuite) Test99DeleteVehicle() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/vehicles/"+testVehicle.UUID.String(), nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/vehicles/"+testVehicle.UUID.String(), nil)
	rec = httptest.NewRecorder()
	assert.NoError(ts.T(), err)
	ts.Serve(req, rec)

	outputMsg := utils.APIError{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputMsg)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), outputMsg.ErrorCode, http.StatusNotFound)
	assert.Equal(ts.T(), outputMsg.ErrorMessage, "record not found")
}

package route

import (
	"bytes"
	"car-rental/internal/model"
	"car-rental/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	testRent = &model.Rent{}
)

func (ts *testSuite) Test11AddRent() {
	testRent.UserUUID = testUser.UUID
	testRent.VehicleUUID = testVehicle.UUID
	byteRent, err := json.Marshal(testRent)
	assert.NoError(ts.T(), err)

	rec, err := ts.CallRentAPI("", http.MethodPost, bytes.NewBuffer(byteRent))
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	outputRent := &model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), outputRent)
	assert.NoError(ts.T(), err)

	assert.False(ts.T(), outputRent.UUID.IsNil())
	assert.Equal(ts.T(), testRent.UserUUID, outputRent.UserUUID)
	assert.Equal(ts.T(), testRent.VehicleUUID, outputRent.VehicleUUID)
	*testRent = *outputRent
}

func (ts *testSuite) Test21GetRent() {
	rec, err := ts.CallRentAPI(testRent.UUID.String(), http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	outputRent := model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)

	assert.Equal(ts.T(), outputRent.CreatedAt.Format(utils.TimeFormat), testRent.CreatedAt.Format(utils.TimeFormat))
	assert.NoError(ts.T(), err)
}

func (ts *testSuite) Test31GetRents() {
	rec, err := ts.CallRentAPI("", http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	outputRents := []model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRents)
	assert.NoError(ts.T(), err)
	assert.True(ts.T(), len(outputRents) > 0)
	found := false
find:
	for i := range outputRents {
		if outputRents[i].UUID == testRent.UUID {
			found = true
			break find
		}
	}
	assert.True(ts.T(), found)
}

func (ts *testSuite) Test41UpdateRent() {
	t := time.Now()
	testRent.PickUpAt = &t
	byteData, err := json.Marshal(testRent)
	assert.NoError(ts.T(), err)

	rec, err := ts.CallRentAPI(testRent.UUID.String(), http.MethodPut, strings.NewReader(string(byteData)))
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	outputRent := model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputRent, *testRent)
}

func (ts *testSuite) Test50ScheduleRent() {
	// deleting rent so I can recreate it using schedulement
	rec, err := ts.CallRentAPI(testRent.UUID.String(), http.MethodDelete)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
	}

	outputRent := model.Rent{}
	path := fmt.Sprintf("/api/v1/vehicles/%s/%s/schedule", testVehicle.UUID.String(), testUser.UUID.String())

	u, err := url.Parse(path)
	assert.NoError(ts.T(), err)
	values := u.Query()
	values.Add("pickup_time", time.Now().Add(time.Hour).Format(utils.TimeFormat))
	values.Add("dropoff_time", time.Now().Add(time.Duration(2)*time.Hour).Format(utils.TimeFormat))
	u.RawQuery = values.Encode()
	path = u.String()
	req, err := http.NewRequest(http.MethodPut, path, nil)
	rec = httptest.NewRecorder()
	ts.Serve(req, rec)

	assert.NoError(ts.T(), err)

	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputRent, *testRent)
	assert.Equal(ts.T(),
		outputRent.CreatedAt.Format(utils.TimeFormat),
		testRent.CreatedAt.Format(utils.TimeFormat),
	)
	assert.NotEqual(ts.T(),
		outputRent.PickUpAt.Format(utils.TimeFormat),
		testRent.PickUpAt.Format(utils.TimeFormat),
	)
	assert.NotNil(ts.T(), outputRent.DropOffAt)
	*testRent = outputRent
}

func (ts *testSuite) Test51PickRent() {
	path := fmt.Sprintf("/api/v1/vehicles/%s/%s/update?status=active", testVehicle.UUID.String(), testUser.UUID.String())
	req, err := http.NewRequest(http.MethodPut, path, nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	rec, err = ts.CallRentAPI(testRent.UUID.String(), http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
	}

	outputRent := model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputRent, *testRent)
	assert.NotNil(ts.T(), outputRent.PickedAt)
}

func (ts *testSuite) Test52DropRent() {
	path := fmt.Sprintf("/api/v1/vehicles/%s/%s/update?status=inactive", testVehicle.UUID.String(), testUser.UUID.String())
	req, err := http.NewRequest(http.MethodPut, path, nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	rec, err = ts.CallRentAPI(testRent.UUID.String(), http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
	}

	outputRent := model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputRent, *testRent)
	assert.NotNil(ts.T(), outputRent.DroppedAt)
}

func (ts *testSuite) Test53DropRent() {
	path := fmt.Sprintf("/api/v1/vehicles/%s/%s/update?status=canceled", testVehicle.UUID.String(), testUser.UUID.String())
	req, err := http.NewRequest(http.MethodPut, path, nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
		return
	}

	rec, err = ts.CallRentAPI(testRent.UUID.String(), http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
	}

	outputRent := model.Rent{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputRent)
	assert.NoError(ts.T(), err)
	assert.NotEqual(ts.T(), outputRent, *testRent)
	assert.NotNil(ts.T(), outputRent.CanceledAt)
}

func (ts *testSuite) Test63DeleteRent() {
	//deleting rent
	rec, err := ts.CallRentAPI(testRent.UUID.String(), http.MethodDelete)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusOK, rec.Code)
	if http.StatusOK != rec.Code {
		msg := utils.APIError{}
		err = json.Unmarshal(rec.Body.Bytes(), &msg)
		assert.NoError(ts.T(), err)
		ts.T().Log("err msg", msg.LogTxt())
	}

	// getting nonexistent rent
	rec, err = ts.CallRentAPI(testRent.UUID.String(), http.MethodGet)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), http.StatusNotFound, rec.Code)

	outputMsg := utils.APIError{}
	err = json.Unmarshal(rec.Body.Bytes(), &outputMsg)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), outputMsg.ErrorCode, http.StatusNotFound)
	assert.Equal(ts.T(), outputMsg.ErrorMessage, "record not found")
}

func (ts *testSuite) CallRentAPI(uuid string, method string, args ...io.Reader) (*httptest.ResponseRecorder, error) {
	var arg io.Reader
	if len(args) > 0 {
		arg = args[0]
	}
	if uuid != "" {
		ts.T().Log("uuid", uuid)
		uuid = "/" + uuid
	}
	req, err := http.NewRequest(method, "/api/v1/rents"+uuid, arg)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}
	return rec, err
}

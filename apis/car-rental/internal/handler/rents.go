package handler

import (
	"car-rental/internal/model"
	"car-rental/utils"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const datetimeExp = `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`

// @Summary Gets a rent from the database filtered by uuid
// @Produce json
// @Tags    rents
// @Param   rent    path    string     true        "rent uuid"
// @Success 200 {object} model.Rent	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 404 {object} utils.APIError "record not found"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rents/:rent [get]
func (h *Handler) GetRent(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("rent"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	rent, err := h.Client.GetRentByUUID(utils.UUID{UUID: id})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HTTPErrJSON(ctx, http.StatusNotFound, "record not found")
		return
	}

	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetRentByUUID"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, rent)
}

// @Summary Gets all rents from the database
// @Produce json
// @Tags    rents
// @Success 200 {object} []model.Rent	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rents/ [get]
func (h *Handler) GetRents(ctx *gin.Context) {
	rents, err := h.Client.GetAllRents()
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetAllRents"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, rents)
}

// @Summary Adds a rent to the database
// @Produce json
// @Tags    rents
// @Param   rent    body    model.Rent     true        "rent data"
// @Success 200 {object} model.Rent	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rents/ [post]
func (h *Handler) AddRent(ctx *gin.Context) {
	rent := &model.Rent{}
	if err := ctx.Bind(rent); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}
	rent, err := h.Client.AddRent(rent)
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.AddRent"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}
	ctx.JSON(http.StatusOK, rent)
}

// @Summary Updates a rent from the database filtered by uuid
// @Produce json
// @Tags    rents
// @Param   rent    body    model.Rent true        "rent data"
// @Param   rent_id path    string     true        "rent uuid"
// @Success 200 {object} model.Rent	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rents/:rent [put]
func (h *Handler) UpdateRent(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("rent"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	rent := &model.Rent{}
	if err := ctx.Bind(rent); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	if err := rent.Valid(); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rent.UUID = utils.UUID{UUID: id}

	err = h.Client.UpdateRent(rent, utils.UUID{UUID: id})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HTTPErrJSON(ctx, http.StatusNotFound, "record not found")
		return
	}
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.UpdateRent"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}
	ctx.JSON(http.StatusOK, rent)
}

// @Summary Deletes a rent from the database by uuid
// @Produce json
// @Tags    rents
// @Param   rent    path    string     true        "rent uuid"
// @Success 200 "ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rents/:rent [delete]
func (h *Handler) DeleteRent(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("rent"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.Client.DeleteRent(utils.UUID{UUID: id})
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.DeleteRent"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Schedules a vehicle rent
// @Produce json
// @Tags    controller
// @Param   user_id			path		string		true		"user uuid"
// @Param   vehicle_id		path		string		true		"vehicle uuid"
// @Param   pickup_time		query		time.Time	true		"scheduled pickup date"
// @Param   dropoff_time	query		time.Time	true		"schedule dropoff date"
// @Success 200 "ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 404 {object} utils.APIError "record not found"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rent/:vehicle/:user/schedule [put]
func (h *Handler) ScheduleVehicle(ctx *gin.Context) {
	vehicleID, err := uuid.FromString(ctx.Param("vehicle"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	userID, err := uuid.FromString(ctx.Param("user"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid vehicle id")
		return
	}

	pickUpTime := ctx.Query("pickup_time")
	match, err := regexp.Match(datetimeExp, []byte(pickUpTime))
	if err != nil || !match {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid pickup time input; must use YYYY-MM-DD hh:mm:ss format")
		return
	}

	pickUp, err := time.Parse(utils.TimeFormat, pickUpTime)
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, fmt.Sprintf("invalid time input; err %v", err))
		return
	}

	dropOffTime := ctx.Query("dropoff_time")
	match, err = regexp.Match(datetimeExp, []byte(dropOffTime))
	if err != nil || !match {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid time input; must use YYYY-MM-DD hh:mm:ss format")
		return
	}

	dropOff, err := time.Parse(utils.TimeFormat, dropOffTime)
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest,
			fmt.Sprintf("invalid time input; err %v", err),
		)
		return
	}

	rent, err := h.ctrl.Schedule(utils.UUID{UUID: vehicleID}, utils.UUID{UUID: userID}, pickUp, dropOff)
	if err != nil {
		h.log.Errorf("h.ctrl.Schedule - %v", err)
		utils.HTTPErrJSON(ctx, http.StatusBadRequest,
			fmt.Sprintf("invalid time input; err %v", err),
		)
		return
	}

	ctx.JSON(http.StatusOK, rent)
}

// @Summary Pickup or DropOff a vehicle
// @Produce json
// @Tags    controller
// @Param   user_id			path		string		true		"user uuid"
// @Param   vehicle_id		path		string		true		"vehicle uuid"
// @Param   status			query		string		true		"rent new status"
// @Success 200 "ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 404 {object} utils.APIError "record not found"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /rent/:vehicle/:user/update [put]
func (h *Handler) PickupOrDropOffVehicle(ctx *gin.Context) {
	vehicleID, err := uuid.FromString(ctx.Param("vehicle"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid user id")
		return
	}

	userID, err := uuid.FromString(ctx.Param("user"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid vehicle id")
		return
	}

	status := ctx.Query("status")
	if status != "active" && status != "inactive" && status != "canceled" {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest,
			fmt.Sprintf("invalid rent status %s; must be `active` or `inactive`", status),
		)
		return
	}

	fmt.Println(h == nil)
	err = h.ctrl.PickupOrDropOff(utils.UUID{UUID: vehicleID}, utils.UUID{UUID: userID}, status)
	if err != nil {
		h.log.Errorf("h.ctrl.PickupOrDropOff - %v", err)
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "contact system admin")
		return
	}
	ctx.Status(http.StatusOK)
}

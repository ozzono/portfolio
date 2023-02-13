package handler

import (
	"car-rental/internal/model"
	"car-rental/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary Gets a vehicle from the database filtered by uuid
// @Produce json
// @Tags    vehicles
// @Param   vehicle    path    string     true        "vehicle uuid"
// @Success 200 {object} model.Vehicle	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 404 {object} utils.APIError "record not found"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /vehicles/:vehicle [get]
func (h *Handler) GetVehicle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("vehicle"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	vehicle, err := h.Client.GetVehicleByUUID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HTTPErrJSON(ctx, http.StatusNotFound, "record not found")
		return
	}

	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetVehicleByUUID"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

// @Summary Gets all vehicles from the database
// @Produce json
// @Tags    vehicles
// @Success 200 {object} []model.Vehicle	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /vehicles/ [get]
func (h *Handler) GetVehicles(ctx *gin.Context) {
	vehicles, err := h.Client.GetAllVehicles()
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetAllVehicles"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, vehicles)
}

// @Summary Adds a vehicle to the database
// @Produce json
// @Tags    vehicles
// @Param   vehicle    body    model.Vehicle     true        "vehicle data"
// @Success 200 {object} model.Vehicle	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /vehicles/ [post]
func (h *Handler) AddVehicle(ctx *gin.Context) {
	vehicle := &model.Vehicle{}
	if err := ctx.Bind(vehicle); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	vehicle, err := h.Client.AddVehicle(vehicle)
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.AddVehicle"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}
	ctx.JSON(http.StatusOK, vehicle)
}

// @Summary Updatess a vehicle from the database filtered by uuid
// @Produce json
// @Tags    vehicles
// @Param   vehicle    body    model.Vehicle true        "vehicle data"
// @Param   vehicle_id path    string     true        "vehicle uuid"
// @Success 200 {object} model.Vehicle	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /vehicles/:vehicle [put]
func (h *Handler) UpdateVehicle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("vehicle"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	vehicle := &model.Vehicle{}
	if err := ctx.Bind(vehicle); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	vehicle.UUID = id

	err = h.Client.UpdateVehicle(vehicle, id)
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.UpdateVehicle"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

// @Summary Deletes a vehicle from the database by uuid
// @Produce json
// @Tags    vehicles
// @Param   vehicle    path    string     true        "vehicle uuid"
// @Success 200 "ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /vehicles/:vehicle [delete]
func (h *Handler) DeleteVehicle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("vehicle"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.Client.DeleteVehicle(id)
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.DeleteVehicle"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.Status(http.StatusOK)
}

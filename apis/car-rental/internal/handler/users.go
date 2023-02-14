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

// @Summary Gets a user from the database filtered by uuid
// @Produce json
// @Tags    users
// @Param   user    path    string     true        "user uuid"
// @Success 200 {object} model.User	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 404 {object} utils.APIError "record not found"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /users/:user [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("user"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := h.Client.GetUserByUUID(utils.UUID{UUID: id})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.HTTPErrJSON(ctx, http.StatusNotFound, "record not found")
		return
	}

	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetUserByUUID"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Gets all users from the database
// @Produce json
// @Tags    users
// @Success 200 {object} []model.User	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /users/ [get]
func (h *Handler) GetUsers(ctx *gin.Context) {
	users, err := h.Client.GetAllUsers()
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.GetAllUsers"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary Adds a user to the database
// @Produce json
// @Tags    users
// @Param   user    body    model.User     true        "user data"
// @Success 200 {object} model.User	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /users/ [post]
func (h *Handler) AddUser(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		h.log.Info("ctx.Bind", err)
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}
	user, err := h.Client.AddUser(user)
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.AddUser"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary Updatess a user from the database filtered by uuid
// @Produce json
// @Tags    users
// @Param   user    body    model.User true        "user data"
// @Param   user_id path    string     true        "user uuid"
// @Success 200 {object} model.User	"ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /users/:user [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("user"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	user.UUID = utils.UUID{UUID: id}

	err = h.Client.UpdateUser(user, utils.UUID{UUID: id})
	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.UpdateUser"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Deletes a user from the database by uuid
// @Produce json
// @Tags    users
// @Param   user    path    string     true        "user uuid"
// @Success 200 "ok"
// @Failure 400 {object} utils.APIError "invalid input"
// @Failure 500 {object} utils.APIError "internal error"
// @Router  /users/:user [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("user"))
	if err != nil {
		utils.HTTPErrJSON(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.Client.DeleteUser(utils.UUID{UUID: id})

	if err != nil {
		h.log.Error(errors.Wrap(err, "h.Client.DeleteUser"))
		utils.HTTPErrJSON(ctx, http.StatusInternalServerError, "contact system admin")
		return
	}

	ctx.Status(http.StatusOK)
}

package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"ports/internal/models"
	"ports/internal/repository"
	"ports/pkg/log"
	"ports/utils"
)

type portHandlers struct {
	group  *gin.RouterGroup
	logger log.Logger
	helper utils.Helper
	svc    repository.Service
}

func NewPortHandlers(
	group *gin.RouterGroup,
	logger log.Logger,
	svc repository.Service,
	helper utils.Helper,
) *portHandlers {
	return &portHandlers{
		group:  group,
		logger: logger,
		svc:    svc,
		helper: helper,
	}
}

var (
	EmptySignVersion = errors.New("signature version not found")
	EmptySignature   = errors.New("signature not found")
	EmptyObjId       = errors.New("invalid object id; cannot be empty")
)

func (h *portHandlers) get() gin.HandlerFunc {
	return func(c *gin.Context) {
		strID := c.Param("id")
		if strID == "" {
			h.helper.HandleHTTPError(c, http.StatusBadRequest, "invalid id", errors.New("cannot be empty"))
			return
		}
		_, err := strconv.Atoi(strID)
		if err != nil {
			h.helper.HandleHTTPError(c, http.StatusBadRequest, "invalid id", err)
			return
		}
		port, err := h.svc.Get(c.Request.Context(), strID)
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when fetching port data", err) {
			return
		}
		c.JSON(http.StatusOK, port)
	}
}

func (h *portHandlers) getByCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")

		port, err := h.svc.GetByCode(c.Request.Context(), code)
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when fetching port data", err) {
			return
		}
		c.JSON(http.StatusOK, port)
	}
}

func (h *portHandlers) query() gin.HandlerFunc {
	return func(c *gin.Context) {
		port, err := h.svc.Query(c.Request.Context())
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when fetching all ports data", err) {
			return
		}
		c.JSON(http.StatusOK, port)
	}
}

func (h *portHandlers) del() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			h.helper.HandleHTTPError(c, http.StatusBadRequest, "cannot be empty", errors.New("invalid id"))
			return
		}
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when fetching all ports data", h.svc.Delete(c.Request.Context(), c.Param("id"))) {
			return
		}
		c.JSON(http.StatusOK, models.ResponseOK{Message: "port successfully removed"})
	}
}

func (h *portHandlers) upsert() gin.HandlerFunc {
	return func(c *gin.Context) {

		upd := repository.UpSertPortRequest{}
		if err := c.BindJSON(&upd); h.helper.HandleHTTPError(c, http.StatusBadRequest, "invalid request body", err) {
			return
		}

		port, err := h.svc.UpSert(c, upd)
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when updating port", err) {
			return
		}
		c.JSON(http.StatusOK, port)
	}
}

func (h *portHandlers) parseJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		if h.helper.HandleHTTPError(c, http.StatusInternalServerError, "error when parsing all ports from json file", h.svc.ParseJson(c)) {
			return
		}
		c.JSON(http.StatusOK, models.ResponseOK{Message: "successfully parsed and added to db all ports from json file"})
	}
}

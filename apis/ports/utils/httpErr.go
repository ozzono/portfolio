package utils

import (
	"net/http"

	"ports/internal/models"
	"ports/pkg/log"

	"github.com/gin-gonic/gin"
)

type Helper struct {
	log log.Logger
}

// This function is to handle all HTTP client errors
// please, make sure you write proper message error
func (h Helper) HandleHTTPError(c *gin.Context, code int, message string, err error) bool {
	if err != nil && code == http.StatusBadRequest {
		h.log.Errorf("Error -> %v", message, err)
		c.JSON(http.StatusBadRequest, models.FailureResponse{
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusInternalServerError {
		h.log.Errorf("Error -> %v", message, err)
		c.JSON(http.StatusInternalServerError, models.FailureResponse{
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusUnauthorized {
		h.log.Errorf("Error -> %v", message, err)
		c.JSON(http.StatusUnauthorized, models.FailureResponse{
			Message: message,
			Error:   err.Error(),
		})
		return true
	} else if err != nil && code == http.StatusConflict {
		h.log.Errorf("Error -> %v", message, err)
		c.JSON(http.StatusConflict, models.FailureResponse{
			Message: message,
			Error:   err.Error(),
		})
		return true
	}
	return false
}

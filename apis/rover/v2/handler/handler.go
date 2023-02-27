package handler

import (
	"fmt"
	"net/http"
	"rover/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Sets the Plateau
// @Description shows current plateau state
// @Success 200 {string} string	"ok"
// @Router /plateau/show [get]
func ShowPlateau(c *gin.Context) {
	plateau, err := controller.Plateau.Show()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}
	c.String(http.StatusOK, plateau)
}

// @Summary Sets the Plateau
// @Description Sets a new plateau or blanks the existent one
// @Produce  json
// @Param   width     query    int     true        "Plateau width"
// @Param   height    query    int     true        "Plateau height"
// @Success 200 "ok"
// @Failure 400 {object} APIError "width and height are required"
// @Router /plateau/set [get]
func SetPlateau(c *gin.Context) {
	w := c.Query("width")
	width, err := strconv.Atoi(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: "invalid width; must be integer"})
		return
	}

	h := c.Query("height")
	height, err := strconv.Atoi(h)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: "invalid width; must be integer"})
		return
	}

	controller.SetPlateau(width, height)

	c.String(http.StatusOK, "blank plateau seted")
}

// @Summary Land rover
// @Description adds a new rover to the plateau
// @Param   x            query    int        true        "x coordinate"
// @Param   y            query    int        true        "y coordinate"
// @Param   direction    query    string     true        "rover direction"
// @Success 200 {int} int	"ok"
// @Failure 400 {object} APIError "width, height and direction are required"
// @Router /rover [put]
func LandRover(c *gin.Context) {
	w := c.Query("x")
	width, err := strconv.Atoi(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: fmt.Sprintf("invalid width %s; must be integer", w)})
		return
	}

	h := c.Query("y")
	height, err := strconv.Atoi(h)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: fmt.Sprintf("invalid height %s; must be integer", h)})
		return
	}

	direction := c.Query("direction")

	rover, err := controller.NewRover(width, height, direction)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: fmt.Sprintf("failed to create rover: %v", err)})
		return
	}

	err = controller.Plateau.AddRover(rover)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: fmt.Sprintf("failed to add rover to the plateau: %v", err)})
		return
	}

	c.String(http.StatusOK, fmt.Sprint(rover.ID))
}

// @Summary Moves the Rover
// @Description Moves the rover in the plateau
// @Accept  json
// @Produce  json
// @Param   id          path     int        true        "rover id"
// @Param   movement    query    string     true        "rover movement"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "invalid movement"
// @Failure 404 {object} APIError "rover not found"
// @Router /rover/{id} [get]
func MoveRover(c *gin.Context) {
	movement := c.Query("movement")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: fmt.Sprintf("invalid id: %v", err)})
		return
	}

	rover, err := controller.Plateau.GerRoverByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, APIError{ErrorCode: http.StatusNotFound, ErrorMessage: err.Error()})
		return
	}

	for _, m := range movement {
		err = rover.Move(string(m))
		if err != nil {
			c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: err.Error()})
			return
		}
	}

	position, err := rover.Position()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{ErrorCode: http.StatusBadRequest, ErrorMessage: err.Error()})
		return
	}
	c.String(http.StatusOK, position)
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

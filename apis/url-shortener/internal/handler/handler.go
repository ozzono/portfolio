package handler

import (
	"log"
	"net/http"
	"url-shortener/internal/database"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	Router *gin.Engine
	DB     *database.Client
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func NewHandler() (*Handler, error) {

	db, err := database.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "database.NewClient")
	}
	h := new(Handler)
	h.DB = db

	router := gin.Default()
	router.GET("/ping", pong)
	router.GET("/:id", h.redirect)

	g := router.Group("/api")
	g.PUT("", h.AddURL)
	g.GET("/:id", h.GetURL)
	g.DELETE("/:id", h.DelURL)

	h.Router = router
	return h, nil
}

func (h *Handler) AddURL(c *gin.Context) {
	url := new(models.URL)
	if err := c.BindJSON(url); err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "invalid input ", Err: err}, http.StatusBadRequest, err)
		return
	}

	url, err := h.DB.AddURL(url, true)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "internal error"}, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *Handler) GetURL(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "invalid id ", Err: err}, http.StatusBadRequest, err)
		return
	}
	url, found, err := h.DB.FindURLByID(&models.URL{ID: objID}, true)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "internal error"}, http.StatusBadRequest, err)
		return
	}

	if !found {
		models.HTTPErr(c, models.ErrMsg{Msg: "url not found"}, http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *Handler) DelURL(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "invalid id ", Err: err}, http.StatusBadRequest, err)
		return
	}

	if err = h.DB.DelURL(&models.URL{ID: objID}); err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "internal error"}, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) redirect(c *gin.Context) {
	id := c.Param("id")
	url, found, err := h.DB.FindURLByShortened(&models.URL{Shortened: id}, false)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "internal error"}, http.StatusInternalServerError, err)
		return
	}

	if !found {
		models.HTTPErr(c, models.ErrMsg{Msg: "invalid url; path not found"}, http.StatusNotFound, nil)
		return
	}

	_, err = h.DB.IncrementURL(url, true)
	if err != nil {
		models.HTTPErr(c, models.ErrMsg{Msg: "internal error"}, http.StatusInternalServerError, err)
		return
	}
	log.Printf("redirecting to %s", url.Source)
	http.Redirect(c.Writer, c.Request, url.Source, http.StatusPermanentRedirect)
	// c.Redirect(http.StatusOK, url.Source)
}

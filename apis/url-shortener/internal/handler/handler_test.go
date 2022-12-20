package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	sourceURL = "https://go.dev"
)

var (
	testURL = new(models.URL)
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	handler.Router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "pong", w.Body.String())
}

func TestAPIPutRoute(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	t.Log("putting url")
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPut, "/api", models.URLToReader(sourceURL))
	require.NoError(t, err, "http.NewRequest")
	handler.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err, "io.ReadAll")

	err = json.Unmarshal(body, testURL)
	require.NoError(t, err, "json.Unmarshal")
}

func TestAPIGetRoute(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	t.Log("getting url")
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/"+testURL.ID.Hex(), nil)
	require.NoError(t, err, "http.NewRequest")
	handler.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err, "io.ReadAll")

	url := new(models.URL)
	err = json.Unmarshal(body, url)
	require.NoError(t, err, "json.Unmarshal")
}

func TestAPIRedirect(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	t.Log("redirecting url")
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/"+testURL.Shortened, nil)
	require.NoError(t, err, "http.NewRequest")
	handler.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusPermanentRedirect, w.Code)
}

func TestAPIDelRoute(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	t.Log("deleting url")
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodDelete, "/api/"+testURL.ID.Hex(), nil)
	require.NoError(t, err, "http.NewRequest")
	handler.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestAPIDelCheck(t *testing.T) {
	handler, err := NewHandler()
	require.NoError(t, err, "NewHandler")

	t.Log("checking deletion")
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/"+testURL.ID.Hex(), nil)
	require.NoError(t, err, "http.NewRequest")
	handler.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNoContent, w.Code)
}

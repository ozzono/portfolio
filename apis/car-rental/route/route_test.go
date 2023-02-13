package route

import (
	"car-rental/internal/handler"
	"car-rental/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

// all tests are named considering alphabetical order
// as thats the default behaviour
type testSuite struct {
	suite.Suite
	routes *gin.Engine
}

func TestRoute(t *testing.T) {
	suite.Run(t, &testSuite{
		routes: Routes(
			handler.NewHandler(
				repository.NewMockClient(),
				zap.NewExample().Sugar(), false,
			), false,
		)})
}

func (ts *testSuite) Serve(req *http.Request, rec *httptest.ResponseRecorder) {
	ts.routes.ServeHTTP(rec, req)
}

func (ts *testSuite) Test0Ping() {
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	ts.Serve(req, rec)
	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), rec.Body.String(), "pong")
}

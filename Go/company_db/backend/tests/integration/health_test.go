package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var container testcontainers.Container
var router *gin.Engine

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	TestRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"ok"`)
}

package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupVersionTestRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(testRequestLogger())
	r.Use(testResponseLogger())

	r.GET("/version", VersionHandler)
	return r
}

func TestVersionRoute(t *testing.T) {
	Version = "My Version"
	router := setupVersionTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"Version\":\"My Version\"}", w.Body.String())
}

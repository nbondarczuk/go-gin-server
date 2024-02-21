package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health provides simple feedback about living server.
func Health(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

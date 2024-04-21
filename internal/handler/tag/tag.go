package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

func ReadHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

func UpdateHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

func DeleteHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

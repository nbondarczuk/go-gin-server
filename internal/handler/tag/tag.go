package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/controller/entity"
)

// CreateHandler makes a new resource with give attributes.
func CreateHandler(c *gin.Context) {
	var tag entity.Tag
	// Check input ie. new object attributes from request body.
	if err := ShouldBindJSON(&tag); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagController()
	if err != nil {
		// Handle error in controller allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.Create(&tag)
	if err != nil {
		// Handle error in object creation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"object": rval,
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

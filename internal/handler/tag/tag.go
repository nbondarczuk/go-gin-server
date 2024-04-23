package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/controller/entity"
)

// CreateHandler makes a new resource with given attributes.
func CreateHandler(c *gin.Context) {
	var tag entity.Tag
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&tag); err != nil {
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
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHandler reaks one or all records from the controller.
func ReadHandler(c *gin.Context) {
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagController()
	if err != nil {
		// Handle error in controller allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	var rval any
	// Dispatch in case the id was provided in the request.
	if id != "" {
		rval, err = tc.ReadOne(id)
	} else {
		rval, err = tc.ReadAll()
	}
	if err != nil {
		// Handle error in controller read operation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// UpdateHandler replaces lla attributes of a given object.
func UpdateHandler(c *gin.Context) {
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagController()
	if err != nil {
		// Handle error in controller allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		// Handle error in request without a parameter.
		c.JSON(http.StatusBadRequest,
			gin.H{"error": ErrEmptyTagId})
		return
	}
	var tag entity.Tag
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&tag); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	err = tc.Update(id, &tag)
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

// DeleteHandler removes an object from backend.
func DeleteHandler(c *gin.Context) {
	// The controller gives access to particular collection.
	tc, err := entity.NewTagController()
	if err != nil {
		// Handle error in controller allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		// Handle error in request without a parameter.
		c.JSON(http.StatusBadRequest,
			gin.H{"error": ErrEmptyTagId})
		return
	}
	err = tc.Delete(id)
	if err != nil {
		// Handle error in object deleteing.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

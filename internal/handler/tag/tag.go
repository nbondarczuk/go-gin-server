package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/repository/entity"
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
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
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

// ReadHOneandler reaks one or all records from the repository.
func ReadOneHandler(c *gin.Context) {
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	rval, err := tc.ReadOne(id)
	if err != nil {
		// Handle error in repository read operation.
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

// ReadHandler reads one or all records from the repository.
func ReadHandler(c *gin.Context) {
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.Read()
	if err != nil {
		// Handle error in repository read operation.
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
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
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
	err = tc.UpdateOne(id, &tag)
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

// DeleteHandler removes an object from backend.
func DeleteHandler(c *gin.Context) {
	// The repository gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
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
	err = tc.DeleteOne(id)
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

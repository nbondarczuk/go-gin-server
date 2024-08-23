package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/cache/entity/tag"
	"go-gin-server/internal/repository/entity"
)

// CreateHandler creates a new tag with given attributes.
//
// swagger:operation POST /api/entity/tag tag CreateHandler
// Creates new tag in the repository.
// ---
// produces:
//  - application/json
//
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '500':
//	   description: Internal Server Error
//
func CreateHandler(c *gin.Context) {
	var tag entity.Tag
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&tag); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// The controller gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.Create(&tag)
	if err != nil {
		// Handle error in object creation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHOneHandler reads one tag from repository.
//
// swagger:operation GET /api/entity/tag/{id} tag ReadOneHandler
// Reads one one tag from repository.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the tag
//     required: true
//     type: string
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func ReadOneHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty tag id provided"})
	}
	cache, err := tag.NewTagCache()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rval, found, err := cache.Check(id)
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		// The controlle gives access to particular collection.
		tc, err := entity.NewTagRepository()
		if err != nil {
			// Handle error in repository allocation.
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rval, err = tc.ReadOne(id)
		if err != nil {
			// Handle error in repository read operation.
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// UpdateHandler replaces some attributes of a given object.
//
// swagger:operation PATCH /api/entity/tag/{id} tag UpdateHandler
// Reads one one tag from repository.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the tag
//     required: true
//     type: string
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func UpdateHandler(c *gin.Context) {
	cache, err := tag.NewTagCache()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		// Handle error in request without a parameter.
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyTagId})
		return
	}
	var tag entity.Tag
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&tag); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = cache.Flush(id)
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = tc.UpdateOne(id, &tag)
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

// DeleteHandler removes tag from backend.
//
// swagger:operation DELETE /api/entity/tag/{id} tag DeleteHandler
// Deletes one tag from repository.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the tag
//     required: true
//     type: string
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func DeleteHandler(c *gin.Context) {
	cache, err := tag.NewTagCache()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// The repository gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		// Handle error in request without a parameter.
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyTagId})
		return
	}
	err = cache.Flush(id)
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = tc.DeleteOne(id)
	if err != nil {
		// Handle error in object deleteing.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHandler reads some tags from the repository.
//
// swagger:operation GET /api/entity/tags tag ReadHandler
// Reads some tags from repository.
// ---
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func ReadHandler(c *gin.Context) {
	// The controlle gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.Read()
	if err != nil {
		// Handle error in repository read operation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// DropHandler removes collection from backend.
//
// swagger:operation DELETE /api/entity/tags tag DropHandler
// Drops collection from repository.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the tag
//     required: true
//     type: string
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func DropHandler(c *gin.Context) {
	cache, err := tag.NewTagCache()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// The repository gives access to particular collection.
	tc, err := entity.NewTagRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = cache.Purge()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = tc.Drop()
	if err != nil {
		// Handle error in dropping the collection.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}

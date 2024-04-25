package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Version string

func SetVersion(version string) {
	Version = version
}

// Version provides info about version
func VersionHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Version": Version,
	}
	c.JSON(http.StatusOK, r)
	return
}

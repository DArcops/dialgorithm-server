package v1

import (
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course models.Course

	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}

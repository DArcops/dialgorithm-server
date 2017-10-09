package v1

import (
	"net/http"
	"strings"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course models.Course

	if err := c.BindJSON(&course); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	name := strings.ToLower(course.Name)

	if !models.First(&models.Course{}, "name = ?", name).RecordNotFound() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "This name of course can't be repeated",
		})
		return
	}

	if err := course.Add(); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{})
	return
}

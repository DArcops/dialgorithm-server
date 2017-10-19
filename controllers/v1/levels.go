package v1

import (
	"net/http"
	"strconv"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func GetLevels(c *gin.Context) {

	courseID := c.Query("course_id")
	formatID, err := strconv.ParseInt(courseID, 10, 32)

	if err != nil {
		Respond(http.StatusNotFound, gin.H{"error": "invalid course"}, c)
		return
	}

	levels, err := models.GetLevels(int(formatID))
	Respond(http.StatusOK, levels, c)
	return
}

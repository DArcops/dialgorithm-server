package v1

import (
	"fmt"
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func AddLesson(c *gin.Context) {
	var requestLesson models.RequestNewLesson

	user := c.MustGet("user").(models.User)

	if !user.CanWrite {
		Respond(http.StatusForbidden, gin.H{}, c)
		return
	}

	if err := c.BindJSON(&requestLesson); err != nil {
		fmt.Println("que pedito", err)
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	if err := requestLesson.Add(); err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}

	Respond(http.StatusOK, gin.H{}, c)
	return
}

func GetLessons(c *gin.Context) {
	levelID := c.Query("level_id")

	lessons, err := models.GetLessons(levelID)
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}

	Respond(http.StatusOK, lessons, c)
	return
}

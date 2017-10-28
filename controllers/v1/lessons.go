package v1

import (
	"fmt"
	"net/http"
	"strconv"

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
	pagination := c.Query("pagination")
	courseID := c.Query("course_id")

	var last int64

	level := models.Level{}
	if models.First(&level, "id = ? and course_id = ?", levelID, courseID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	if pagination == "true" {
		last, _ = strconv.ParseInt(c.Query("last"), 10, 32)
	} else {
		last = -1
	}

	lessons, err := models.GetLessons(level.ID, last)
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}
	Respond(http.StatusOK, lessons, c)
	return
}

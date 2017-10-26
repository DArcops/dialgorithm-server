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

	lessons, err := models.GetLessons(levelID)
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}
	Respond(http.StatusOK, lessons, c)
	return
}

func GetRelatedLessons(c *gin.Context) {

	user := c.MustGet("user").(models.User)

	course_id := c.Query("course_id")
	level_number := c.Query("level_number")

	if models.First(&models.Subscription{}, "course_id = ? and user_id = ?", course_id, user.ID).RecordNotFound() {
		Respond(http.StatusForbidden, gin.H{"error": "You are not in this course"}, c)
		return
	}

	level := models.Level{}
	if models.First(&level, "course_id = ? and number = ?", course_id, level_number).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	lessons, err := models.GetLessons(strconv.FormatUint(uint64(level.ID), 10))
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}
	Respond(http.StatusOK, lessons, c)
	return
}

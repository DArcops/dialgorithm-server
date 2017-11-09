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

	level := c.MustGet("level").(models.Level)

	lessons, err := level.GetLessons()
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}
	Respond(http.StatusOK, lessons, c)
	return
}

func GetLesson(c *gin.Context) {
	lessonID := c.Param("lesson_id")
	lesson := models.Lesson{}

	if models.First(&lesson, "id = ?", lessonID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}
	err := lesson.FillCode()
	if err != nil {
		Respond(Err[err], err, c)
		return
	}
	Respond(http.StatusOK, lesson, c)
	return
}

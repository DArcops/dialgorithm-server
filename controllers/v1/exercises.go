package v1

import (
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func AddExercise(c *gin.Context) {
	var requestExercise models.RequestNewExercise

	user := c.MustGet("user").(models.User)

	if !user.CanWrite {
		Respond(http.StatusForbidden, gin.H{}, c)
		return
	}

	if err := c.BindJSON(&requestExercise); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	if err := requestExercise.Add(); err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}

	Respond(http.StatusOK, gin.H{}, c)
	return
}

func GetExercises(c *gin.Context) {
	lessonID := c.Query("lesson_id")

	lesson := models.Lesson{}
	if models.First(&lesson, "id = ?", lessonID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	exercises, err := lesson.GetExercises()
	if err != nil {
		Respond(Err[err], err, c)
		return
	}
	Respond(http.StatusOK, exercises, c)
	return
}

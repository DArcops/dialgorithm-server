package v1

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetExercise(c *gin.Context) {
	lessonID := c.Query("lesson_id")

	lesson := models.Lesson{}
	if models.First(&lesson, "id = ?", lessonID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	exerciseID := c.Param("exercise_id")
	exrID, _ := strconv.ParseUint(exerciseID, 10, 32)

	exercise, err := lesson.GetExercise(uint(exrID))
	if err != nil {
		Respond(Err[err], err, c)
		return
	}

	user := c.MustGet("user").(models.User)

	if user.CanWrite && c.Query("to_update") == "true" {
		response := exercise.FillIO()
		Respond(http.StatusOK, response, c)
		return
	}

	Respond(http.StatusOK, exercise, c)
	return
}

func TestSolution(c *gin.Context) {
	var solution models.Exercise

	lessonID := c.Query("lesson_id")

	lesson := models.Lesson{}
	if models.First(&lesson, "id = ?", lessonID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	exerciseID := c.Param("exercise_id")
	exrID, _ := strconv.ParseUint(exerciseID, 10, 32)

	exercise, err := lesson.GetExercise(uint(exrID))
	if err != nil {
		Respond(Err[err], err, c)
		return
	}

	if err := c.BindJSON(&solution); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	response := exercise.TestSolution(solution.Code)
	fmt.Println("RESPUESTA!", response)
	Respond(http.StatusOK, response, c)
	return
}

func Solve(c *gin.Context) {
	var solution models.Exercise

	lessonID := c.Query("lesson_id")
	user := c.MustGet("user").(models.User)

	lesson := models.Lesson{}
	if models.First(&lesson, "id = ?", lessonID).RecordNotFound() {
		Respond(http.StatusNotFound, gin.H{}, c)
		return
	}

	exerciseID := c.Param("exercise_id")
	exrID, _ := strconv.ParseUint(exerciseID, 10, 32)

	exercise, err := lesson.GetExercise(uint(exrID))
	if err != nil {
		Respond(Err[err], err, c)
		return
	}

	if err := c.BindJSON(&solution); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	courseID, _ := strconv.ParseUint(c.Query("course_id"), 10, 32)
	intLessonID, _ := strconv.ParseUint(lessonID, 10, 32)

	compiled, status := exercise.Solve(solution.Code, user.ID, uint(courseID), uint(intLessonID))
	Respond(http.StatusOK, gin.H{
		"output": compiled,
		"status": status,
	}, c)
	return

}

func UpdateExercise(c *gin.Context) {
	var requestExercise models.RequestNewExercise

	user := c.MustGet("user").(models.User)

	if !user.CanWrite {
		Respond(http.StatusForbidden, gin.H{}, c)
		return
	}

	if err := c.BindJSON(&requestExercise); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		fmt.Println("ioad", err)
		return
	}

	exerciseID := c.Param("exercise_id")
	exercise := models.Exercise{}
	models.First(&exercise, "id = ?", exerciseID)

	if err := exercise.Update(requestExercise); err != nil {
		Respond(http.StatusInternalServerError, err, c)
		return
	}

	Respond(http.StatusOK, "updated", c)
	return
}

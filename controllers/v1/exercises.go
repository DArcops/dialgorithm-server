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

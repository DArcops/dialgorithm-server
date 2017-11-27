package v1

import (
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func GetSolutions(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	if !user.CanWrite {
		Respond(http.StatusForbidden, "You dont have enough permissions", c)
		return
	}

	countExercises := models.GetSolutions()
	Respond(http.StatusOK, countExercises, c)
	return
}

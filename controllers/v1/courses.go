package v1

import (
	"fmt"
	"net/http"
	"strings"

	b64 "encoding/base64"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/darcops/dialgorithm-server/modules/encrypt"
	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course models.Course

	user := c.MustGet("user").(models.User)
	fmt.Println("Usuario identificado", user)

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

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if len(token) == 0 {
			RespondWithError(http.StatusUnauthorized, "login required", c)
			return
		}

		decoded, err := b64.StdEncoding.DecodeString(token)
		if err != nil {
			RespondWithError(http.StatusUnauthorized, "login required", c)
			return
		}
		userInfo, err := encrypt.Decrypt(decoded)
		if err != nil {
			RespondWithError(http.StatusUnauthorized, "login required", c)
			return
		}

		data := strings.Split(string(userInfo), "+")

		if len(data) < 2 {
			RespondWithError(http.StatusUnauthorized, "login required", c)
			return
		}

		user := models.User{}
		models.First(&user, "id = ?", data[1])
		c.Set("user", user)
		c.Next()
	}
}

func RespondWithError(code int, message string, c *gin.Context) {
	response := map[string]string{"error: ": message}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(code, response)
	c.AbortWithStatus(code)
}

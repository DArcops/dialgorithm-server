package v1

import (
	"net/http"
	"strconv"
	"strings"

	b64 "encoding/base64"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/darcops/dialgorithm-server/modules/encrypt"
	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course models.Course

	user := c.MustGet("user").(models.User)

	if !user.CanWrite {
		Respond(http.StatusForbidden, gin.H{}, c)
		return
	}

	if err := c.BindJSON(&course); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	name := strings.ToLower(course.Name)

	if !models.First(&models.Course{}, "name = ?", name).RecordNotFound() {

		Respond(http.StatusConflict, gin.H{
			"error": "This name of course can't be repeated",
		}, c)
		return
	}

	if err := course.Add(); err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}

	Respond(http.StatusOK, gin.H{}, c)
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

func GetCourses(c *gin.Context) {
	last := c.Query("last")
	elem, err := strconv.ParseInt(last, 10, 32)
	if err != nil {
		elem = 0
	}
	courses, err := models.GetCourses(int(elem))
	if err != nil {
		Respond(http.StatusInternalServerError, gin.H{}, c)
		return
	}
	Respond(http.StatusOK, courses, c)
	return
}

func RespondWithError(code int, message string, c *gin.Context) {
	response := map[string]string{"error: ": message}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(code, response)
	c.AbortWithStatus(code)
}

func Respond(code int, data interface{}, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(code, data)
}

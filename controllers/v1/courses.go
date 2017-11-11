package v1

import (
	"fmt"
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

func CourseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID := c.Query("course_id")
		course := models.Course{}
		if models.First(&course, "id = ?", courseID).RecordNotFound() {
			RespondWithError(http.StatusNotFound, "Course not found", c)
			return
		}

		user := c.MustGet("user").(models.User)
		if models.First(&models.Subscription{}, "user_id = ? and course_id = ?", user.ID, course.ID).RecordNotFound() && !user.CanWrite {
			RespondWithError(http.StatusForbidden, "you dont have any susbciption to this course", c)
			return
		}
		c.Set("course", course)
		c.Next()
	}
}

func LevelMiddleware() gin.HandlerFunc {
	fmt.Println("entyra al middleware de level")

	return func(c *gin.Context) {
		levelID := c.Query("level_id")
		course := c.MustGet("course").(models.Course)
		level := models.Level{}
		if models.First(&level, "id = ? and course_id = ?", levelID, course.ID).RecordNotFound() {
			RespondWithError(http.StatusNotFound, "", c)
			return
		}
		c.Set("level", level)
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

func SuscribeToCourse(c *gin.Context) {
	var subscription models.Subscription

	user := c.MustGet("user").(models.User)

	if err := c.BindJSON(&subscription); err != nil {
		Respond(http.StatusBadRequest, gin.H{}, c)
		return
	}

	if subscription.UserPass != user.Password {
		Respond(http.StatusForbidden, "Invalid password", c)
		return
	}

	subscription.UserID = user.ID

	if err := subscription.Add(); err != nil {
		Respond(Err[err], err, c)
		return
	}

	Respond(http.StatusCreated, gin.H{}, c)
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

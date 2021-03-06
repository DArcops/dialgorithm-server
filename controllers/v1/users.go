package v1

import (
	"fmt"
	"net/http"
	"strconv"

	b64 "encoding/base64"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	fmt.Println("aqui se devolveran los usuarios")
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := models.Create(&user).Error; err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusCreated, gin.H{})
	return
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if models.First(&user, "email = ? and password = ?", user.Email, user.Password).RecordNotFound() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	token, err := models.GenerateToken([]byte(user.Email + "+" + strconv.FormatUint(uint64(user.ID), 10)))

	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"token":      b64.StdEncoding.EncodeToString(token),
		"user_name":  user.Name,
		"user_email": user.Email,
	})
	return

}

func GetProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	user.Password = ""
	Respond(http.StatusOK, user, c)
	return
}

func GetUserCourses(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	courses, err := user.GetCourses()
	if err != nil {
		Respond(http.StatusInternalServerError, err, c)
		return
	}
	Respond(http.StatusOK, courses, c)
	return
}

func AddAdmin(c *gin.Context) {
	var newAdmin models.NewAdmin

	if err := c.Bind(&newAdmin); err != nil {
		Respond(http.StatusBadRequest, "", c)
		return
	}

	user := c.MustGet("user").(models.User)
	if !user.CanWrite {
		Respond(http.StatusForbidden, "you dont have enough permissions", c)
		return
	}

	if err := user.AddNewAdmin(newAdmin.Email); err != nil {
		Respond(Err[err], err, c)
		return
	}
	Respond(http.StatusCreated, "created", c)
	return
}

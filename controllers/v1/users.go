package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

var user models.User

func GetUsers(c *gin.Context) {
	fmt.Println("aqui se devolveran los usuarios")
}

func Register(c *gin.Context) {

	if err := c.Bind(&user); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	fmt.Println("que pedoo", user)

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

	token, err := models.GenerateToken([]byte(gocql.TimeUUID().String() + "+" + strconv.FormatUint(uint64(user.ID), 10)))

	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusInternalServerError, gin.H{})
		fmt.Println("hubo pedo", err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	fmt.Println("HHSGSSS", token)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return

}

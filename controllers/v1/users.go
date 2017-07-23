package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	fmt.Println("aqui se devolveran los usuarios")
}

package routes

import (
	V1 "github.com/darcops/dialgorithm-server/controllers/v1"
	"github.com/gin-gonic/gin"
)

var users *gin.RouterGroup

func userRoutes() {
	users = v1.Group("users")
	users.GET("", V1.GetUsers)
	users.POST("/register", V1.Register)
	users.POST("/login", V1.Login)
	users.GET("/profile", V1.UserMiddleware(), V1.GetProfile)
	users.GET("/courses", V1.UserMiddleware(), V1.GetUserCourses)
}

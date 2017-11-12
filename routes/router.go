package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var (
	router *gin.Engine
	api    *gin.RouterGroup
	v1     *gin.RouterGroup
)

func Init() {
	router = gin.New()
	api = router.Group("api")
	v1 = api.Group("v1")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,DELETE,PUT",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	userRoutes()
	courseRoutes()
	lessonRoutes()
	levelsRoutes()
	exerciseRoutes()
	solutionRoutes()

	router.Run(":8088")
}

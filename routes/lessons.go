package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func lessonRoutes() {
	courses := v1.Group("lessons").Use(V1.UserMiddleware())
	courses.POST("/new", V1.AddLesson)
}

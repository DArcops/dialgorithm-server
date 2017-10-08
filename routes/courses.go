package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func courseRoutes() {
	courses := v1.Group("courses")
	courses.POST("/new", V1.AddCourse)
}

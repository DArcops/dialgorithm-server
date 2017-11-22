package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func courseRoutes() {
	courses := v1.Group("courses").Use(V1.UserMiddleware())
	courses.POST("/new", V1.AddCourse)
	courses.POST("/suscribe", V1.SuscribeToCourse)
	courses.GET("/", V1.GetCourses)
	courses.GET("/:course_id", V1.GetCourse)
	courses.POST("/update/:course_id", V1.UpdateCourse)
}

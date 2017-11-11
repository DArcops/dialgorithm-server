package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func exerciseRoutes() {
	exercises := v1.Group("exercises").Use(V1.UserMiddleware())
	exercises.POST("/new", V1.AddExercise)
	exercises.GET("", V1.CourseMiddleware(), V1.GetExercises)
	exercises.GET("/:exercise_id", V1.CourseMiddleware(), V1.GetExercise)
}

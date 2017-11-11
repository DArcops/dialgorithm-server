package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func lessonRoutes() {
	lessons := v1.Group("lessons").Use(V1.UserMiddleware())
	lessons.POST("/new", V1.AddLesson)
	lessons.GET("", V1.CourseMiddleware(), V1.LevelMiddleware(), V1.LevelMiddleware(), V1.GetLessons)
	lessons.GET("/:lesson_id", V1.GetLesson)
}

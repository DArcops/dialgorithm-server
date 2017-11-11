package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func levelsRoutes() {
	levels := v1.Group("levels").Use(V1.UserMiddleware()).Use(V1.CourseMiddleware())
	levels.GET("", V1.GetLevels)
}

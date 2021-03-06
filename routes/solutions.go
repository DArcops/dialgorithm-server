package routes

import V1 "github.com/darcops/dialgorithm-server/controllers/v1"

func solutionRoutes() {
	solutions := v1.Group("solutions").Use(V1.UserMiddleware())
	{
		solutions.POST("/:exercise_id", V1.CourseMiddleware(), V1.TestSolution)
		solutions.POST("/:exercise_id/solve", V1.CourseMiddleware(), V1.Solve)
		solutions.GET("", V1.GetSolutions)
	}

}

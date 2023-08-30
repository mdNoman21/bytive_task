package routes

import (
	"bytive-task/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/createProject", controllers.CreateProject())
	incomingRoutes.GET("/getProjects", controllers.GetProjects())
	incomingRoutes.GET("/getProject/:id", controllers.GetProject())
	incomingRoutes.DELETE("/deleteProject/:id", controllers.DeleteProject())
	incomingRoutes.PATCH("/updateEndTimeAll", controllers.UpdateAllProjects())

}

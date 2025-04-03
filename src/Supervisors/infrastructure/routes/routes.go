package routes

import (
	"github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/dependencies"
	"github.com/BryanChanona/backend_users/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/supervisors")
	registerSupervisorController := dependencies.GetSaveSupervisorController().Execute
	deleteSupervisorController := dependencies.GetDeleteSupervisorController().Execute

	routes.POST("/",middlewares.AuthMiddleware(),registerSupervisorController)
	routes.DELETE("/:id_supervisor",middlewares.AuthMiddleware(),deleteSupervisorController)

}
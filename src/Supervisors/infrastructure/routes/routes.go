package routes

import (
	"github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/supervisors")
	registerSupervisorController := dependencies.GetSaveSupervisorController().Execute

	routes.POST("/",registerSupervisorController)

}
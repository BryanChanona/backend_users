package routes

import (
	"github.com/BryanChanona/backend_users/src/Mqtt/infrastructure/adapters"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/publish")


	publishHandler := adapters.PublishHandler


	routes.POST("/",publishHandler)

}
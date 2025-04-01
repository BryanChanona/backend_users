package routes

import (

	"github.com/BryanChanona/backend_users/src/User/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/users")

	saveUserController := dependencies.GetSaveUserController().Execute
	loginController := dependencies.GeyLoginController().Execute
	

	routes.POST("/", saveUserController)
	routes.POST("/login",loginController)

}
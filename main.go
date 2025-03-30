package main

import (
	"github.com/BryanChanona/backend_users/src/Mqtt/infrastructure/adapters"
	"github.com/BryanChanona/backend_users/src/User/infrastructure/dependencies"
	"github.com/BryanChanona/backend_users/src/User/infrastructure/routes"
	"github.com/gin-gonic/gin"
	routesMQTT "github.com/BryanChanona/backend_users/src/Mqtt/infrastructure/routes"

	
)

func main() {
	adapters.InitMQTT()
	dependencies.Init()

	r := gin.Default()

	routes.Routes(r)
	routesMQTT.Routes(r)

	r.Run(":8080")
}
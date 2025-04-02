package main

import (
	//"github.com/BryanChanona/backend_users/src/Mqtt/infrastructure/adapters"
	dependenciesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/dependencies"
	dependenciesSupervisor "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/dependencies"
	routesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/routes"
	"github.com/gin-gonic/gin"
	routesMQTT "github.com/BryanChanona/backend_users/src/Mqtt/infrastructure/routes"
	routesSupervisors "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/routes"

	
)

func main() {
	//adapters.InitMQTT()
	dependenciesUser.Init()
	dependenciesSupervisor.Init()

	r := gin.Default()

	routesUser.Routes(r)
	routesSupervisors.Routes(r)

	routesMQTT.Routes(r)

	r.Run(":8080")
}
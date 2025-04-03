package main

import (
	dependenciesSupervisor "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/dependencies"
	routesSupervisors "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/routes"
	//"github.com/BryanChanona/backend_users/src/User/infrastructure/adapters"
	dependenciesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/dependencies"
	routesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//adapters.InitMQTT()
	dependenciesUser.Init()
	dependenciesSupervisor.Init()

	r := gin.Default()

	routesUser.Routes(r)
	routesSupervisors.Routes(r)


	r.Run(":8080")
}
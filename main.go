package main

import (
	dependenciesSupervisor "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/dependencies"
	routesSupervisors "github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/routes"
	"github.com/BryanChanona/backend_users/src/User/infrastructure/adapters"
	dependenciesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/dependencies"
	routesUser "github.com/BryanChanona/backend_users/src/User/infrastructure/routes"
	"github.com/BryanChanona/backend_users/src/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	adapters.InitMQTT()
	dependenciesUser.Init()
	dependenciesSupervisor.Init()
	

	r := gin.Default()
	helpers.InitCORS(r)
	routesUser.Routes(r)
	routesSupervisors.Routes(r)
	


	r.Run(":8080")
}
package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_users/src/Supervisors/application/UseCase"
	"github.com/BryanChanona/backend_users/src/Supervisors/infrastructure"
	"github.com/BryanChanona/backend_users/src/Supervisors/infrastructure/controllers"
	"github.com/BryanChanona/backend_users/src/helpers"
)

var (
	mySQL infrastructure.MySQL
)

func Init() {
	db, err := helpers.ConnMySQL()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL = *infrastructure.NewMySQL(db)

}

func GetSaveSupervisorController() *controllers.RegisterSupervisorController {
	useCaseSaveSupervisor := UseCase.NewRegisterSupervisorUc(&mySQL)
	return controllers.NewRegisterSupervisorController(useCaseSaveSupervisor)
}
func GetDeleteSupervisorController()*controllers.DeleteSupervisorController{
	useCase := UseCase.NewDeleteSupervisorUc(&mySQL)
	return controllers.NewDeleteSupervisorController(useCase)
}
func GetGetSupervisorsController()*controllers.GetSupervisorsByUserController{
	useCase := UseCase.NewGetSupervisorByUserUC(&mySQL)
	return controllers.NewGetSupervisorsByUserController(useCase)
}


package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_users/src/User/application/UseCase"
	"github.com/BryanChanona/backend_users/src/User/infrastructure"
	"github.com/BryanChanona/backend_users/src/User/infrastructure/controllers"
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

func GetSaveUserController() *controllers.RegisterUsercontroller{
	useCaseSaveUser := UseCase.NewRegisterUserUC(&mySQL)
	return controllers.NewRegisterUserController(*useCaseSaveUser)
}
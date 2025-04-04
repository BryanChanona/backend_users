package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_users/src/Supervisors/application/UseCase"
	"github.com/BryanChanona/backend_users/src/Supervisors/domain"
	"github.com/BryanChanona/backend_users/src/helpers"
	"github.com/gin-gonic/gin"
)

type LogInController struct {
	useCase *UseCase.LoginUseCase
}
func NewLogInController(useCase *UseCase.LoginUseCase) *LogInController {
	return &LogInController{useCase: useCase}
}

func (controller *LogInController)Execute (ctx *gin.Context){
	var supervisor domain.SupervisorsModel
	if err := ctx.ShouldBindJSON(&supervisor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	email :=supervisor.Email
	password := supervisor.Password

	authenticatedSupervisor, err := controller.useCase.Execute(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	if authenticatedSupervisor.Id_supervisor == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de supervisor"})
		return
	}

	token, err := helpers.GenerateSupervisorJWT(authenticatedSupervisor.Id_supervisor,authenticatedSupervisor.Id_usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"supervisor": authenticatedSupervisor, "token": token})
}
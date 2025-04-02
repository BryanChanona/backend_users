package controllers

import (
	"net/http"
	"strconv"

	"github.com/BryanChanona/backend_users/src/Supervisors/application/UseCase"
	"github.com/BryanChanona/backend_users/src/Supervisors/domain"
	"github.com/gin-gonic/gin"
)

type RegisterSupervisorController struct {
	useCase *UseCase.RegisterSupervisorUC
}

func NewRegisterSupervisorController(useCase *UseCase.RegisterSupervisorUC) *RegisterSupervisorController {
	return &RegisterSupervisorController{useCase: useCase}
}

func (controller *RegisterSupervisorController) Execute(ctx *gin.Context) {
	var supervisor domain.SupervisorsModel

	// Obtener el idUser desde la URL y validar que sea un número
	idUser, err := strconv.Atoi(ctx.Param("idUser"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	// Validar el JSON recibido
	if err := ctx.ShouldBindJSON(&supervisor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Asignar el ID de usuario al supervisor
	supervisor.Id_usuario = idUser

	// Ejecutar el caso de uso para registrar el supervisor
	if err := controller.useCase.Execute(supervisor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Supervisor registrado con éxito"})
}

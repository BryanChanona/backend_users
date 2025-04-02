package controllers

import (
	"net/http"

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

	// Obtener el ID del usuario logueado (ejemplo: de un token JWT)
	userID, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
		return
	}

	// Asegurarse de que userID es de tipo int
	idUser, ok := userID.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	// Validar datos recibidos
	if err := ctx.ShouldBindJSON(&supervisor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos"})
		return
	}

	// Asignar el ID del usuario al supervisor
	supervisor.Id_usuario = idUser

	// Ejecutar caso de uso
	if err := controller.useCase.Execute(supervisor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "supervisor registrado exitosamente"})
}
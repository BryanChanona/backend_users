package controllers

import (
	"net/http"
	"strconv"

	"github.com/BryanChanona/backend_users/src/User/application/UseCase"
	"github.com/BryanChanona/backend_users/src/User/domain"
	"github.com/gin-gonic/gin"
)

type UpdateStatusController struct {
	useCase *UseCase.UpdateStatusUC
}

func NewUpdateStatusUcController(useCase *UseCase.UpdateStatusUC) *UpdateStatusController {
	return &UpdateStatusController{
		useCase: useCase,
	}
}

func (controller *UpdateStatusController) Execute(ctx *gin.Context) {
	var user domain.User

	// Convertir el ID del usuario desde la URL
	idUser, err := strconv.Atoi(ctx.Param("idUser"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Validar los datos del JSON
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	status := user.Premium

	// Llamar al caso de uso
	err = controller.useCase.Execute(idUser, status)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el usuario con ese ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Actualizado con éxito"})
}

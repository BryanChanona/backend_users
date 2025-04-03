package controllers

import (
	"net/http"
	"strconv"

	"github.com/BryanChanona/backend_users/src/Supervisors/application/UseCase"
	"github.com/gin-gonic/gin"
)

type DeleteSupervisorController struct {
	useCase *UseCase.DeleteSupervisorUC
}

func NewDeleteSupervisorController(useCase *UseCase.DeleteSupervisorUC) *DeleteSupervisorController {
	return &DeleteSupervisorController{
		useCase: useCase,
	}
}

func (controller *DeleteSupervisorController) Execute(ctx *gin.Context) {
	userID, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
		return
	}

	idUser, ok := userID.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	idSupervisor, err := strconv.Atoi(ctx.Param("id_supervisor"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de supervisor no válido"})
		return
	}

	err = controller.useCase.Execute(idUser,idSupervisor)
	if err != nil {
		if err.Error() == "no existe el supervisor" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Supervisor no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{"message": "Supervisor deleted successfully"})
}

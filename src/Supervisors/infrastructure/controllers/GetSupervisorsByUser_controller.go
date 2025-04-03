package controllers

import (
	"github.com/BryanChanona/backend_users/src/Supervisors/application/UseCase"
	"github.com/gin-gonic/gin"
)

type GetSupervisorsByUserController struct {
	useCase *UseCase.GetSupervisorsByUserUC
}

func NewGetSupervisorsByUserController(useCase *UseCase.GetSupervisorsByUserUC) *GetSupervisorsByUserController {
	return &GetSupervisorsByUserController{
		useCase: useCase,
	}
}

func (controller *GetSupervisorsByUserController)Execute (ctx *gin.Context){
	idUser, exists := ctx.Get("id_user")
	if !exists{
		ctx.JSON(400, gin.H{"error": "id_user not found in context"})
		return
	}
	supervisors, err := controller.useCase.Execute(idUser.(int))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error retrieving supervisors"})
		return
	}
	ctx.JSON(200, gin.H{"supervisors": supervisors})
}

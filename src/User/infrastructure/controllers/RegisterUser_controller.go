package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_users/src/User/application/UseCase"
	"github.com/BryanChanona/backend_users/src/User/domain"
	"github.com/gin-gonic/gin"
)

type RegisterUsercontroller struct {
	useCase UseCase.RegisterUserUC
}

func NewRegisterUserController(useCase UseCase.RegisterUserUC) *RegisterUsercontroller{
	return &RegisterUsercontroller{useCase: useCase}
}

func (controller *RegisterUsercontroller)Execute (ctx *gin.Context){
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the use case to create the book
	err := controller.useCase.Execute(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Registered user"})
	}
}
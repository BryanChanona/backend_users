package controllers

import (
	"fmt"
	"net/http"

	"github.com/BryanChanona/backend_users/src/User/application/UseCase"
	"github.com/BryanChanona/backend_users/src/User/domain"
	"github.com/BryanChanona/backend_users/src/helpers"
	"github.com/gin-gonic/gin"
)
type LogInController struct {
	useCase *UseCase.LogInUc
}

func NewLogInController (useCase *UseCase.LogInUc)*LogInController{
	return &LogInController{useCase: useCase}
}

func (controller *LogInController)Execute(ctx *gin.Context){
	var user domain.User

	//Obtenemos los datos del Json de la solicitud
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	email := user.Email
	password := user.Password	
	
	authenticatedUser, err := controller.useCase.Execute(email,password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}
	token, err:= helpers.GenerateJWT(email)
	if err !=nil {
		fmt.Println("Error al generar el token")
	}

	
	ctx.JSON(http.StatusOK, gin.H{"user": authenticatedUser,"token":token})

}
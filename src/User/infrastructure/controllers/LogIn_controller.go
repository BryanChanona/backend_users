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

func NewLogInController(useCase *UseCase.LogInUc) *LogInController {
	return &LogInController{useCase: useCase}
}

func (controller *LogInController) Execute(ctx *gin.Context) {
	var user domain.User

	// Obtenemos los datos del JSON de la solicitud
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	email := user.Email
	password := user.Password	

	// Autenticamos al usuario
	authenticatedUser, err := controller.useCase.Execute(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Aseguramos que authenticatedUser tenga un ID válido
	if authenticatedUser.Id_usuario == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ID de usuario"})
		return
	}

	// Generamos el token con el ID del usuario
	token, err := helpers.GenerateJWT(authenticatedUser.Id_usuario)
	if err != nil {
		fmt.Println("Error al generar el token:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": authenticatedUser, "token": token})
}

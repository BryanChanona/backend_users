package UseCase

import (
	"fmt"

	"github.com/BryanChanona/backend_users/src/User/domain"
	"github.com/BryanChanona/backend_users/src/helpers"
)

type RegisterUserUC struct {
	db domain.IUserRepository
}

func NewRegisterUserUC(db domain.IUserRepository) *RegisterUserUC{
	return &RegisterUserUC{db: db}
}

func (useCase *RegisterUserUC) Execute(user domain.User) error{
	password := user.Password
	user.Premium = false

	hashPassword, err := helpers.EncryptPassword(password)

	if err != nil {
		fmt.Print("Hubo un error al hashear la contraseña.")
	}
	user.Password = string(hashPassword)
	
	return useCase.db.RegisterUser(user)
}
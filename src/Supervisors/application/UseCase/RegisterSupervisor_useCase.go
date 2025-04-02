package UseCase

import (
	"fmt"

	"github.com/BryanChanona/backend_users/src/Supervisors/domain"
	"github.com/BryanChanona/backend_users/src/helpers"
)

type RegisterSupervisorUC struct {
	db domain.ISupervisorRepository
}

func NewRegisterSupervisorUc(db domain.ISupervisorRepository) *RegisterSupervisorUC {
	return &RegisterSupervisorUC{
		db: db,
	}
}
func (useCase *RegisterSupervisorUC) Execute(supervisor domain.SupervisorsModel) error {
	// Verificar si el usuario es premium
	isPremium, err := useCase.db.GetUserPremiumStatus(supervisor.Id_usuario)
	if err != nil {
		return err
	}

	// Si el usuario no es premium, solo puede tener un supervisor
	if !isPremium {
		count, err := useCase.db.CountSupervisors(supervisor.Id_usuario)
		if err != nil {
			return err
		}
		if count >= 1 {
			return fmt.Errorf("el usuario ya tiene un supervisor registrado")
		}
	}

	// Verificar si el correo ya está registrado
	emailAlreadyExist, err := useCase.db.EmailAlreadyExists(supervisor.Email)
	if err != nil {
		return fmt.Errorf("error verificando el email: %w", err)
	}
	if emailAlreadyExist {
		return fmt.Errorf("el email ya está en uso")
	}

	// Encriptar la contraseña
	hashPassword, err := helpers.EncryptPassword(supervisor.Password)
	if err != nil {
		return fmt.Errorf("error al hashear la contraseña: %w", err)
	}
	supervisor.Password = string(hashPassword)

	// Registrar supervisor
	return useCase.db.RegisterSupervisor(supervisor)
}
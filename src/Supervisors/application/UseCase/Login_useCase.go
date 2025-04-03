package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type LoginUseCase struct {
	db domain.ISupervisorRepository
}

func NewLoginUseCase(db domain.ISupervisorRepository) *LoginUseCase {
	return &LoginUseCase{db: db}
}

func (useCase *LoginUseCase)Execute (email string, password string)(domain.SupervisorsModel, error){
	return useCase.db.LoginSupervisors(email,password)
}
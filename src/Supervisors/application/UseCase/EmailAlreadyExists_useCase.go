package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type EmailAlreadyExistsUC struct {
	db domain.ISupervisorRepository
}

func NewEmailAlreadyExistsUc(db domain.ISupervisorRepository) *EmailAlreadyExistsUC {
	return &EmailAlreadyExistsUC{
		db: db,
	}
}

func (useCase *EmailAlreadyExistsUC) Execute(email string) (bool, error) {
	return useCase.db.EmailAlreadyExists(email)
}
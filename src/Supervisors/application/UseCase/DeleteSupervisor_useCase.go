package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type DeleteSupervisorUC struct {
	db domain.ISupervisorRepository
}

func NewDeleteSupervisorUc(db domain.ISupervisorRepository) *DeleteSupervisorUC {
	return &DeleteSupervisorUC{
		db: db,
	}
}

func (useCase *DeleteSupervisorUC) Execute(idUser ,idSupervisor int) error {
	return useCase.db.DeleteSupervisor(idUser, idSupervisor)
}
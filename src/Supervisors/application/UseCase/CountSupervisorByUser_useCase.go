package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type CountSupervisorByUserUC struct {
	db domain.ISupervisorRepository
}

func NewCountSupervisorByUserUc(db domain.ISupervisorRepository) *CountSupervisorByUserUC {
	return &CountSupervisorByUserUC{
		db: db,
	}
}
func (useCase *CountSupervisorByUserUC)Execute(idUser int)(int, error){
	return useCase.db.CountSupervisors(idUser)
}
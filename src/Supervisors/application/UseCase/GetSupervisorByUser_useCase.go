package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type GetSupervisorsByUserUC struct {
	db domain.ISupervisorRepository
}

func NewGetSupervisorByUserUC(db domain.ISupervisorRepository) *GetSupervisorsByUserUC {
	return &GetSupervisorsByUserUC{
		db: db,
	}
}

func (uc *GetSupervisorsByUserUC) Execute(idUser int) ([]domain.SupervisorsResponse, error) {
	return uc.db.GetSupervisorsByUser(idUser)
}
package UseCase

import "github.com/BryanChanona/backend_users/src/Supervisors/domain"

type GetUserPremiumStatusUC struct {
	db domain.ISupervisorRepository
}

func NewGetUserPremiumStatusUc(db domain.ISupervisorRepository) *GetUserPremiumStatusUC {
	return &GetUserPremiumStatusUC{
		db: db,
	}
}

func (useCase *GetUserPremiumStatusUC) Execute(idUser int) (bool, error) {
	return useCase.db.GetUserPremiumStatus(idUser)
}
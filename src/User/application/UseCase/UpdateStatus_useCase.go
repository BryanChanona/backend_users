package UseCase

import "github.com/BryanChanona/backend_users/src/User/domain"

type UpdateStatusUC struct {
	db domain.IUserRepository
}

func NewUpdateStatusUc(db domain.IUserRepository) *UpdateStatusUC {
	return &UpdateStatusUC{
		db: db,
	}
}

func (useCase *UpdateStatusUC)Execute (idUser int, status bool) error {
	return useCase.db.UpdateStatus(idUser,status)
}
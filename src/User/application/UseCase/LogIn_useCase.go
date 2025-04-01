package UseCase

import "github.com/BryanChanona/backend_users/src/User/domain"

type LogInUc struct {
	db domain.IUserRepository
}

func NewLogInUc(db domain.IUserRepository) *LogInUc{
	return &LogInUc{db: db }
}
func (useCase *LogInUc) Execute(email string, password string) (domain.User, error){
	return useCase.db.LogIn(email, password)
}
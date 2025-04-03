package domain

type ISupervisorRepository interface {
	RegisterSupervisor(supervisor SupervisorsModel) error
	CountSupervisors(id_User int)(int, error)
	GetUserPremiumStatus(idUser int) (bool, error)
	EmailAlreadyExists(email string) (bool, error)
	DeleteSupervisor(idUser,idSupervisor int) error
	GetSupervisorsByUser(idUser int) ([]SupervisorsResponse, error)
	LoginSupervisors(email string, password string) (SupervisorsModel, error)
}
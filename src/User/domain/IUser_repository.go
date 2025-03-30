package domain

type IUserRepository interface{
	RegisterUser(user User)	error

}
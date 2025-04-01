package domain

type IUserRepository interface{
	RegisterUser(user User)	error
	LogIn(email string, password string) (User,error)
}
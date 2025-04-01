package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/BryanChanona/backend_users/src/User/domain"
	"golang.org/x/crypto/bcrypt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) RegisterUser(user domain.User) error {
	query, err := sql.db.Prepare("INSERT INTO `usuario` (nombre,correo,password,premium,id_dispositivo) VALUES (?,?,?,?,?)")

	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(user.Name, user.Email, user.Password, user.Premium,user.Id_device)

	if err != nil {
		log.Println("Error saving the user:", err)
		return err
	}
	return nil
}

func (sql *MySQL)LogIn(email string, password string) (domain.User,error){
	var user domain.User
	var hashedPassword string

	query:= "SELECT id_usuario, nombre, correo, password, id_dispositivo, premium FROM usuario WHERE correo = ?"
	
	err := sql.db.QueryRow(query,email).Scan(&user.Id_usuario,&user.Name,&user.Email,&hashedPassword,&user.Id_device,&user.Premium)
	if err != nil {
		return domain.User{}, fmt.Errorf("error al buscar usuario: %w", err)

	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if err != nil {
		return domain.User{}, errors.New("contraseña incorrecta")
	}
	
	user.Password = "" // Se limpia la contraseña por seguridad
	return user, nil


}
func (sql *MySQL)UpdateStatus(id_usuario int, status bool) error {
	query, err := sql.db.Prepare("UPDATE usuario SET premium = ? WHERE id_usuario = ?")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(status, id_usuario)
	if err != nil {
		log.Println("Error updating user status:", err)
		return err
	}
	return nil
}
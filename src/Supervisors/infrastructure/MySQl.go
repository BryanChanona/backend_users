package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BryanChanona/backend_users/src/Supervisors/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) RegisterSupervisor(supervisor domain.SupervisorsModel) error {
	query, err:= sql.db.Prepare("INSERT INTO supervisor (nombre,correo,password,id_usuario) VALUES (?,?,?,?)") 

	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(supervisor.Name, supervisor.Email,supervisor.Password,supervisor.Id_usuario)

	if err != nil {
		log.Println("Error guardando el supervisor:", err)
		return err
	}
	return nil
		


}

func(sql *MySQL)CountSupervisors(idUser int)(int,error){
	query, err := sql.db.Prepare("SELECT COUNT(*) FROM supervisor WHERE id_usuario = ?")
	if err != nil {
		fmt.Println("Error preparing query:", err)
	}
	defer query.Close()

	var count int 
	err = query.QueryRow(idUser).Scan(&count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}	

	return count, nil
}

func (sql *MySQL) GetUserPremiumStatus(idUser int) (bool, error) {
	query, err := sql.db.Prepare("SELECT premium FROM usuario WHERE id_usuario = ?")
	if err != nil {
		fmt.Println("Error preparing query:", err)
	}
	defer query.Close()

	var premium bool
	err = query.QueryRow(idUser).Scan(&premium)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return false, err
	}

	return premium, nil
}

func (sql *MySQL) EmailAlreadyExists(email string) (bool, error) {
	query, err := sql.db.Prepare("SELECT COUNT(*) FROM supervisor WHERE correo = ?")
	if err != nil {
		return false, fmt.Errorf("error preparando la consulta: %w", err)
	}
	defer query.Close() // Ahora solo se ejecuta si `Prepare` tiene éxito

	var count int
	err = query.QueryRow(email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	return count > 0, nil
}

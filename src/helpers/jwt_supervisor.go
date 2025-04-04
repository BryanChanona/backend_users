package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSupervisorKey = []byte(os.Getenv("JWT_SUPERVISOR_SECRET"))

// Claims específicos para supervisor
type SupervisorClaims struct {
	IdSupervisor int    `json:"id_supervisor"`
	IdUser       int    `json:"id_user"`
	jwt.StandardClaims
}

// Genera un JWT para el supervisor
func GenerateSupervisorJWT(IdSupervisor int, IdUser int) (string, error) {
	expirationTime := time.Now().Add(8 * time.Hour) // Puedes cambiar la expiración si lo deseas
	claims := &SupervisorClaims{
		IdSupervisor: IdSupervisor, // Id del supervisor
		IdUser:       IdUser,       // Id del usuario asignado al supervisor
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSupervisorKey)
}

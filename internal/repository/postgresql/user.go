package postgresql

import (
	"github.com/SunilKividor/internal/models"
	"github.com/google/uuid"
)

func GetUserQuery(id uuid.UUID) (models.User, error) {
	var user models.User
	smt := `SELECT * FROM users WHERE id = $1`

	err := db.QueryRow(smt, id).Scan(&user.Id, &user.Username, &user.Password, &user.RefreshToken, &user.Name, &user.Category)
	if err != nil {
		return user, err
	}
	return user, nil
}

func DeletUserQuery(id uuid.UUID) error {
	smt := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(smt, id)
	return err
}

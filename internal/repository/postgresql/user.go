package postgresql

import "github.com/google/uuid"

func DeletUserQuery(id uuid.UUID) error {
	smt := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(smt, id)
	return err
}

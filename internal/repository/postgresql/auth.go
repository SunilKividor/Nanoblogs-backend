package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/SunilKividor/internal/models"
	"github.com/google/uuid"
)

func GetUserDetailsQuery(username string) (models.User, error) {
	var user models.User

	smt := `SELECT * FROM users WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&user.Id, &user.Username, &user.Password, &user.AccessToken, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("failed to fetch user details: %w", err)
	}

	return user, nil
}

func GetUsernamePasswordQuery(username string) (string, string, error) {
	var (
		name     string
		password string
	)

	smt := `SELECT username,password FROM users WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&name, &password)
	if err != nil {
		return "", "", err
	}

	return name, password, nil
}

func RegisterNewUserQuery(user models.User) (uuid.UUID, error) {
	var id uuid.UUID
	smt := `INSERT INTO users(username,password,access_token,refresh_token) VALUES($1,$2,$3,$4) RETURNING id`
	err := db.QueryRow(smt, user.Username, user.Password, user.AccessToken, user.RefreshToken).Scan(&id)
	return id, err
}

func UpdateUserTokensQuery(accessToken string, refreshToken string, username string) (uuid.UUID, error) {
	var id uuid.UUID
	smt := `UPDATE users SET access_token = $1, refresh_token = $2 WHERE username = $3 RETURNING id`
	err := db.QueryRow(smt, accessToken, refreshToken, username).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return id, fmt.Errorf("no user with id %d", id)
	case err != nil:
		return id, fmt.Errorf("query error: %v", err)
	}
	return id, nil
}

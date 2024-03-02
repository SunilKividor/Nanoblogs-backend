package postgresql

import (
	"fmt"

	"github.com/SunilKividor/internal/models"
)

func GetUserDetails(username string) (models.UserDetails, error) {
	var user models.UserDetails

	smt := `SELECT * FROM user_details WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&user.Id, &user.Username, &user.Password, &user.AccessToken, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("failed to fetch user details: %w", err)
	}

	return user, nil
}

func GetUsernamePassword(username string) (string, string, error) {
	var (
		name     string
		password string
	)

	smt := `SELECT username,password FROM user_details WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&name, &password)
	if err != nil {
		return "", "", err
	}

	return name, password, nil
}

func RegisterNewUser(user models.UserDetails) error {
	smt := `INSERT INTO user_details(username,password,access_token,refresh_token) VALUES($1,$2,$3,$4)`
	_, err := db.Exec(smt, user.Username, user.Password, user.AccessToken, user.RefreshToken)

	return err
}

func UpdateUserTokens(accessToken string, refreshToken string, username string) error {

	smt := `UPDATE user_details SET access_token = $1, refresh_token = $2 WHERE username = $3`
	result, err := db.Exec(smt, accessToken, refreshToken, username)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 1 {
		return nil
	}

	return fmt.Errorf("error updating")
}

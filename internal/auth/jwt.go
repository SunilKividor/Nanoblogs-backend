package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(username string) (string, string, error) {
	privateKey := []byte(os.Getenv("APISECRET"))

	refreshClaims := jwt.MapClaims{
		"authorized": true,
		"username":   username,
		"exp":        time.Now().Add(time.Hour * 24 * time.Duration(365)).Unix(),
	}

	accessClaims := jwt.MapClaims{
		"authorized": true,
		"username":   username,
		"exp":        time.Now().Add(time.Minute * time.Duration(10)).Unix(),
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func RefreshAccessToken(refreshToken string) (string, string, error) {
	parsedToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(os.Getenv("APISECRET")), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", "", fmt.Errorf("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", "", fmt.Errorf("missing or invalid 'exp' claim")
	}
	expTime := time.Unix(int64(exp), 0)

	if time.Now().After(expTime) {
		return "", "", fmt.Errorf("token is expired")
	}

	privateKey := []byte(os.Getenv("APISECRET"))
	username := claims["username"]
	accessClaims := jwt.MapClaims{
		"authorized": true,
		"username":   username,
		"exp":        time.Now().Add(time.Minute * time.Duration(10)).Unix(),
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	return accessToken, username.(string), nil
}

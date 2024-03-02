package handlers

import (
	"fmt"
	"net/http"

	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/models"
	"github.com/SunilKividor/internal/repository/postgresql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body models.AuthReqModel

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
	}

	//check in db if the user exists with the username
	username, password, err := postgresql.GetUsernamePassword(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// verify the password
	isVerified := comparePassword(password, body.Password)
	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password did not match",
		})
		return
	}
	// generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//save tokens in db
	err = postgresql.UpdateUserTokens(accessToken, refreshToken, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "updating table",
		})
		return
	}
	//send the tokens to frontend
	var res models.AuthResModel
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func Signup(c *gin.Context) {
	var body models.AuthReqModel

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	//hash the password
	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//create user
	var user models.UserDetails
	user.Username = body.Username
	user.Password = hashedPassword
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	//save user in db
	err = postgresql.RegisterNewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResModel
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func RefreshToken(c *gin.Context) {
	var body models.Refreshreq
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	refreshToken := body.RefreshToken

	accessToken, username, err := auth.RefreshAccessToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = postgresql.UpdateUserTokens(accessToken, refreshToken, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResModel
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return string(hash), fmt.Errorf("error hashing the password : %v", err)
	}

	return string(hash), nil
}

func comparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

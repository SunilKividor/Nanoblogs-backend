package handlers

import (
	"net/http"

	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/repository/postgresql"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err = postgresql.DeletUserQuery(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not delete the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

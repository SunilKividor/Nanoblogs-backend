package handlers

import (
	"net/http"

	"github.com/SunilKividor/internal/repository/neo4j"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	//get all blogs from postgres db
	res, err := neo4j.GetTopCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

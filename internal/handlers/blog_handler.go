package handlers

import (
	"net/http"
	"time"

	"github.com/SunilKividor/internal/models"
	"github.com/SunilKividor/internal/repository/postgresql"
	"github.com/gin-gonic/gin"
)

func PostBlog(c *gin.Context) {
	var req models.PostBlogReqModel
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid req body",
		})
	}

	//create a blog
	var blog models.Blog
	blog.UserId = req.UserId
	blog.Title = req.Title
	blog.Content = req.Content
	blog.Category = req.Category
	blog.Created_At = time.Now()
	blog.Updated_At = time.Now()

	//save the blog in db
	err = postgresql.PostBlogQuery(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateBlog(c *gin.Context) {
	var req models.UpdateBlogReqModel
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid req body",
		})
	}

	//update the blog in db
	err = postgresql.UpdateBlogQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

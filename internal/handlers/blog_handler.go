package handlers

import (
	"log"
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
		return
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
	log.Println(req)
	//update the blog in db
	err = postgresql.UpdateBlogQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetAllUserBlogs(c *gin.Context) {
	var req models.GetBlogsReqModel
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid req body",
		})
		return
	}
	log.Println(req)
	// get all blogs by user id from db
	blogs, err := postgresql.GetAllUserBlogsQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

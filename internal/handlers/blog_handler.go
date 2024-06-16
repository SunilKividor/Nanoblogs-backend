package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/models"
	"github.com/SunilKividor/internal/repository/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//create a blog
	var blog models.Blog
	blog.UserId = id
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
}

func UpdateBlog(c *gin.Context) {
	var req models.UpdateBlogReqModel
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid req body",
		})
		return
	}
	log.Println(req)
	if err = uuid.Validate(req.BlogId.String()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id format",
		})
		return
	}

	var updatedBlog models.UpdateBlogDBModel

	updatedBlog.BlogId = req.BlogId
	updatedBlog.Title = req.Title
	updatedBlog.Content = req.Content
	updatedBlog.Category = req.Category
	//update the blog in db
	err = postgresql.UpdateBlogQuery(updatedBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetAllUserBlogs(c *gin.Context) {

	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// get all blogs by user id from db
	blogs, err := postgresql.GetAllUserBlogsQuery(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func GetAllBlogs(c *gin.Context) {
	blogs, err := postgresql.GetAllBlogsQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func DeleteBlog(c *gin.Context) {
	var req models.DeleteBlogReqModel
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid req body",
		})
		return
	}
	log.Println(req)
	if err = uuid.Validate(req.BlogId.String()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id format",
		})
		return
	}
	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//delete blog in db
	err = postgresql.DeleteBlogQuery(req.BlogId, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

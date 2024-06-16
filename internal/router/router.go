package router

import (
	"net/http"

	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.Use(func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, access-control-allow-origin")

		// If the request method is OPTIONS, handle it and return 200 status code
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continue to the next middleware or route handler
		c.Next()
	})
	//auth
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/refresh", handlers.RefreshToken)
	r.GET("/auth/category", handlers.GetAllCategories)
	user := r.Group("/user")
	authorized := r.Group("/blog")
	//middlewares
	user.Use(auth.AuthMiddleware())
	authorized.Use(auth.AuthMiddleware())
	//user-blogs
	user.POST("/blog/post", handlers.PostBlog)
	user.POST("/blog/update", handlers.UpdateBlog)
	user.GET("/blog/get", handlers.GetAllUserBlogs)
	user.DELETE("/blog/delete", handlers.DeleteBlog)
	//user
	user.GET("/user/profile", handlers.GetUser)
	user.DELETE("/user/delete", handlers.DeleteUser)
	//blogs
	authorized.GET("/all", handlers.GetAllBlogs)

	//health
	r.POST("/health", handlers.HelathCheck)
}

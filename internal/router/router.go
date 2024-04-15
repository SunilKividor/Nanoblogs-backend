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
	authorized := r.Group("/user")
	//blogs
	authorized.Use(auth.AuthMiddleware())
	authorized.POST("/blog/post", handlers.PostBlog)
	authorized.POST("/blog/update", handlers.UpdateBlog)
	authorized.GET("/blog/get", handlers.GetAllUserBlogs)
	authorized.DELETE("/blog/delete", handlers.DeleteBlog)
	//user
	authorized.GET("/user/profile", handlers.GetUser)
	authorized.DELETE("/user/delete", handlers.DeleteUser)

	//health
	r.POST("/health", handlers.HelathCheck)
}

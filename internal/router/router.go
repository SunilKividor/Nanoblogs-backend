package router

import (
	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//auth
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/refresh", handlers.RefreshToken)
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

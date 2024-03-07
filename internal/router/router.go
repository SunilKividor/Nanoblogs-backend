package router

import (
	"github.com/SunilKividor/internal/auth"
	"github.com/SunilKividor/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/user/login", handlers.Login)
	r.POST("/user/signup", handlers.Signup)
	r.POST("/user/refresh", handlers.RefreshToken)
	authorized := r.Group("/user")
	authorized.Use(auth.AuthMiddleware())
	authorized.POST("/blog/post", handlers.PostBlog)
	authorized.POST("/blog/update", handlers.UpdateBlog)
	authorized.GET("/blog/get", handlers.GetAllUserBlogs)
	authorized.DELETE("/blog/delete", handlers.DeleteBlog)
}

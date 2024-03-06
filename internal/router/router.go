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
	authorized.GET("/user", handlers.ProtectedRoute)
	authorized.POST("/postblog", handlers.PostBlog)
	authorized.POST("/updateblog", handlers.UpdateBlog)
	authorized.POST("/blogs", handlers.GetAllUserBlogs)
}

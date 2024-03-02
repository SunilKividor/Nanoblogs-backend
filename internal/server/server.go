package server

import (
	"github.com/SunilKividor/internal/router"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	r := gin.Default()
	router.Router(r)
	return r.Run(":8080")
}

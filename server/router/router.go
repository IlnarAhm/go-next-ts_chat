package router

import (
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHanlder *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHanlder.CreateUser)
	r.POST("/login", userHanlder.Login)
	r.GET("/logout", userHanlder.Logout)
}

func Start(addr string) error {
	return r.Run(addr)
}

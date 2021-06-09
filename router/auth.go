package router

import (
	"willow/api"

	"github.com/gin-gonic/gin"
)

func authRouter(engine *gin.RouterGroup) {
	engine.POST("login", api.Login)
}

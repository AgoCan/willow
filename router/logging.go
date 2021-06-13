package router

import (
	"willow/api"

	"github.com/gin-gonic/gin"
)

// 把之前的代码放在这里
func loggingRouter(engine *gin.RouterGroup) {
	engine.GET("logging/:value", api.Logging)
}

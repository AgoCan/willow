package router

import (
	"willow/api"

	"github.com/gin-gonic/gin"
)

func machineRouter(engine *gin.RouterGroup) {
	engine.POST("machine", api.CreateMachine)
	engine.POST("machine/:id", api.UpdateMachine)
}

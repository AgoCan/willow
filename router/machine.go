package router

import (
	"willow/api"

	"github.com/gin-gonic/gin"
)

func machineRouter(engine *gin.RouterGroup) {
	engine.GET("machine", api.QueryMachine)
	engine.GET("machine/:id", api.GetMachine)
	engine.POST("machine", api.CreateMachine)
	engine.PUT("machine/:id", api.UpdateMachine)
	engine.DELETE("machine/:id", api.DeleteMachine)
}

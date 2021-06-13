package router

import (
	"github.com/gin-gonic/gin"

	"willow/api"
	"willow/middleware/auth"
	"willow/middleware/log"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine {
	router := gin.New()
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(log.GinLogger(log.Logger),
		log.GinRecovery(log.Logger, true))
	router.GET("/health", api.Health)

	v1NoAuth := router.Group("/api/v1")
	authRouter(v1NoAuth)
	loggingRouter(v1NoAuth)

	v1Auth := router.Group("/api/v1")
	v1Auth.Use(auth.JWTAuth())
	v1Auth.GET("/healthauth", api.Health)
	return router
}

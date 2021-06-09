package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"willow/service"
)

func Health(c *gin.Context) {
	service := service.Health{}
	res := service.Status()
	c.JSON(http.StatusOK, res)
}

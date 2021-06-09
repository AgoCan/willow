package api

import (
	"net/http"
	"willow/response"
	"willow/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	service := service.User{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		res := service.Login()
		c.JSON(http.StatusOK, res)
	}
}

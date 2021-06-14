package api

import (
	"net/http"
	"willow/response"
	"willow/service"

	"github.com/gin-gonic/gin"
)

func Logging(c *gin.Context) {
	value := c.Param("value")
	if value != "" {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	}

	service := service.Log{
		Value: value,
	}
	res := service.Search()
	c.JSON(http.StatusOK, res)
}

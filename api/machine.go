package api

import (
	"net/http"
	"strconv"
	"willow/response"
	"willow/service"

	"github.com/gin-gonic/gin"
)

func CreateMachine(c *gin.Context) {
	service := service.Machine{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(http.StatusOK, res)
	}
}

func UpdateMachine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.Machine{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		service.ID = id
		res := service.Update()
		c.JSON(http.StatusOK, res)
	}
}

func DeleteMachine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.Machine{}
	service.ID = id
	res := service.Delete()
	c.JSON(http.StatusOK, res)

}

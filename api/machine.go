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

func QueryMachine(c *gin.Context) {
	service := service.Machine{}

	res := service.Query()
	c.JSON(http.StatusOK, res)
}

func GetMachine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.Machine{}

	res := service.Get(id)
	c.JSON(http.StatusOK, res)
}

func CreateMachineGroup(c *gin.Context) {
	service := service.MachineGroup{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		res := service.Create()
		c.JSON(http.StatusOK, res)
	}
}

func UpdateMachineGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.MachineGroup{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		service.ID = id
		res := service.Update()
		c.JSON(http.StatusOK, res)
	}
}

func DeleteMachineGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.MachineGroup{}
	service.ID = id
	res := service.Delete()
	c.JSON(http.StatusOK, res)

}

func QueryMachineGroup(c *gin.Context) {
	service := service.MachineGroup{}

	res := service.Query()
	c.JSON(http.StatusOK, res)
}

func GetMachineGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
		return
	}
	service := service.MachineGroup{}

	res := service.Get(id)
	c.JSON(http.StatusOK, res)
}

func MachineExcute(c *gin.Context) {
	service := service.MachineExcute{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, response.Error(response.ErrCodeParameter))
	} else {
		res := service.Excute()
		c.JSON(http.StatusOK, res)
	}
}

package service

import (
	"database/sql"
	"errors"
	"willow/global"
	"willow/model"
	"willow/response"

	"gorm.io/gorm"
)

type Machine struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	Host       string `json:"host" binding:"required"`
	Port       int    `json:"port"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Type       string `json:"type"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

func (m *Machine) Create() response.Response {
	var machine model.Machine

	if m.Port == 0 {
		machine.Port = 22
	}
	if m.Type == "password" {
		if m.Password == "" {
			return response.Error(response.MachinePasswordIsNull)
		}
		machine.Password = sql.NullString{Valid: true, String: m.Password}
	} else if m.Type == "private" {
		if m.PrivateKey == "" {
			return response.Error(response.MachinePrivateKeyIsNull)
		}
		machine.PrivateKey = sql.NullString{Valid: true, String: m.PrivateKey}
	} else {
		return response.Error(response.MachinePrivateKeyIsNull)
	}
	if m.User == "" {
		machine.User = "root"
	} else {
		machine.User = m.User
	}

	machine.Name = m.Name
	machine.Host = m.Host
	machine.Port = m.Port
	machine.Type = m.Type

	if !errors.Is(global.GDB.Where("name = ?", machine.Name).First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineNameExist)
	}

	if !errors.Is(global.GDB.Where("Host = ?", machine.Host).First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineHostExist)
	}

	if err := global.GDB.Create(&machine).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Success("成功创建机器")
}

func (m *Machine) Update() response.Response {

	var machine model.Machine
	if errors.Is(global.GDB.Where("id = ?", m.ID).First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineNameExist)
	}
	machine.Name = m.Name
	machine.Host = m.Host
	machine.Port = m.Port
	machine.Type = m.Type

	return response.Success("成功更新机器")
}

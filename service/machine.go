package service

import (
	"errors"
	"willow/global"
	"willow/model"
	"willow/response"

	"github.com/jinzhu/copier"
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

	machine, err := model.NewMachine(
		model.SetPort(m.Port),
		model.SetUser(m.User),
		model.SetName(m.Name),
		model.SetHost(m.Host),
		model.SetType(m.Type),
		model.SetPassword(m.Password),
		model.SetPrivateKey(m.PrivateKey),
	)
	if err == model.MachinePasswordIsNull {
		return response.Error(response.MachinePasswordIsNull)
	} else if err == model.MachinePrivateKeyIsNull {
		return response.Error(response.MachinePrivateKeyIsNull)
	}

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
		return response.Error(response.MachineNameNotExist)
	}
	machine, err := model.NewMachine(
		model.SetPort(m.Port),
		model.SetUser(m.User),
		model.SetName(m.Name),
		model.SetHost(m.Host),
		model.SetType(m.Type),
		model.SetPassword(m.Password),
		model.SetPrivateKey(m.PrivateKey),
	)
	if err == model.MachinePasswordIsNull {
		return response.Error(response.MachinePasswordIsNull)
	} else if err == model.MachinePrivateKeyIsNull {
		return response.Error(response.MachinePrivateKeyIsNull)
	}

	if err := global.GDB.Model(&model.Machine{}).Where("id = ?", m.ID).Updates(machine).Error; err != nil {
		return response.Error(response.ErrSQL)
	}

	return response.Success("成功更新机器")
}

func (m *Machine) Delete() response.Response {
	var machine model.Machine
	if errors.Is(global.GDB.Where("id = ?", m.ID).First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineNameNotExist)
	}
	if err := global.GDB.Where("id = ?", m.ID).Delete(&machine).Error; err != nil {
		return response.Error(response.ErrSQL)
	}

	return response.Success("删除OK")
}

func (m *Machine) Query() response.Response {
	var machines []model.Machine
	global.GDB.Find(&machines)

	ms := make([]Machine, len(machines))

	for i, item := range machines {
		copier.Copy(m, item)
		ms[i] = *m
	}
	return response.Success(ms)
}

func (m *Machine) Get(id int) response.Response {
	var machine model.Machine
	if errors.Is(global.GDB.Where("id = ?", id).First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineNameNotExist)
	}
	copier.Copy(m, machine)
	return response.Success(m)
}

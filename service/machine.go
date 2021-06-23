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
	ID           int          `json:"id"`
	Name         string       `json:"name" binding:"required"`
	Host         string       `json:"host" binding:"required"`
	Port         int          `json:"port"`
	PrivateKey   string       `json:"privateKey"`
	PublicKey    string       `json:"publicKey"`
	Type         string       `json:"type"`
	User         string       `json:"user"`
	Password     string       `json:"password"`
	MachineGroup MachineGroup `json:"machine_group"`
}

type MachineGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MachineExcute struct {
	MachineIDS      []int  `json:"machine_ids"`
	MachineGroupIDS []int  `json:"machine_group_ids"`
	Command         string `json:"command" binding:"required"`
}

func (m *Machine) Create() response.Response {
	var machine model.Machine
	var group model.MachineGroup

	if errors.Is(global.GDB.Where("name = ?", m.MachineGroup.Name).First(&group).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineGroupNotExist)
	}

	machine, err := model.NewMachine(
		model.SetPort(m.Port),
		model.SetUser(m.User),
		model.SetName(m.Name),
		model.SetHost(m.Host),
		model.SetType(m.Type),
		model.SetPassword(m.Password),
		model.SetPrivateKey(m.PrivateKey),
		model.SetGroup(group.ID),
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
	if errors.Is(global.GDB.Where("machines.id = ?", id).Joins("MachineGroup").First(&machine).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineNameNotExist)
	}
	copier.Copy(m, machine)
	return response.Success(m)
}

func (g *MachineGroup) Create() response.Response {
	var group model.MachineGroup
	if !errors.Is(global.GDB.Where("name = ?", g.Name).First(&group).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineGroupExist)
	}
	group.Name = g.Name
	if err := global.GDB.Create(&group).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Success("创建OK")
}

func (g *MachineGroup) Delete() response.Response {
	var group model.MachineGroup
	if errors.Is(global.GDB.Where("id = ?", g.ID).First(&group).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineGroupNotExist)
	}
	if err := global.GDB.Where("id = ?", g.ID).Delete(&group).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Success("删除OK")
}

func (g *MachineGroup) Query() response.Response {
	var group []model.MachineGroup
	global.GDB.Find(&group)

	gs := make([]MachineGroup, len(group))

	for i, item := range group {
		copier.Copy(g, item)
		gs[i] = *g
	}
	return response.Success(gs)
}

func (g *MachineGroup) Update() response.Response {
	var group model.MachineGroup
	if errors.Is(global.GDB.Where("id = ?", g.ID).First(&group).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.MachineGroupNotExist)
	}
	if err := global.GDB.Model(&model.MachineGroup{}).Where("id = ?", g.ID).Updates(group).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	group.Name = g.Name

	return response.Success("更新OK")
}
func (g *MachineGroup) Get(id int) response.Response {
	// get 请求应该吧machine也给查询出来进行输送
	var machines []model.Machine
	m := new(Machine)
	global.GDB.Where("machine_group_id = ?", id).Find(&machines)

	ms := make([]Machine, len(machines))

	for i, item := range machines {
		copier.Copy(m, item)
		ms[i] = *m
	}
	return response.Success(ms)
}

func (e *MachineExcute) Excute() response.Response {
	if len(e.MachineGroupIDS) == 0 && len(e.MachineIDS) == 0 {
		return response.Error(response.MachineGroupIDIsNull)
	}
	var machinesByGroup []model.Machine
	if len(e.MachineGroupIDS) != 0 {
		global.GDB.Where("machine_group_id IN ?", e.MachineGroupIDS).Find(&machinesByGroup)
	}
	var machines []model.Machine
	if len(e.MachineIDS) != 0 {
		global.GDB.Find(&machines, e.MachineIDS)
	}

	machines = append(machines, machinesByGroup...)

	// 去重,此处不用int用空结构体更加省空间
	machinesMap := make(map[model.Machine]int, len(machines))
	for _, item := range machines {
		if _, ok := machinesMap[item]; !ok {
			machinesMap[item] = 1
		} else {
			machinesMap[item]++
		}
	}

	m := []model.Machine{}
	for i := range machinesMap {
		m = append(m, i)
	}
	return response.Success("执行ok")
}

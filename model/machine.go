package model

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

var (
	MachinePasswordIsNull   = errors.New("Password is null.")
	MachinePrivateKeyIsNull = errors.New("Private key is null.")
)

type Machine struct {
	gorm.Model
	Name           string `gorm:"type:varchar(128);not null"`
	Host           string `gorm:"type:varchar(128);not null"`
	Port           int
	PrivateKey     sql.NullString `gorm:"type:text"`
	PublicKey      sql.NullString `gorm:"type:text"`
	Type           string         `gorm:"type:varchar(128);not null"`
	User           string         `gorm:"type:varchar(64);not null"`
	Password       sql.NullString `gorm:"type:varchar(128)"`
	MachineGroup   MachineGroup
	MachineGroupID sql.NullInt64
}

type MachineGroup struct {
	gorm.Model
	Name string `gorm:"type:varchar(128);not null;unique"`
}

type MachineOption func(*Machine)

func SetPort(port int) MachineOption {
	if port == 0 {
		port = 22
	}
	return func(m *Machine) {
		m.Port = port
	}
}

func SetType(t string) MachineOption {
	return func(m *Machine) {
		m.Type = t
	}
}

func SetPassword(p string) MachineOption {
	return func(m *Machine) {
		m.Password = sql.NullString{Valid: true, String: p}
	}
}

func SetPrivateKey(p string) MachineOption {
	return func(m *Machine) {
		m.PrivateKey = sql.NullString{Valid: true, String: p}
	}
}

func SetUser(u string) MachineOption {
	if u == "" {
		u = "root"
	}
	return func(m *Machine) {
		m.User = u
	}
}

func SetName(n string) MachineOption {
	return func(m *Machine) {
		m.Name = n
	}
}
func SetHost(h string) MachineOption {
	return func(m *Machine) {
		m.Host = h
	}
}

func SetGroup(id uint) MachineOption {
	if id == 0 {
		return func(m *Machine) {
			m.MachineGroup = MachineGroup{}
		}
	}
	return func(m *Machine) {
		m.MachineGroup = MachineGroup{
			Model: gorm.Model{
				ID: id,
			},
		}
	}
}

func NewMachine(m ...MachineOption) (machine Machine, err error) {
	for _, opt := range m {
		opt(&machine)
	}
	if machine.Type == "password" {
		if machine.Password.String == "" {
			return machine, MachinePasswordIsNull
		}
	} else if machine.Type == "private" {
		if machine.PrivateKey.String == "" {
			return machine, MachinePrivateKeyIsNull
		}
	}

	return
}

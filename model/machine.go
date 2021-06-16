package model

import (
	"database/sql"

	"gorm.io/gorm"
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
	Name string `gorm:"type:varchar(128);not null"`
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

func NewMachine(m ...MachineOption) (machine Machine, err error) {
	for _, opt := range m {
		opt(&machine)
	}
	return
}

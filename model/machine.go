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



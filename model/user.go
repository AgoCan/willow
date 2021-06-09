package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string         `gorm:"type:varchar(128);not null"`
	Password  string         `gorm:"type:varchar(128);not null"`
	Nickname  sql.NullString `gorm:"type:varchar(128)"`
	Role      Role           `gorm:"foreignKey:Role"`
	RoleID    int
	IsSupper  uint8
	IsActive  uint8
	LastLogin time.Time
	LastIP    string `gorm:"type:varchar(64)"`
}

type Role struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(128);not null"`
	Desc        sql.NullString `gorm:"type:varchar(128)"`
	PagePerms   string         `gorm:"type:text"`
	DeployPerms string         `gorm:"type:text"`
	HostPerms   string         `gorm:"type:text"`
}

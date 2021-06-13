package global

import (
	"willow/config"
	"willow/model"
	"willow/pkg/es"

	"gorm.io/gorm"
)

var (
	GDB *gorm.DB
	ES  *es.Elastic
)

func Init() {
	GDB = model.New()
	model.AutoMigrate(GDB)
	if config.Conf.ES.Enable {
		ES = es.New(config.Conf.ES.Address)
	}

}

package global

import (
	"willow/config"
	"willow/model"
	"willow/pkg/es"
	"willow/pkg/redis"

	"gorm.io/gorm"
)

var (
	GDB   *gorm.DB
	ES    *es.Elastic
	Redis *redis.Redis
)

func Init() {
	GDB = model.New()
	model.AutoMigrate(GDB)
	if config.Conf.ES.Enable {
		ES = es.New(config.Conf.ES.Address)
	}
	r := config.Conf.Db
	Redis = redis.New(r.RedisHost(), r.Redis.Password, r.Redis.Port)
}

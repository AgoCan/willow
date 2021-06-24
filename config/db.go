package config

import "fmt"

// DbConfig 数据库配置文件
type DbConfig struct {
	Mysql struct {
		DbName   string
		Password string
		Username string
		Port     string
		Host     string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Db       int
	}
}

func (d *DbConfig) Dsn() string {
	return Conf.Db.Mysql.Username + ":" + Conf.Db.Mysql.Password + "@tcp(" +
		Conf.Db.Mysql.Host + ":" + Conf.Db.Mysql.Port + ")/" + Conf.Db.Mysql.DbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}

func (d *DbConfig) RedisHost() string {
	return fmt.Sprintf("%s:%d", Conf.Db.Redis.Host, Conf.Db.Redis.Port)
}

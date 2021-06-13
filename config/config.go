package config

// 配置文件导入yaml文件是configstruct.go
//
// 配置文件可以使用 -c 的参数
// https://github.com/go-yaml/yaml

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// Conf 单例
var Conf *Config

// Config 配置文件
type Config struct {
	Db  DbConfig
	Log LogConfig
	Jwt Jwt
	ES  ES
}

// LogConfig 日志配置文件
type LogConfig struct {
	LogDirector      string
	LogInfoFilename  string
	LogInfoFilePath  string
	LogErrorFilename string
	LogErrorFilePath string
	LogMaxSize       int
	LogMaxBackups    int
	LogMaxAge        int
	LogLevel         string
}

// 设置配置文件的 环境变量
var (

	// LogDirector 日志目录
	LogDirector string
	// LogInfoFile info日志文件
	LogInfoFilename string
	LogMaxSize      int
	LogMaxBackups   int
	LogMaxAge       int
	LogLevel        string
)

// 获取文件绝对路径
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// InitConfig 初始化配置项
func InitConfig(opt *Option) (err error) {
	viper.SetConfigFile(opt.ConfigFile)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	LogDirector = Conf.Log.LogDirector
	if LogDirector == "" {
		LogDirector = path.Join(path.Dir(getCurrPath()), "log")
	}
	Conf.Log.LogInfoFilePath = path.Join(LogDirector, viper.GetString("log.logInfoFilename"))
	Conf.Log.LogErrorFilePath = path.Join(LogDirector, viper.GetString("log.logErrorFilename"))

	return nil
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"willow/config"
	"willow/model"
	"willow/response"
	"willow/router"
)

var (
	err error
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "c",
				Value:       "config/config.yaml",
				Usage:       "配置文件位置",
				Destination: &config.Opt.ConfigFile,
			},
		},
		Action: func(context *cli.Context) error {
			return nil
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
	response.Init()
	// 初始化配置文件
	config.InitConfig(&config.Opt)
	// 连接数据库并在代码结束后关闭
	model.New()

	// 调用路由组
	router := router.SetupRouter()

	err = router.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
}

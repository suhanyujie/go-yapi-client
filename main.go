package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-yapi-client/config"
	"log"
	"os"
)

var CONFIG_DATA map[string]string

// 入口函数
func main() {
	config.GetConfig()
	app := &cli.App{
		Name:  "yapi client",
		Usage: "commit yapi document",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
		},
		Action: func(context *cli.Context) error {
			argc := context.Args().Len()
			if argc == 1 {
				// 直接执行提交 提交文档 todo
				return nil
			}
			// 当参数个数多于 1 个，需要按值处理 todo
			fileName := context.String("f")
			fmt.Println("%V", fileName)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

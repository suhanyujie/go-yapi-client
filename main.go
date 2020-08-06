package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-yapi-client/config"
	"go-yapi-client/docFile"
	"go-yapi-client/help"
	"log"
	"os"
)

var CONFIG_DATA map[string]string

// 入口函数
func main() {
	config.GetConfig("./env.ini")
	app := &cli.App{
		Name:  "yapi client",
		Usage: "commit yapi document",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
		},
		Action: func(context *cli.Context) error {
			// argc := context.Args().Len()
			argc := len(os.Args)
			fmt.Printf("参数个数为：%d", argc)
			paramLen := argc - 1
			if paramLen == 0 { // todo: show help message
				fmt.Println(help.GetHelpMsg())
				return nil
			} else if paramLen == 1 { // todo: submit a markdown doc. such as: yc a.md

			} else if paramLen == 2 { // todo: submit a markdown doc
				// 直接执行提交 提交文档 todo
				// defaultToken := os.Getenv("YC_TOKEN")
				fileName := context.String("f")
				fullPath, _ := docFile.GetFileFullPath(fileName)
				content, _ := docFile.GetFileContent(fullPath)
				println(string(content))
			} else if paramLen == 3 { // todo: submit a file or some files under a dir with a token from shell ENV

			} else if paramLen == 4 { // todo: submit a file or some files under a dir with a token from token flag

			}
			// 当参数个数多于 1 个，需要按值处理 todo
			fileName := context.String("f")
			fmt.Printf("this is a test: %+v", fileName)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

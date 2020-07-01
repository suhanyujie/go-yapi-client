package config

import (
	"gopkg.in/gcfg.v1"
	"log"
	"reflect"
)

var Config1 Section

type Section struct {
	FirstSection FirstSection `gcfg:"first-section"`
	BaseSection  BaseSection  `gcfg:"base-section"`
	TokenSection TokenSection `gcfg:"token-section"`
}

type FirstSection struct {
	TestKey string `gcfg:"test-key"`
}
type BaseSection struct {
	Host        string `gcfg:"host"`
	ExampleFile string `gcfg:"example-file"`
}
type TokenSection struct {
	HrStaffCenter string `gcfg:"hr-staff-center"`
	DefaultToken  string `gcfg:"default-token"`
}

// 读取配置文件
func GetConfig(filename string) Section {
	err := InitConfigServer(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// curPath, _ := os.Getwd() // 获取当前可执行文件所在路径
	return Config1
}

// 初始化配置读取
func InitConfigServer(filename string) error {
	if filename == "" {
		filename = "env.ini"
	}
	err := gcfg.ReadFileInto(&Config1, filename)
	return err
}

func GetTokenList(config Section) (tokenList map[string]string, err error) {
	s1 := config.TokenSection
	var tokenMap = make(map[string]string)
	obj1 := reflect.TypeOf(s1)
	obj2 := reflect.ValueOf(s1)
	for i := 0; i < obj1.NumField(); i++ {
		tokenMap[obj1.Field(i).Name] = obj2.Field(i).String()
	}
	return tokenMap, err
}

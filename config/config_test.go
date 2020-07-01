package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// 执行单元测试时，执行文件的路径位于 当前 test 文件所在目录
	allConfig := GetConfig("../env.ini")
	if allConfig.FirstSection.TestKey != "123" {
		t.Errorf("current value is %s, expect value is 123", allConfig.FirstSection.TestKey)
	}
	if allConfig.TokenSection.DefaultToken != "1xx12xaadsdaaa3" {
		t.Errorf("current value is %s, expect value is 1xx12xaadsdaaa3", allConfig.FirstSection.TestKey)
	}
}

func TestGetTokenList(t *testing.T) {
	config := GetConfig("../env.ini")
	tokenMap, err := GetTokenList(config)
	if err != nil {
		t.Errorf("%s", err)
	}
	for key, val := range tokenMap {
		fmt.Printf("%s => %s\n", key, val)
	}
	if _, ok := tokenMap["HrStaffCenter"]; !ok {
		t.Errorf("error, no value of key HrStaffCenter")
	}
}

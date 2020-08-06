package docFile

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetFileFullPath(file string) (fullPath string, err error) {
	curPath, _ := os.Getwd()
	fullPath = fmt.Sprintf("%s/%s", curPath, file)
	_, err = ioutil.ReadFile(fullPath)
	CheckErr(err)

	return fullPath, nil
}

func GetFileContent(file string) (content []byte, err error) {
	return ioutil.ReadFile(file)
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

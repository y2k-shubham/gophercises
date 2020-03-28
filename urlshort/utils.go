package urlshort

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func InputFilesPath() string {
	_, path, _, _ := runtime.Caller(0)
	return strings.TrimSuffix(path, "utils.go") + "input_files/"
}

func ReadFile(fileName string) string {
	filePath := InputFilesPath() + fileName
	byteArr, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(byteArr)
}

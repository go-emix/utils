package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetWdPath() (wd string, err error) {
	var dir string
	dir, err = os.Getwd()
	if err != nil {
		return
	}
	wd = strings.ReplaceAll(dir, "\\", "/")
	return
}

func FileIsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func Md5Str(data string) string {
	bytes := []byte(data)
	hash := md5.New()
	hash.Write(bytes)
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

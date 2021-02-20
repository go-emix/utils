package utils

import (
	"testing"
)

func TestMd5Str(t *testing.T) {
	str := Md5Str("123456")
	t.Log(str)
}

func TestGetWdPath(t *testing.T) {
	path, err := GetWdPath()
	PanicError(err)
	t.Log(path)
}

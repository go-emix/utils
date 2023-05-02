package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/fs"
	"os"
	"path/filepath"
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

// Src is file or dir,dist must be dir
// and if it does not exist, create it.
func CopyFile(src, dist string, replace bool) error {
	if dist == "" {
		return nil
	}
	dist, _ = strings.CutSuffix(dist, string(filepath.Separator))
	_ = os.MkdirAll(dist, os.ModePerm)
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		file, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		fn := dist + string(filepath.Separator) + stat.Name()
		err = os.WriteFile(fn, file, os.ModePerm)
		return err
	}
	src, _ = strings.CutSuffix(src, string(filepath.Separator))
	srcName := ""
	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if src == path {
			srcName = info.Name()
			return nil
		}
		_, cut, _ := strings.Cut(path, src)
		if cut == "" {
			return nil
		}
		fn := dist + cut
		if !replace {
			fn = dist + string(filepath.Separator) + srcName + cut
		}
		if info.IsDir() {
			_ = os.MkdirAll(fn, os.ModePerm)
		} else {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			err = os.WriteFile(fn, file, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

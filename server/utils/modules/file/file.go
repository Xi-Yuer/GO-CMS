package file

import "os"

var File = &file{}

type file struct{}

// PathExists 检查文件是否存在
func (f *file) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

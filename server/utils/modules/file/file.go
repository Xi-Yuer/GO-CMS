package file

import "os"

var File = &file{}

type file struct{}

// PathExists 判断文件是否存在
func (f *file) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

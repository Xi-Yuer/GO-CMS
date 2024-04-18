package file

import (
	"fmt"
	"os"
)

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

// CheckOrCreateFolder 检查文件夹是否存在
func (f *file) CheckOrCreateFolder(folderPath string) error {
	// 检查文件夹是否存在
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return fmt.Errorf("创建文件夹失败: %v", err)
		}
		//fmt.Println("文件夹不存在，已创建:", folderPath)
	} else if err != nil {
		// 其他错误，返回错误信息
		return fmt.Errorf("检查文件夹失败: %v", err)
	} else {
		//fmt.Println("文件夹已存在:", folderPath)
	}
	return nil
}

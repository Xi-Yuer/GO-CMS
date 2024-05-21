package buildZip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateFilesAndZip 函数接收文件路径和内容的映射，创建文件并打包成zip文件
func CreateFilesAndZip(files map[string]string) ([]byte, error) {
	// 创建临时目录
	tempDir, err := ioutil.TempDir("", "tempDir")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempDir)

	// 创建所有文件及其目录结构
	for filePath, content := range files {
		fullPath := filepath.Join(tempDir, filePath)
		err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create directories for %s: %v", filePath, err)
		}

		// 写入内容到文件
		err = ioutil.WriteFile(fullPath, []byte(content), os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to write to file %s: %v", filePath, err)
		}
	}

	// 创建zip文件
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// 遍历临时目录并将文件添加到zip中
	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 忽略目录
		if info.IsDir() {
			return nil
		}

		// 创建zip文件中的文件
		relPath := strings.TrimPrefix(path, tempDir+string(os.PathSeparator))
		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("failed to create zip entry for %s: %v", relPath, err)
		}

		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", path, err)
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		// 将文件内容写入zip文件中
		_, err = io.Copy(zipFile, file)
		if err != nil {
			return fmt.Errorf("failed to write file %s to zip: %v", path, err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to add files to zip: %v", err)
	}

	// 关闭zip writer
	err = zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close zip writer: %v", err)
	}

	return buf.Bytes(), nil
}

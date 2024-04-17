package uploadServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var UploadService = &uploadService{}

type uploadService struct{}

// CheckChunk 文件检查
func (u *uploadService) CheckChunk(params *dto.UploadBigFileRequest) (dto.CheckChunkResponse, error) {
	hashPath := "./uploadFile/" + params.Identifier
	var chunkList []string

	var checkChunk dto.CheckChunkResponse
	// 检查路径是否存在
	exists, err := utils.File.PathExists(hashPath)
	if err != nil {
		return checkChunk, err
	}
	if exists {
		// 如果文件夹存在，读取所有文件名并返回
		entries, err := os.ReadDir(hashPath)
		if err != nil {
			return checkChunk, err
		}
		// 文件是否上传完成
		state := 0
		for _, entry := range entries {
			fileName := entry.Name()
			chunkList = append(chunkList, fileName)
			fileBaseName := strings.Split(fileName, ".")[0]

			if fileBaseName == hashPath {
				state = 1
			}
		}
		return dto.CheckChunkResponse{
			Status:    state,
			ChunkList: chunkList,
		}, nil
	}
	return checkChunk, err
}

// UploadChunk 上传区块
func (u *uploadService) UploadChunk(context *gin.Context, params *dto.UploadBigFileRequest) error {

	fileHash := params.Identifier
	file := params.UpFile
	hashPath := "./uploadFile/" + fileHash

	exists, err := utils.File.PathExists(hashPath)
	if err != nil {
		return err
	}
	if !exists {
		if err := os.Mkdir(hashPath, os.ModePerm); err != nil {
			return err
		}
	}

	if err := context.SaveUploadedFile(file, hashPath+"/"+file.Filename); err != nil {
		return err
	} else {
		var chunkList []string
		dir, err := os.ReadDir(hashPath)
		if err != nil {
			return err
		}
		for _, entry := range dir {
			fileName := entry.Name()
			// 本地 mac 文件排除
			if fileName == ".DS_Store" {
				continue
			}
			chunkList = append(chunkList, fileName)
		}
		return nil
	}
}

// MergeChunk 区块合并
func (u *uploadService) MergeChunk(params dto.UploadBigFileRequest) (string, error) {
	fileHash := params.Identifier
	file := params.UpFile
	hashPath := "./uploadFile/" + fileHash

	exists, err := utils.File.PathExists(hashPath)
	if err != nil {
		return "", err
	}
	if exists {
		return "/" + hashPath + "/" + file.Filename, nil
	}
	dir, err := os.ReadDir(hashPath)
	if err != nil {
		return "", err
	}
	// 创建文件
	create, err := os.Create(hashPath + "/" + file.Filename)
	if err != nil {
		return "", err
	}
	defer func(create *os.File) {
		_ = create.Close()
	}(create)

	for _, file := range dir {
		if file.Name() == ".DS_Store" {
			continue
		}
		// 文件缓冲区
		fileBuf, err := os.ReadFile(hashPath + "/" + file.Name())
		if err != nil {
			return "", err
		}
		if _, err := create.Write(fileBuf); err != nil {
			return "", err
		}
	}
	return "/" + hashPath + "/" + file.Filename, nil
}

package uploadServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
	"io"
	"mime/multipart"
	"os"
	"sync"
)

var UploadService = &uploadService{}

var mu sync.Mutex

type uploadService struct{}

func (u *uploadService) CheckFile(params *dto.CheckChunkRequest) (dto.CheckChunkResponse, error) {
	hashPath := config.Config.FILEPATH + params.Identifier
	// 检查路径是否存在
	exists := utils.File.PathExists(hashPath)
	if exists {
		// 如果文件存在，读取所有文件大小并返回
		stat, err := os.Stat(hashPath)
		if err != nil {
			return dto.CheckChunkResponse{}, err
		}
		return dto.CheckChunkResponse{
			HasReadySize: stat.Size(),
		}, err
	}
	return dto.CheckChunkResponse{HasReadySize: 0}, nil
}

func (u *uploadService) UploadChunk(params *dto.UploadBigFileRequest) error {
	mu.Lock()
	defer mu.Unlock()
	fileHash := params.Identifier
	file := params.UpFile
	filePath := config.Config.FILEPATH + fileHash

	var newFile *os.File
	var err error

	if exists := utils.File.PathExists(filePath); !exists {
		newFile, err = os.Create(filePath)
		if err != nil {
			return err
		}
	} else {
		newFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer func(newFile *os.File) {
		_ = newFile.Close()
	}(newFile)

	open, err := file.Open()
	if err != nil {
		return err
	}
	defer func(open multipart.File) {
		_ = open.Close()
	}(open)

	// 逐块读取并写入文件
	_, err = io.Copy(newFile, io.LimitReader(open, 1024*1024))
	if err != nil {
		return err
	}

	return nil
}

func (u *uploadService) FinishUpload(account string, params *dto.UploadFinishRequest) (string, error) {
	if err := repositories.UploadRepository.CreateRecord(account, params); err != nil {
		return "", err
	}
	return config.Config.FILEPATH + params.Identifier, nil
}

func (u *uploadService) DeleteFile(id string) error {
	// 判断文件是否存在
	if _, err := os.Stat(config.Config.FILEPATH + id); err != nil {
		return errors.New("文件不存在")
	}
	// 删除文件
	if err := os.Remove(config.Config.FILEPATH + id); err != nil {
		return err
	}
	// 删除数据库记录
	return repositories.UploadRepository.DeleteRecord(id)
}

func (u *uploadService) GetFileList(params *dto.Page) (*dto.HasTotalResponseData, error) {
	return repositories.UploadRepository.GetRecords(params)
}

func (u *uploadService) DownloadFile(id string) (*dto.UploadRecordResponse, error) {
	info, err := repositories.UploadRepository.GetFileInfo(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}

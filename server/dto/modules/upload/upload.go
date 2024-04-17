package uploadResponsiesModules

import "mime/multipart"

type UploadBigFileRequest struct {
	Identifier string                `form:"identifier" binding:"required"`
	UpFile     *multipart.FileHeader `form:"file" binding:"required"`
}

type CheckChunkResponse struct {
	HasReadySize int64 `json:"hasReadySize"`
}

type CheckChunkRequest struct {
	Identifier string `form:"identifier" binding:"required"`
}

type UploadFinishRequest struct {
	Identifier string `form:"identifier" binding:"required"`
	FileName   string `form:"fileName" binding:"required"`
	FileSize   int64  `form:"fileSize" binding:"required"`
}
type UploadRecordResponse struct {
	FileID     string `json:"fileID"`
	FileName   string `json:"fileName"`
	FileSize   int64  `json:"fileSize"`
	UploadUser string `json:"uploadUser"`
	UploadTime string `json:"uploadTime"`
}

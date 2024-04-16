package uploadResponsiesModules

import "mime/multipart"

type UploadBigFileRequest struct {
	Identifier string                `form:"identifier" binding:"required"`
	Filename   string                `form:"filename" binding:"required"`
	UpFile     *multipart.FileHeader `form:"file" binding:"required"`
}

type CheckChunkResponse struct {
	ChunkList []string `json:"chunkList"`
	Status    int      `json:"status"`
}

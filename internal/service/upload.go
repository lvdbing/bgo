package service

import (
	"mime/multipart"

	"github.com/lvdbing/bgo/internal/model"
)

func (svc *Service) UploadFile(fileType model.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*model.FileInfo, error) {
	return model.Uploader.UploadFile(fileType, file, fileHeader)
}

package model

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/pkg/upload"
	"github.com/lvdbing/bgo/pkg/utils"
)

type FileType int

const (
	FileImage FileType = iota + 1
	FileExcel
	FileWord
	FilePPT
	FileTxt
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

type myUploader struct {
}

var Uploader = &myUploader{}

func (upl *myUploader) UploadFile(fileType FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filename := upl.getFilename(fileHeader.Filename)
	if !upl.checkUploadExt(fileType, filename) {
		return nil, errors.New("file suffix is not supported")
	}
	if !upl.checkUploadMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	uploadPath := upl.getUploadPath()
	if upl.IsNotExist(uploadPath) {
		err := upl.createPath(uploadPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if upl.isPermission(uploadPath) {
		return nil, errors.New("insufficient file permissions")
	}

	if !strings.HasSuffix(uploadPath, "/") {
		uploadPath = uploadPath + "/"
	}
	dst := uploadPath + filename
	err := upload.UploadFile(fileHeader, dst)
	if err != nil {
		return nil, err
	}
	accessUrl := upl.getAccessUrl(filename)
	return &FileInfo{Name: filename, AccessUrl: accessUrl}, nil
}

func (upl *myUploader) getFilename(name string) string {
	ext := upl.getFileExt(name)
	filename := strings.TrimSuffix(name, ext)
	filename = utils.EncodeMD5(filename)
	return filename + ext
}

func (upl *myUploader) getFileExt(name string) string {
	return path.Ext(name)
}

func (upl *myUploader) getUploadPath() string {
	return global.AppSetting.UploadPath
}

func (upl *myUploader) IsNotExist(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func (upl *myUploader) checkUploadExt(t FileType, name string) bool {
	ext := upl.getFileExt(name)
	ext = strings.ToUpper(ext)
	var allowExts []string
	switch t {
	case FileImage:
		allowExts = global.AppSetting.UploadExtsImg
	case FileExcel:
		allowExts = global.AppSetting.UploadExtsExcel
	case FileWord:
		allowExts = global.AppSetting.UploadExtsWord
	case FilePPT:
		allowExts = global.AppSetting.UploadExtsPPT
	case FileTxt:
		allowExts = global.AppSetting.UploadExtsTxt
	}
	for _, allowExt := range allowExts {
		if strings.ToUpper(allowExt) == ext {
			return true
		}
	}
	return false
}

func (upl *myUploader) checkUploadMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	var maxSize int
	switch t {
	case FileImage:
		maxSize = global.AppSetting.UploadMaxSizeImg
	case FileExcel:
		maxSize = global.AppSetting.UploadMaxSizeExcel
	case FileWord:
		maxSize = global.AppSetting.UploadMaxSizeWord
	case FilePPT:
		maxSize = global.AppSetting.UploadMaxSizePPT
	case FileTxt:
		maxSize = global.AppSetting.UploadMaxSizeTxt
	}

	return size < maxSize*1024*1024
}

func (upl *myUploader) isPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func (upl *myUploader) createPath(dst string, perm os.FileMode) error {
	return os.MkdirAll(dst, perm)
}

func (upl *myUploader) getAccessUrl(filename string) string {
	uploadUrl := global.AppSetting.UploadUrl
	if !strings.HasSuffix(uploadUrl, "/") {
		uploadUrl = uploadUrl + "/"
	}
	return uploadUrl + filename
}

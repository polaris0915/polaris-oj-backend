package file_service

import (
	"errors"
	"mime/multipart"
)

type FileValidator struct {
	MaxSize   int64
	AllowType []string
}

func (v *FileValidator) ValidateFile(file *multipart.FileHeader) error {
	// 检查文件大小
	if file.Size > v.MaxSize {
		return errors.New("file size exceeds limit")
	}
	return nil
}

package file_dto

import (
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type FileUploadRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (u *FileUploadRequest) GetValidator() *validator.Validate {
	return validator.New()
}

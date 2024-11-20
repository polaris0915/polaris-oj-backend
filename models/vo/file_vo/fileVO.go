package file_vo

import (
	"github.com/go-playground/validator/v10"
)

type FileVO struct {
	FileUrl string `json:"fileUrl"`
}

func (u *FileVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *FileVO) GetResponseVo(fileUrl string) error {
	u.FileUrl = fileUrl
	return nil
}

package file_service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"polaris-oj-backend/database/mysql"
	"polaris-oj-backend/polaris_logger"
	"polaris-oj-backend/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const maxSize = 2 << 20

// const maxSize = 2 << 8

var allowType = []string{".txt", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".rtf", ".md", ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".tiff", ".webp", ".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a", ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm"}

var FileValidator *fileValidator = initFileValidator()

func initFileValidator() *fileValidator {
	return &fileValidator{
		ctx:       context.Background(),
		maxSize:   maxSize,
		allowType: allowType,
	}
}

type fileValidator struct {
	// TODO: 这里也许有问题
	ctx       context.Context
	maxSize   int64
	allowType []string
}

// 目前只对文件的大小，文件类型的检查
func (v *fileValidator) ValidateFile(file *multipart.FileHeader) (error, bool) {
	// 检查文件大小
	if file.Size > v.maxSize {

		return polaris_logger.Error(v.ctx, "file size exceeds limit"), false
	}
	// 检查文件类型
	ext := filepath.Ext(strings.ToLower(file.Filename))
	for _, path := range v.allowType {
		if path == ext {
			return nil, true
		}
	}

	return polaris_logger.Error(v.ctx, "file type is invalid"), false
}

func (v *fileValidator) GetRandomFileName(oldName string) string {
	ext := strings.ToLower(filepath.Ext(oldName))
	uuid := utils.GetUUID()
	return fmt.Sprintf("%s-%d%s", uuid, time.Now().Unix(), ext)
}

type UploadFileInfo struct {
	File     *multipart.FileHeader
	FileName string
}

func UploadFile(c *gin.Context, request *UploadFileInfo) (string, error) {
	// 校验文件
	if err, ok := FileValidator.ValidateFile(request.File); !ok {
		return "", polaris_logger.Error(c, err.Error())
	}
	// 给文件生成随机名
	fileName := FileValidator.GetRandomFileName(request.File.Filename)
	/*
		已经登录的用户统一用用户的uuid作为文件夹的名字
	*/
	if user, exits := c.Get("user"); exits {
		userInfo, _ := user.(*utils.Claims)
		uploadPath := "./upload/" + userInfo.Identity + "/" + fileName
		// 保存都本地
		if err := c.SaveUploadedFile(request.File, uploadPath); err != nil {
			return "", polaris_logger.Error(c, err.Error())
		}
		return "/api/file/download/" + userInfo.Identity + "/" + fileName, nil
	}
	return "", polaris_logger.Error(c, "未登录")

}

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewService(c *gin.Context) *Service {
	return &Service{
		ctx: c,
		db:  mysql.DB,
	}
}

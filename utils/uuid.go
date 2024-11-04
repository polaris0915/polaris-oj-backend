package utils

import uuid "github.com/satori/go.uuid"

// 生成uuid
func GetUUID() string {
	return uuid.NewV4().String()
}

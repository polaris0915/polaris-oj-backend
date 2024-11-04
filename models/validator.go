package models

import "github.com/go-playground/validator/v10"

// 考虑到不管哪种模型都有可能设计到字段的验证
// 现在将验证器抽象出来
type ModelValidator interface {
	GetValidator() *validator.Validate
}

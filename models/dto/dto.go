package dto

import "github.com/go-playground/validator/v10"

type DtoTransfer interface {
	DtoToModel(modelDst interface{}) error
	ModelToDto(modelSrc interface{}) error
	GetValidator() *validator.Validate
}

// type DtoValidator struct {
// 	V *validator.Validate
// }

// func (dv *DtoValidator) Validator() {
// 	dv.V = validator.New()
// }

// func (dv *DtoValidator) GetValidator() *validator.Validate {
// 	return dv.V
// }

// // GetValidator 返回验证器实例
// func (dv *DtoValidator) GetValidator() *validator.Validate {
// 	return dv.V
// }

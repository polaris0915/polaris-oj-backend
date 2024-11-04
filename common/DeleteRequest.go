package common

import (
	"errors"
	"polaris-oj-backend/constant"

	"reflect"

	"github.com/go-playground/validator"
)

type DeleteRequest struct {
	Identity string `json:"identity" validate:"required,uuid"`
}

func (u *DeleteRequest) GetValidator() *validator.Validate {
	return validator.New()
}

func NewDeleteRequest() *DeleteRequest {
	u := new(DeleteRequest)
	return u
}

func (dr *DeleteRequest) DtoToModel(model any) error {
	// 校验
	if err := dr.GetValidator().Struct(dr); err != nil {
		return err
	}
	// 用泛型处理model中的Identity字段
	modelValue := reflect.ValueOf(model)

	if modelValue.Kind() != reflect.Ptr && modelValue.Elem().Kind() != reflect.Struct {
		return errors.New(constant.PARAMS_ERROR.Message)
	}

	identityField := modelValue.Elem().FieldByName("Identity")
	if !identityField.CanSet() || !identityField.IsValid() {
		return errors.New("field Identity in model cannot be set or valid")
	}
	identityField.Set(reflect.ValueOf(dr.Identity))
	return nil
}

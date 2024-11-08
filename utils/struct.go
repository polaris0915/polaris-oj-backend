package utils

import (
	"errors"
	"reflect"
)

func CopyModels(toValue any, fromValue any) error {
	// 检查传入的是否是指针类型
	to := reflect.ValueOf(toValue)
	from := reflect.ValueOf(fromValue)
	// 判断是否为指针类型并指向结构体
	if to.Kind() != reflect.Ptr || from.Kind() != reflect.Ptr && to.Elem().Kind() != reflect.Struct || from.Elem().Kind() != reflect.Struct {
		// TODO: 需要完善开发人员的日志
		// 做到返回的错误与日志要隔离
		return errors.New("both toValue and fromValue should be pointers to structs")
	}

	to = to.Elem()
	from = from.Elem()
	fromType := from.Type()

	for i := 0; i < from.NumField(); i++ {
		fromField := from.Field(i)
		fromFieldName := fromType.Field(i).Name
		toFiled := to.FieldByName(fromFieldName)
		// 这里设置值的逻辑
		// 只要toValue有效且能设置值
		if toFiled.IsValid() && toFiled.CanSet() {
			// 并且fromValue中对应的有值就赋值，没有就跳过
			// 即只变动传入进来的有有效值的字段
			if !fromField.IsZero() {
				if fromField.Kind() == reflect.Struct || fromField.Kind() == reflect.Slice {
					var str string
					var err error
					if str, err = ModelToJson(fromField); err != nil {
						return errors.New("Failed")
					}
					toFiled.Set(reflect.ValueOf(str))
					continue
				}
				toFiled.Set(fromField)
			}
		}
	}
	return nil
}

func GetBoolPtr(b bool) *bool {
	return &b
}

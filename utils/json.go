package utils

import (
	"encoding/json"
	"errors"
	"polaris-oj-backend/constant"
)

func ModelToJson(obj interface{}) (string, error) {
	tags, err := json.Marshal(obj)
	if err != nil {
		return "", errors.New(constant.PARAMS_ERROR.Message)
	}
	return string(tags), nil
}

func JsonToModel(jsonStr string, obj any) error {
	if err := json.Unmarshal([]byte(jsonStr), obj); err != nil {
		return err
	}
	// fmt.Printf("obj: %v", obj)
	return nil
}

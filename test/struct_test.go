package test

import (
	"fmt"
	"polaris-oj-backend/database/mysql"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"reflect"
	"testing"
)

type DeleteRequest struct {
	Identity string
}

type Question struct {
	Identity string
	Name     string `json:"username"` // 编码后的字段名为 username
	Age      int    `json:"userage"`  // 编码后的字段名为 userage
	Sex      string // 编码后的字段名为 Sex
	Hobby    string `json:"omitempy"` // 字段不进行序列化
}

func (dr *DeleteRequest) xx(model any) {
	modelValue := reflect.ValueOf(model)

	if modelValue.Kind() != reflect.Ptr && modelValue.Elem().Kind() != reflect.Struct {
		fmt.Println("model must a pointer to a structure")
		return
	}

	identityField := modelValue.Elem().FieldByName("Identity")
	if !(identityField.CanSet() && identityField.IsValid()) {
		fmt.Println("model must a pointer to a structure")
		return
	}
	identityField.Set(reflect.ValueOf(dr.Identity))
}

func TestXxx(t *testing.T) {
	DelRequest := &DeleteRequest{
		Identity: "bsdhwiudhiuq",
	}
	question := new(allModels.Question)
	DelRequest.xx(question)
	fmt.Printf("question: %+v", question)

}

func TestGormPreload(t *testing.T) {
	question := &allModels.Question{
		Identity: "c63c9ebe-51f3-40bf-b0ee-fb6f1195425e",
	}

	// user := new(allModels.User)
	res := mysql.DB.Preload("User").First(question, "identity = ?", "c63c9ebe-51f3-40bf-b0ee-fb6f1195425e")
	if res.Error != nil {
		fmt.Println(res.Error.Error())
	}
	t.Logf("question: %+v\n", question)
	// TODO OK: 想办法解决gen生成的模型没有外键关联的代码
	t.Logf("user: %+v\n", question.User)
}

// func TestGormPreload(t *testing.T) {
// 	question := &allModels.Question{
// 		Identity: "c63c9ebe-51f3-40bf-b0ee-fb6f1195425e",
// 	}

// 	// 使用 Preload 加载关联的 User 数据
// 	res := mysql.DB.Preload("User").First(question, "identity = ?", question.Identity)
// 	if res.Error != nil {
// 		t.Fatalf("Error: %s", res.Error.Error())
// 	}
// 	t.Logf("Loaded Question: %+v", question)
// 	t.Logf("Loaded Associated User: %+v", question.User)
// }

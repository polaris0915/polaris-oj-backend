# 迭代内容：

1. 在vscode中使用git



## ⚠️⚠️⚠️更改数据库之后一定要重新生成模型数据

## ⚠️⚠️⚠️更新数据表外键相关的内容之后一定要去/gen/main.go中手动设置外键添加

## ⚠️⚠️⚠️更改api注解或者参数模型一定要重新生成swagger

## ⚠️⚠️⚠️更改接口一定要在vue中重新生成调用函数



#### 使用gorm.io/gen 自动根据 MySQL 数据库中的表生成 GORM 模型代码

1. 安装gormt

   ```shell	
   go get -u gorm.io/gen
   ```

2. 编写代码: 在项目中创建一个代码生成脚本（例如 `gen/main.go`）：

   ```go
   // gen/main.go
   package main
   
   import (
       "gorm.io/driver/mysql"
       "gorm.io/gen"
       "gorm.io/gorm"
   )
   
   func main() {
       // 数据库连接
       dsn := "user:password@tcp(127.0.0.1:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Local"
       db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
       if err != nil {
           panic("failed to connect database")
       }
   
       // 初始化代码生成器
       g := gen.NewGenerator(gen.Config{
           OutPath:      "./model", // 输出路径
           ModelPkgPath: "model",   // 模型包路径
       })
   
       // 使用数据库结构生成模型
       g.UseDB(db)
   
       // 自动生成所有表的模型
       g.ApplyBasic(g.GenerateAllTable()...)
   
       // 生成代码
       g.Execute()
   }
   
   ```

#### 根据后端接口文档自动生成前端请求代码

   ```shell
openapi --input http://localhost:8000/swagger/doc.json --output ./generated --client axios
   ```

 

#### 常用的占位符号

```go
在Go语言中，常用的占位符主要包括以下几：

通用占位：
%v：按值的默认格式输出，适用于任何类型。
%+v：在输出结构体时，会添加字段名。
%#v：输出值的Go语法表示形式。
%T：输出值的类型。
%%：输出一个字面的百分号（%）。

布尔值占位：
%t：将布尔值格式化为字符串“true”或“false”。

整数占位：
%d：十进制表示形式。
%o：八进制表示形式。
%x：不带前缀的十六进制表示形式（小写字母）。
%X：带前缀的十六进制表示形式（大写字母）。
%b：二进制表示形式。
%U：将整数值格式化为Unicode格式（例如“U+1234”）。

浮点数和复数占位：
%f：带小数点的浮点数表示形式。
%e：带小数点的科学计数法表示形式（小写字母e）。
%E：带小数点的科学计数法表示形式（大写字母E）。
%g：根据值的大小选择%e或%f格式化。
%G：根据值的大小选择%E或%f格式化。

字符串和字节切片占位：
%s：字符串表示形式。
%q：将字符串格式化为带引号的字符。
%.(数字)s：截取指定长度的字符串。

指针占位：
%p：将指针值格式化为十六进制表示形式，前缀为0x。
```












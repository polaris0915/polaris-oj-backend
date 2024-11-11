package main

import (
	_ "polaris-oj-backend/config" // 导入配置文件
	"polaris-oj-backend/router"
)

func main() {
	r := router.Router()
	r.Run(":8000")
}

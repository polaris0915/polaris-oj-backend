package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MYSQL_DB_NAME = "polaris-oj"
	// MYSQL_DB_NAME     = "my_db"
	MYSQL_DB_PASSWORD = "ALin0915="
)

var DB = Init()

// 处理mysql数据库连接
func Init() *gorm.DB {
	// 连接本地数据库，数据库的信息应该写入配置文件
	dsn := "root:" + MYSQL_DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + MYSQL_DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:ALin0915=@tcp(127.0.0.1:3306)/gin_gorm-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init error: ", err)
	}
	return db
}

package model

import (
	"fmt"

	"gorm.io/driver/mysql" // 需执行 go get gorm.io/driver/mysql
	"gorm.io/gorm"         // 需执行 go get gorm.io/gorm
)

var DB *gorm.DB
var err error

func init() {
	var dsn = "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")
}

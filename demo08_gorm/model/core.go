package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局变量 DB 和 err
var DB *gorm.DB
var err error

// 请确保数据库 db01 已经存在，如果不存在需要手动在 MySQL 中创建: CREATE DATABASE db1;
var dsn = "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")

	// 自动迁移（创建表）
	// AutoMigrate 会自动创建表、缺失的列和索引，但不会删除未使用的列（为了保护数据）

	err = DB.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println("创建表失败:", err)
	} else {
		fmt.Println("表 authors, students 创建/迁移成功")
	}

}

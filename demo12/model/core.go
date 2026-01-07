package model

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 定义全局变量 DB 和 err
var DB *gorm.DB
var err error

func init() {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()

	// 请确保数据库 db01 已经存在，如果不存在需要手动在 MySQL 中创建: CREATE DATABASE db1;
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		// 打印 SQL 语句
		Logger: logger.Default.LogMode(logger.Info),
		// 开启/关闭事务---true/false
		SkipDefaultTransaction: false,
	})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")

	// 自动迁移（创建表）
	// AutoMigrate 会自动创建表、缺失的列和索引，但不会删除未使用的列（为了保护数据）

	err = DB.AutoMigrate(&Bank{})
	if err != nil {
		fmt.Println("创建表失败:", err)
	} else {
		fmt.Println("表 authors, students 创建/迁移成功")
	}

}

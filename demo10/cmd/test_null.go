package main

import (
	"demo10/model"
	"fmt"
)

func main() {
	fmt.Println("开始测试查询 NULL 值...")

	// 注意：这里依赖 model 包的 init() 函数自动连接数据库
	// 如果 model.DB 为 nil，说明连接失败
	if model.DB == nil {
		fmt.Println("错误：数据库连接未初始化")
		return
	}

	var results []model.Salary
	// 开启 Debug 模式
	err := model.DB.Debug().Where("category IS NULL").Find(&results).Error
	if err != nil {
		fmt.Printf("查询出错: %v\n", err)
		return
	}

	// 发现问题：Category 字段有 NOT NULL 约束，且存的是字符串 "NULL"

	// 1. 修改表结构，允许 NULL
	fmt.Println("修改表结构，允许 category 字段为 NULL...")
	model.DB.Exec("ALTER TABLE sales MODIFY category VARCHAR(30) NULL")

	// 2. 自动修复数据
	fmt.Println("正在将字符串 'NULL' 修复为真正的 SQL NULL...")
	model.DB.Exec("UPDATE sales SET category = NULL WHERE category = 'NULL'")

	// 再次查询验证
	var resultsAfter []model.Salary
	fmt.Println("修复后再次查询 category IS NULL...")
	model.DB.Debug().Where("category IS NULL").Find(&resultsAfter)
	fmt.Printf("修复后查询成功！找到 %d 条数据\n", len(resultsAfter))

	// for _, v := range resultsAfter {
	// 	// categoryStr := "nil"
	// 	if v.Category != nil {
	// 		categoryStr = *v.Category
	// 	}
	// 	// fmt.Printf("ID: %d, Product: %s, Category: %s\n", v.ID, v.Product, categoryStr)
	// }
}

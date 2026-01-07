package model

// Author 定义表结构体
// gorm.Model 包含了 ID, CreatedAt, UpdatedAt, DeletedAt 字段
type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

// TableName 自定义表名，默认是结构体名称的复数形式（students）
// 如果不写这个方法，GORM 默认也会创建名为 students 的表
func (s *Student) TableName() string {
	return "students"
}

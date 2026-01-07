package dept

type Employee struct {
	ID      int    `gorm:"primaryKey"`      // 标准主键 ID
	EmpName string `gorm:"column:emp_name"` // 指定列名
	DeptID  int    `gorm:"column:dept_id"`  // 外键字段

	// Belongs To:
	// foreignKey: 指明本结构体的 DeptID 字段是外键
	// references: 默认引用对方的 ID 字段，所以这里不需要写了！
	Department Department `gorm:"foreignKey:DeptID"`
}

func (e *Employee) TableName() string {
	return "employees"
}

package dept

type Department struct {
	ID       int    `gorm:"primaryKey"`       // 标准主键 ID，Tag 其实也可以省略，GORM 默认就认 ID
	DeptName string `gorm:"column:dept_name"` // 指定列名
	// foreignKey:DeptID 告诉 GORM：
	// "去 Employee 表里找，那个叫 DeptID 的字段存的就是我的 ID"
	Employees []Employee `gorm:"foreignKey:DeptID"` // Has Many: 指明 Employee 表里的 DeptID 是外键
}

func (d *Department) TableName() string {
	return "departments"
}

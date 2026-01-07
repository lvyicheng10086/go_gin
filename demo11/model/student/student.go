package student

//foreignKey外键  如果是表名称加上Id的话默认也可以不配置   如果不是，我们需要通过foreignKey配置外键
//references表示的是主键    默认就是Id   如果是Id的话可以不配置
// 学生表结构体
type Student struct {
	StudentId   int32  `gorm:"primaryKey;column:student_id;type:int"` // 强制使用 int (32位)
	StudentName string `gorm:"column:student_name;size:255"`
	// 修正：完整指定所有 Key，防止 GORM 乱猜
	Courses []Courses `gorm:"many2many:student_courses;joinForeignKey:student_id;References:CourseId;joinReferences:course_id"`
}

func (s *Student) TableName() string {
	return "students"
}

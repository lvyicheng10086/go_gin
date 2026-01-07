package student

//课程表
type Courses struct {
	CourseId   int32  `gorm:"primaryKey;column:course_id;type:int"` // 强制使用 int
	CourseName string `gorm:"column:course_name;size:255"`
	// 反向关联也必须写全
	Student []Student `gorm:"many2many:student_courses;foreignKey:CourseId;joinForeignKey:course_id;References:StudentId;joinReferences:student_id"`
}

func (c *Courses) TableName() string {
	return "courses"
}

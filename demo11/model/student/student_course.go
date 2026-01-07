package student

// 中间表学生课程表
type StudentCourse struct {
	StudentID int32 `gorm:"column:student_id;type:int"`
	CourseID  int32 `gorm:"column:course_id;type:int"`
}

func (s *StudentCourse) TableName() string {
	return "student_courses"
}

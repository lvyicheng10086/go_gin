package apis

import (
	"demo11/model"
	"demo11/model/dept"
	"demo11/model/student"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
}
type CourseController struct {
}
type DepartmentController struct {
}
type EmployeeController struct {
}

func (u *CourseController) GetCourse(c *gin.Context) {
	//查询学生信息的时候显示，该学生选修的所有课程
	var StudentList []student.Student
	// 开启 Debug 模式打印 SQL
	// 先查询课程，再查询学生，关联查询
	model.DB.Debug().Preload("Courses").Where("student_name = ?", "王娜").Find(&StudentList)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  StudentList,
	})
}

// join连表查询
func (u *StudentController) GetStudent(c *gin.Context) {

	//查询学科为Python编程的所有学生姓名
	var students []student.Student
	// 开启 Debug 模式打印 SQL
	// 修正：数据库实际表名是 student_courses (复数)，之前写成了单数
	model.DB.Debug().Joins("JOIN student_courses ON student_courses.student_id = students.student_id").
		Joins("JOIN courses ON courses.course_id = student_courses.course_id").
		Where("courses.course_name = ?", "Python编程").
		Find(&students)

	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  students,
	})
}

// 一对多查询，当前主表是department
func (u *DepartmentController) GetDepartment(c *gin.Context) {

	//查询部门为开发部的所有员工姓名
	var dept []dept.Department
	// 开启 Debug 模式打印 SQL
	model.DB.Debug().Preload("Employees").
		Where("dept_name = ?", "技术研发部").
		Find(&dept)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  dept,
	})
}

// 一对一查询，当前主表是employee
func (u *EmployeeController) GetEmployee(c *gin.Context) {

	//查询员工为王洋的部门信息
	var employee dept.Employee
	// 开启 Debug 模式打印 SQL
	model.DB.Debug().Preload("Department").
		Where("emp_name = ?", "王洋").
		Find(&employee)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  employee,
	})
}

// 查询科目表中没有王娜选修的所有课程
func (u *CourseController) GetCourseNoWangNa(c *gin.Context) {
	//查询学生信息的时候显示，该学生选修的所有课程
	var StudentList []student.Student
	// 开启 Debug 模式打印 SQL
	// 先查询课程，再查询学生，关联查询
	model.DB.Debug().Preload("Courses").Where("student_id != ?", 1).Find(&StudentList)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  StudentList,
	})
}

//查询选修“计算机科学与导论”的学生姓名，结果按照倒叙输出

func (s *CourseController) GetStudentName(c *gin.Context) {
	//当前主表是students
	var StudentList []student.Student
	model.DB.Debug().
		Joins("JOIN student_courses ON student_courses.student_id = students.student_id").
		Joins("JOIN courses ON courses.course_id = student_courses.course_id").
		Where("courses.course_name = ?", "计算机科学导论"). // ✅ 现在可以访问 courses 表字段了
		Order("students.student_id Desc").           // 最好指定表名，防止 ID 冲突
		Preload("Courses").                          //如果你还想在结果里看到课程详情，需要加上这个
		Find(&StudentList)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  StudentList,
	})
}

// 查询指定条件offset ，limit 分页查询
func (s *CourseController) GetStudentNamePage(c *gin.Context) {
	var StudentList []student.Student
	model.DB.Debug().
		Joins("JOIN student_courses ON student_courses.student_id = students.student_id").
		Joins("JOIN courses ON courses.course_id = student_courses.course_id").
		Where("courses.course_name = ?", "计算机科学导论"). // ✅ 现在可以访问 courses 表字段了
		Order("students.student_id").                // 最好指定表名，防止 ID 冲突
		Preload("Courses").                          //如果你还想在结果里看到课程详情，需要加上这个
		Limit(10).
		Offset(10). //跳过第一条数据
		Find(&StudentList)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  StudentList,
	})
}

// Preload中指定条件，在学生表中排除掉李静
func (u *CourseController) GetCourseNoLiJing(c *gin.Context) {
	//查询学生信息的时候显示，该学生选修的所有课程
	var CourseList []student.Student
	// 开启 Debug 模式打印 SQL
	// 先查询课程，再查询学生，关联查询
	model.DB.Debug().Preload("Courses").Where("student_name != ?", "李静").Find(&CourseList).Limit(10)
	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  CourseList,
	})
}

// 自定义预加载函数，
// 场景查询：哪些课程被学生选修，要求按照学生的id倒叙输出
/*
func(db *gorm.DB) *gorm.DB {
		return db.Order("student_id Desc")
	}

*/
func (u *CourseController) GetCourseByStudentIDDesc(c *gin.Context) {
	var CourseList []student.Courses
	// 开启 Debug 模式打印 SQL
	model.DB.Debug().Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Order("student_id Desc")
	}).Find(&CourseList).Limit(10)

	c.JSON(200, gin.H{
		"message": "查询成功",
		"result":  CourseList,
	})
}

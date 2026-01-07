package routers

import (
	"demo11/controllers/apis"
	"demo11/middlewares"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func ApiRouters(r *gin.Engine) {

	apiRouters := r.Group("/api", middlewares.ApiAuth)
	api := apis.StudentController{}
	department := apis.DepartmentController{}
	employee := apis.EmployeeController{}
	course := apis.CourseController{} // 确保已定义 CourseController 类型
	simulation := apis.SimulationController{}

	{
		// 新增模拟接口
		apiRouters.POST("/login", simulation.Login)                        // 登录模块
		apiRouters.GET("/medical-records", simulation.GetMedicalRecords)   // 门诊病历查询
		apiRouters.GET("/policies", simulation.GetPolicies)                // 政策权益查询
		apiRouters.GET("/settlements", simulation.GetSettlements)          // 结算信息查询
		apiRouters.GET("/", api.GetStudent)
		apiRouters.GET("/department", department.GetDepartment)
		apiRouters.GET("/employee", employee.GetEmployee)
		apiRouters.GET("/course", course.GetCourse)
		apiRouters.GET("/course/no/wangna", course.GetCourseNoWangNa)
		apiRouters.GET("/course/student/name", course.GetStudentName)
		apiRouters.GET("/course/student/name/page", course.GetStudentNamePage)
		apiRouters.GET("/course/no/lijing", course.GetCourseNoLiJing)
		apiRouters.GET("/course/student/desc", course.GetCourseByStudentIDDesc)

	}
}

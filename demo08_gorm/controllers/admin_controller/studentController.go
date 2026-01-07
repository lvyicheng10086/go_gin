package admin_controller

import (
	"demo08_gorm/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	BaseController
}

func (u *StudentController) AddUser(c *gin.Context) {
	//新增单个学生到数据库
	// student := model.Student{
	// 	Name:  "王五",
	// 	Age:   19,
	// 	Grade: "高二",
	// }

	// err := model.DB.Create(&student).Error

	// 新增多个学生记录到数据库
	students := []*model.Student{
		{Name: "Jinzhu", Age: 18, Grade: "高一"},
		{Name: "Jackson", Age: 19, Grade: "高二"},
		{Name: "Jannette", Age: 20, Grade: "高三"},
		{Name: "Jake", Age: 21, Grade: "复读"},
	}
	// 新增多个学生记录到数据库
	err := model.DB.Create(students).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "新增用户失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回新增的学生ID
	c.JSON(200, gin.H{
		"message": "新增用户",
		"id":      students,
	})
}
func (u *StudentController) DeleteUser(c *gin.Context) {

	//初始化一个学生对象
	student := []model.Student{}

	//按照主键ID删除
	model.DB.Delete(&student, 6)

	//批量删除
	// model.DB.Where("id>6").Delete(&student)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}
func (u *StudentController) GetUser(c *gin.Context) {
	//查询全部学生,注意是指针类型
	student := []model.Student{}
	// model.DB.Find(&student)

	//按照条件过滤查询
	//查询年龄大于18的学生
	model.DB.Where("age > ?", 18).Find(&student)
	fmt.Println("len(student)", len(student))
	if len(student) == 0 {
		c.JSON(404, gin.H{
			"message": "用户不存在",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "获取用户列表",
		"result":  student,
	})
}

func (u *StudentController) UpdateUser(c *gin.Context) {
	//保存切片类型
	user := []model.Student{}

	//单个更新单列
	//第一步:先查询要更新的用户是否存在(按照主键id查询)
	model.DB.Where("id = ?", 1).Find(&user)

	//第二步:更新用户的name列
	model.DB.Model(&user).Update("name", "lvyicheng")

	//全量更新save
	model.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"user":    user,
	})
}

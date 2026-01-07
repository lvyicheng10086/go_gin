package admin_controller

import (
	"demo10/model"

	"github.com/gin-gonic/gin"
)

/*
	Where条件
	=
	<
	>
	<=
	>=
	!=
	IS NOT NULL
	IS NULL
	BETWEEN AND
	NOT BETWEEN AND
	IN
	OR
	AND
	NOT
	LIKE

使用Select指定返回的字段

# Order排序 、Limit 、Offset

分页

# Count 统计总数

使用原生 sql 删除 user 表中的一条数据
*/
type NavController struct {
	BaseController //继承
}

// Where条件查询
func (a *NavController) AddProduct(g *gin.Context) {
	// 陷阱提示：在 GORM 中，如果复用同一个切片变量进行多次 Find，数据默认是“追加”而不是“覆盖”！
	// 也就是第一次查询的结果还在，第二次查询的结果会跟在后面。
	// 最佳实践：为不同的查询定义不同的变量。

	// 1. 定义变量接收范围查询结果
	salRange := []model.Salary{}
	// 查询 id > 3 且 id < 9 的数据
	model.DB.Where("id > ? AND id < ?", 3, 9).Find(&salRange)

	// 2. 定义变量接收 IN 查询结果
	salIn := []model.Salary{}
	// 使用 IN 查询 id 在 3, 5, 6 中的数据
	model.DB.Where("id IN ?", []int{3, 5, 6}).Find(&salIn)

	// 3.like语句的模糊查询
	salLike := []model.Salary{}
	// 查询 name 中包含 "三" 的数据
	model.DB.Where("product LIKE ?", "%三%").Find(&salLike)

	//between and 查询
	// 查询 2023年1，2月份之间的数据
	salBeteen := []model.Salary{}
	model.DB.Where("sale_date BETWEEN ? and ?", "2023-01-01", "2023-02-28").Find(&salBeteen)
	// OR的查询
	// 查询 name 中包含 "三" 或 "四" 的数据
	salOr := []model.Salary{}
	model.DB.Where("product LIKE ? OR product LIKE ?", "%三%", "%14%").Find(&salOr)

	isNull := []model.Salary{}
	// 查询 category 为 NULL 的数据
	// 开启 Debug 模式，打印 SQL 和 错误详情
	err := model.DB.Debug().Where("category IS NULL").Find(&isNull).Error
	if err != nil {
		// 如果查询出错，在控制台打印
		// 生产环境建议使用日志框架
		println("查询 category IS NULL 出错:", err.Error())
	} else {
		println("查询 category IS NULL 成功，记录数:", len(isNull))
	}

	isIn := []model.Salary{}
	// 查询 category 为整数的记录
	model.DB.Where("id in (?,?)", 3, 5).Find(&isIn)
	println("查询 category 为整数的记录数:", len(isIn))

	//select 关键字返回指定的字段
	// 解决方案：定义一个专门的结构体（DTO）来接收部分字段
	// 这样 JSON 序列化时就只会包含这几个字段
	type ResultDTO struct {
		ID      int    `json:"id"`
		Product string `json:"product"`
	}

	selectResult := []ResultDTO{}

	// 注意：
	// 1. 使用 Model(&model.Salary{}) 告诉 GORM 查哪张表
	// 2. 使用 Scan(&selectResult) 将结果映射到自定义结构体中
	model.DB.Model(&model.Salary{}).Select("id, product").Where("id in (?,?)", 3, 5).Scan(&selectResult)

	println("查询 id 在 3, 5, 6 中的数据数:", len(selectResult))

	g.JSON(200, gin.H{
		"message":        "查询成功",
		"range_result":   salRange, // 这里的名字要清晰
		"in_result":      salIn,
		"like_result":    salLike,
		"between_result": salBeteen,
		"or_result":      salOr,
		"isnull_result":  isNull,
		"isin_result":    isIn,
		"select_result":  selectResult,
	})
}

// 排序
func (a *NavController) OrderByProduct(g *gin.Context) {
	// 按销售额降序排序
	var salaries []model.Salary
	model.DB.Order("amount desc").Find(&salaries)

	g.JSON(200, gin.H{
		"message":  "查询成功",
		"salaries": salaries,
	})
}

// limit 分页查询
/*

	Offset 偏移量，跳过前面的记录数
	Limit 限制返回的记录数

*/
func (a *NavController) Limit(g *gin.Context) {
	// 按销售额降序排序
	var salaries []model.Salary
	model.DB.Order("amount desc").Offset(2).Limit(5).Find(&salaries)

	g.JSON(200, gin.H{
		"message":  "查询成功",
		"salaries": salaries,
	})
}

// Count 统计总数
func (a *NavController) Count(g *gin.Context) {
	var count int64
	model.DB.Model(&model.Salary{}).Count(&count)

	g.JSON(200, gin.H{
		"message": "查询成功",
		"count":   count,
	})
}

//原生sql 查，删，改

// 1，查询
func (a *NavController) QuerySalary(g *gin.Context) {
	// 定义一个结构体变量，用于接收 JSON 数据
	var newSalary model.Salary
	model.DB.Raw("select * from sales where id = ?", 3).Scan(&newSalary)

	g.JSON(200, gin.H{
		"message": "查询成功",
		"salary":  newSalary,
	})
}

// 2.删除
func (a *NavController) DeleteSalary(g *gin.Context) {
	// 修正1：SQL 语法修正
	// 在 SQL 中，判断 NULL 必须用 "IS NULL"，不能用 "= ''" 或 "= NULL"
	sql := "DELETE FROM sales WHERE category IS NULL"

	// 修正2：执行与结果获取
	// DELETE 操作没有返回结果集，所以不能用 Scan
	// 应该接收返回的 *gorm.DB 对象，从中获取 RowsAffected（受影响行数）
	result := model.DB.Exec(sql)

	// 错误处理
	if result.Error != nil {
		g.JSON(500, gin.H{
			"message": "删除失败",
			"error":   result.Error.Error(),
		})
		return
	}

	g.JSON(200, gin.H{
		"message":       "删除成功",
		"deleted_count": result.RowsAffected, // 告诉前端删除了几条数据
	})
}

// 3.修改
func (a *NavController) UpdateSalary(g *gin.Context) {
	// 定义一个结构体变量，用于接收 JSON 数据
	var newSalary model.Salary
	model.DB.Exec("update sales set amount = ? where id = ?", 999999, 3).Scan(&newSalary)

	g.JSON(200, gin.H{
		"message": "修改成功",
		"salary":  newSalary,
	})
}

// 统计销售额
func (a *NavController) SumSalary(g *gin.Context) {
	var sum float64
	model.DB.Raw("select sum(amount) from sales").Scan(&sum)

	g.JSON(200, gin.H{
		"message": "查询成功",
		"sum":     sum,
	})
}

type SalarySelect struct {
	ID      int     `json:"id"`
	Product string  `json:"product"`
	Amount  float64 `json:"amount"`
}

// 使用select关键字返回指定的字段
func (a *NavController) SelectSalary(g *gin.Context) {
	var salaries []SalarySelect
	// 修正：使用 .Model(&model.Salary{}) 明确指定查询 sales 表
	// 否则 GORM 会根据 []SalarySelect 推断表名为 salary_selects
	model.DB.Model(&model.Salary{}).Select("id, product, amount").Find(&salaries)

	g.JSON(200, gin.H{
		"message":  "查询成功",
		"salaries": salaries,
	})
}

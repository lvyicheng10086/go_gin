package yukuaizheng

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type YukuaizhengController struct{}

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 1. 登录模块
// POST /api/v1/yukuaizheng/auth/login
func (c *YukuaizhengController) Login(ctx *gin.Context) {
	//模拟延迟
	time.Sleep(1000 * time.Millisecond)
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Code: 400, Msg: "参数错误: " + err.Error()})
		return
	}

	// 模拟登录验证
	if loginReq.Username == "admin" && loginReq.Password == "123456" {
		ctx.JSON(http.StatusOK, Response{
			Code: 200,
			Msg:  "登录成功",
			Data: gin.H{
				"token":    "eyJhGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_token_123456",
				"userId":   "1001",
				"username": "admin",
				"role":     "admin",
			},
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, Response{Code: 401, Msg: "用户名或密码错误"})
	}
}

// 2. 医疗救助资金支出统计查询
// GET /api/v1/yukuaizheng/medical-assistance/stats
func (c *YukuaizhengController) GetMedicalFundStats(ctx *gin.Context) {
	//模拟延迟
	time.Sleep(800 * time.Millisecond)
	// 模拟返回统计数据
	data := gin.H{
		"total_expenditure": 5000000.00, // 总支出金额
		"beneficiary_count": 1200,       // 受益人数
		"fund_balance":      2000000.00, // 资金结余
		"monthly_stats": []gin.H{
			{"month": "2023-01", "expenditure": 400000},
			{"month": "2023-02", "expenditure": 420000},
			{"month": "2023-03", "expenditure": 380000},
		},
		"update_time": time.Now().Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, Response{Code: 200, Msg: "查询成功", Data: data})
}

// 3. 低收入人群信息录入
// POST /api/v1/yukuaizheng/low-income-population
func (c *YukuaizhengController) AddLowIncomeInfo(ctx *gin.Context) {
	//模拟延迟
	time.Sleep(1200 * time.Millisecond)
	var person struct {
		Name         string  `json:"name" binding:"required"`
		IdCard       string  `json:"id_card" binding:"required"`
		Address      string  `json:"address"`
		AnnualIncome float64 `json:"annual_income"`
		FamilySize   int     `json:"family_size"`
	}

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Code: 400, Msg: "参数错误: " + err.Error()})
		return
	}

	// 模拟录入成功
	ctx.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "录入成功",
		Data: gin.H{
			"id":         "LIP20231220001",
			"created_at": time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}

// 3. 低收入人群信息查询
// GET /api/v1/yukuaizheng/low-income-population?id_card=500103198001011234
func (c *YukuaizhengController) GetLowIncomeInfo(ctx *gin.Context) {
	//模拟延迟
	time.Sleep(500 * time.Millisecond)
	idCard := ctx.Query("id_card")

	// 模拟查询结果
	var list []gin.H
	if idCard != "" {
		list = append(list, gin.H{
			"id":            "LIP20231220001",
			"name":          "张三",
			"id_card":       idCard,
			"address":       "重庆市渝中区解放碑街道",
			"annual_income": 12000.00,
			"family_size":   3,
			"status":        "已建档",
		})
	} else {
		list = []gin.H{
			{
				"id":            "LIP20231220001",
				"name":          "张三",
				"id_card":       "500103198001011234",
				"address":       "重庆市渝中区解放碑街道",
				"annual_income": 12000.00,
				"family_size":   3,
				"status":        "已建档",
			},
			{
				"id":            "LIP20231220002",
				"name":          "李四",
				"id_card":       "500103198505056789",
				"address":       "重庆市江北区观音桥街道",
				"annual_income": 8000.00,
				"family_size":   1,
				"status":        "待审核",
			},
		}
	}

	ctx.JSON(http.StatusOK, Response{Code: 200, Msg: "查询成功", Data: gin.H{"list": list, "total": len(list)}})
}

// 4. 预警信息推送 (模拟获取预警列表)
// GET /api/v1/yukuaizheng/warnings
func (c *YukuaizhengController) GetWarnings(ctx *gin.Context) {
	//模拟延迟
	time.Sleep(569 * time.Millisecond)
	// 模拟预警数据
	warnings := []gin.H{
		{
			"id":           "WARN2023122001",
			"type":         "收入异常",
			"level":        "High",
			"content":      "监测到低保户[张三]近期有大额支出，请核实。",
			"target_user":  "张三",
			"created_time": time.Now().Add(-2 * time.Hour).Format("2006-01-02 15:04:05"),
			"status":       "Pending", // 待处理
		},
		{
			"id":           "WARN2023122002",
			"type":         "医疗报销",
			"level":        "Medium",
			"content":      "监测到[李四]发生大病医疗费用，建议主动介入救助。",
			"target_user":  "李四",
			"created_time": time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05"),
			"status":       "Processed", // 已处理
		},
	}

	ctx.JSON(http.StatusOK, Response{Code: 200, Msg: "获取预警信息成功", Data: warnings})
}

// 4. 救助状态更新
// PUT /api/v1/yukuaizheng/assistance-status/:id
func (c *YukuaizhengController) UpdateAssistanceStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var statusReq struct {
		Status string `json:"status" binding:"required"` // e.g., "Approved", "Rejected", "Processing"
		Remark string `json:"remark"`
	}

	if err := ctx.ShouldBindJSON(&statusReq); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Code: 400, Msg: "参数错误: " + err.Error()})
		return
	}

	// 模拟更新
	ctx.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "状态更新成功",
		Data: gin.H{
			"id":          id,
			"new_status":  statusReq.Status,
			"update_time": time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}

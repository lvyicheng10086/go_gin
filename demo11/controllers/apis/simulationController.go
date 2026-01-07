package apis

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SimulationController struct {
}

// 登录模块
func (s *SimulationController) Login(c *gin.Context) {
	// 模拟延迟
	time.Sleep(500 * time.Millisecond)

	// 模拟请求参数
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 模拟登录成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Login successful",
		"data": gin.H{
			"token":     "simulated-jwt-token-xyz123",
			"userId":    "user_001",
			"userName":  loginReq.Username,
			"loginTime": time.Now().Format(time.RFC3339),
		},
	})
}

// 门诊病历查询
func (s *SimulationController) GetMedicalRecords(c *gin.Context) {
	// 模拟延迟
	time.Sleep(500 * time.Millisecond)

	// 模拟返回门诊病历列表
	records := []gin.H{
		{
			"recordId":   "MR20231219001",
			"date":       "2023-12-19",
			"department": "Internal Medicine",
			"doctor":     "Dr. Zhang",
			"diagnosis":  "Upper Respiratory Infection",
			"details":    "Patient reported cough and fever.",
		},
		{
			"recordId":   "MR20231105002",
			"date":       "2023-11-05",
			"department": "Orthopedics",
			"doctor":     "Dr. Li",
			"diagnosis":  "Ankle Sprain",
			"details":    "Patient twisted ankle while running.",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    records,
	})
}

// 政策权益查询
func (s *SimulationController) GetPolicies(c *gin.Context) {
	// 模拟延迟
	time.Sleep(500 * time.Millisecond)

	// 模拟返回政策权益数据
	policies := []gin.H{
		{
			"policyId":    "POL001",
			"title":       "Medical Insurance Coverage A",
			"description": "Covers 80% of outpatient visits.",
			"validUntil":  "2024-12-31",
		},
		{
			"policyId":    "POL002",
			"title":       "Critical Illness Support",
			"description": "Additional support for critical illnesses.",
			"validUntil":  "2025-06-30",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    policies,
	})
}

// 结算信息查询
func (s *SimulationController) GetSettlements(c *gin.Context) {
	// 模拟延迟
	time.Sleep(1000 * time.Millisecond)

	// 模拟返回结算信息
	settlements := []gin.H{
		{
			"settlementId": "SET2023121901",
			"amount":       150.00,
			"status":       "Paid",
			"date":         "2023-12-19 10:30:00",
			"items": []string{
				"Consultation Fee",
				"Medicine A",
			},
		},
		{
			"settlementId": "SET2023110502",
			"amount":       300.50,
			"status":       "Paid",
			"date":         "2023-11-05 14:20:00",
			"items": []string{
				"X-Ray",
				"Bandage",
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    settlements,
	})
}

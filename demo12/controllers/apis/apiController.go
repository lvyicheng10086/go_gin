package apis

import (
	"demo12/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankController struct {
	BaseController
}

func (con *BankController) GetBank(c *gin.Context) {
	//开始事务
	tx := model.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			// 打印错误信息
			fmt.Println("事务回滚，错误信息:", err)
			con.error(c)
		}
	}()
	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）

	//张三账户里面减去100元
	var u1 model.Bank
	if err := tx.Where("name = ?", "张三").First(&u1).Error; err != nil {
		tx.Rollback()
		con.error(c)
		return
	}
	// 原子更新：直接在数据库层面减 100 (使用 gorm.Expr)
	if err := tx.Model(&u1).Update("bank", gorm.Expr("bank - ?", 100)).Error; err != nil {
		fmt.Println("更新张三账户失败:", err)
		tx.Rollback()
		con.error(c)
		return
	}

	//在李四的账户里面增加100元
	var u2 model.Bank
	if err := tx.Where("name = ?", "李四").First(&u2).Error; err != nil {
		tx.Rollback()
		con.error(c)
		return
	}
	// 原子更新：直接在数据库层面加 100 (使用 gorm.Expr)
	if err := tx.Model(&u2).Update("bank", gorm.Expr("bank + ?", 100)).Error; err != nil {
		fmt.Println("更新李四账户失败:", err)
		tx.Rollback()
		con.error(c)
		return
	}

	// 提交事务
	tx.Commit()

	c.JSON(200, gin.H{
		"message": "转账成功",
		"result": gin.H{
			"张三": u1,
			"李四": u2,
		},
	})
}

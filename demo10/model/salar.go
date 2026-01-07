package model

type Salary struct {
	ID       int     `gorm:"primaryKey"`
	Product  string  `gorm:"column:product"`
	Category *string `gorm:"column:category"`
	Amount   float64 `gorm:"column:amount"`
	SaleDate string  `gorm:"column:sale_date"`
}

func (s *Salary) TableName() string {
	return "sales"
}

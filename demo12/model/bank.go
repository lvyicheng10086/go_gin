package model

type Bank struct {
	ID   int     `gorm:"primarykey"`
	Name string  `gorm:"column:name"`
	Bank float64 `gorm:"column:bank"`
}

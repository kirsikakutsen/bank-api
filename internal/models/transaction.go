package models

import "time"

type Transaction struct {
	ID 			uint 			`gorm:"primaryKey"`
	FromID 		*uint 			`gorm:"default:null"`
	ToID		*uint 			`gorm:"default:null"`
	ToUsername  string 			`gorm:"default:null"`
	Type		TransactionType `gorm:"not null"`
	AmountCents int64 			`gorm:"not null"`
	CreatedAt   time.Time 
}
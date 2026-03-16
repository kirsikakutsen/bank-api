package response

import (
	"go-bank/internal/models"
	"time"
)

type AccountDto struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	BalanceCents int64     `json:"balance"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewAccountDto(acc models.Account) AccountDto {
	return AccountDto{
		ID:           acc.ID,
		Username:     acc.Username,
		BalanceCents: acc.BalanceCents,
		CreatedAt:    acc.CreatedAt,
	}
}

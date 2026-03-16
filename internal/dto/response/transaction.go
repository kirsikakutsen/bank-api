package response

import (
	"go-bank/internal/models"
	"time"
)

type TransactionResponseDto struct {
    ID           uint                   `json:"id"`
    FromID       *uint                   `json:"fromID,omitempty"`
    ToID         *uint                  `json:"toID,omitempty"`
    ToUsername   string                 `json:"toUsername"`
    Type         models.TransactionType `json:"type"`
    AmountCents  int64                  `json:"amountCents"`
    CreatedAt    time.Time              `json:"createdAt"`
}


func NewTransactionResponseDto(tx models.Transaction) TransactionResponseDto {

    return TransactionResponseDto{
        ID:           tx.ID,
        FromID:       tx.FromID,
        ToID:         tx.ToID,
        ToUsername:   tx.ToUsername,
        Type:         tx.Type,
        AmountCents:  tx.AmountCents,
        CreatedAt:    tx.CreatedAt,
    }
}

func NewTransactionsResponseDto() {
    
}

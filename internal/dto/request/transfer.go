package request

type TransferRequest struct {
	ToUsername 	string 	`json:"toUsername" validate:"required"`
	AmountCents	int64	`json:"amountCents" validate:"required"`
}
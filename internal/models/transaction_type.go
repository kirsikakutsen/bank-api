package models

type TransactionType string

const (
    Transfer TransactionType = "transfer"
    Withdraw TransactionType = "withdrawal"
    Deposit  TransactionType = "deposit"
)
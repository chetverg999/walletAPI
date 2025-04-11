package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletRequest struct {
	WalletId      uuid.UUID       `json:"walletId"`
	Amount        decimal.Decimal `json:"amount"`
	OperationType string          `json:"operationType" validate:"required,oneof=DEPOSIT WITHDRAW"`
}

type Wallet struct {
	WalletId uuid.UUID       `db:"walletid" json:"walletId"`
	Amount   decimal.Decimal `db:"amount" json:"amount"`
}

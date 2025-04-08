package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletRequest struct {
	WalletId      uuid.UUID       `json:"walletid"`
	Amount        decimal.Decimal `json:"amount"`
	OperationType string          `json:"operation_type" validate:"required,oneof=DEPOSIT WITHDRAW"`
}

type Wallet struct {
	WalletId uuid.UUID       `db:"walletid" json:"walletid"`
	Amount   decimal.Decimal `db:"amount" json:"amount"`
}

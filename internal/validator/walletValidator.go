package validator

import (
	"github.com/go-playground/validator/v10"
	"taskAPI/internal/model"
	"taskAPI/internal/repository"
)

type WalletValidator struct {
	wallet     model.WalletRequest
	repository *repository.WalletRepository
}

func NewWalletValidator(repository *repository.WalletRepository) *WalletValidator {
	return &WalletValidator{repository: repository}
}

func (w *WalletValidator) Validate(wallet model.WalletRequest) error {
	w.wallet = wallet

	return validator.New().Struct(wallet)
}

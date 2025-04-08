package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"taskAPI/internal/model"
	"taskAPI/internal/repository"
	"taskAPI/internal/validator"
)

type WalletService struct {
	validator  *validator.WalletValidator
	repository *repository.WalletRepository
	ctx        context.Context
}

func NewWalletService(ctx context.Context) *WalletService {
	repository := repository.NewWalletRepository(ctx)

	return &WalletService{repository: repository, ctx: ctx, validator: validator.NewWalletValidator(repository)}
}

func (w *WalletService) Post(c *gin.Context, wallet model.Wallet) {
	w.ValidateRequest(c, wallet)

	//увеличение то вызываем метод репозитори на увеличение

	//вызываем метод репозитори на уменьшение баланса
}

func (w *WalletService) GetAll(c *gin.Context) {
	//заполняем и возвращаем массив со значениями
	w.repository.GetAll(c)
}

func (w *WalletService) GetWalletID(c *gin.Context) {
	walletID, err := uuid.Parse(c.Param("WALLET_UUID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})

		return
	}

	w.repository.GetWalletId(c, walletID)
}

func (w *WalletService) ValidateRequest(c *gin.Context, wallet model.Wallet) {
	if err := w.validator.Validate(wallet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

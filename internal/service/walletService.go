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

func (w *WalletService) DepositPost(c *gin.Context, wallet model.WalletRequest) {
	w.ValidateRequest(c, wallet)
	w.repository.Deposit(c, wallet.WalletId, wallet.Amount)
}

func (w *WalletService) WithdrawPost(c *gin.Context, wallet model.WalletRequest) {
	w.ValidateRequest(c, wallet)
	w.repository.Withdraw(c, wallet.WalletId, wallet.Amount)
}

func (w *WalletService) GetAll(c *gin.Context) {
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

func (w *WalletService) ValidateRequest(c *gin.Context, wallet model.WalletRequest) {
	if err := w.validator.Validate(wallet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

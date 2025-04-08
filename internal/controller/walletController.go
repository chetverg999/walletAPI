package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"taskAPI/internal/model"
	"taskAPI/internal/service"
)

type WalletController struct {
	service *service.WalletService
	ctx     context.Context
}

func NewWalletController(ctx context.Context) *WalletController {
	return &WalletController{ctx: ctx, service: service.NewWalletService(ctx)}
}

func (w *WalletController) Post(c *gin.Context) {
	var wallet model.Wallet

	if err := c.BindJSON(&wallet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	w.service.Post(c, wallet)
}

func (w *WalletController) Get(c *gin.Context) {
	w.service.GetAll(c)
}

func (w *WalletController) GetWalletID(c *gin.Context) {
	w.service.GetWalletID(c)
}

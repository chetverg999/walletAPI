package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
	"taskAPI/internal/model"
)

type WalletRepository struct {
	ctx context.Context
}

func NewWalletRepository(ctx context.Context) *WalletRepository {
	return &WalletRepository{ctx: ctx}
}

func (r *WalletRepository) getDB(c *gin.Context) *sqlx.DB {
	db, exists := c.Get("db")

	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "База данных не доступна"})
	}

	sqlxDB, ok := db.(*sqlx.DB)

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Неверный тип данных для базы данных"})
	}

	return sqlxDB
}

func (r *WalletRepository) GetAll(c *gin.Context) {
	var wallets []model.Wallet
	sqlxDB := r.getDB(c)
	err := sqlxDB.Select(&wallets, "SELECT * FROM wallets")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить значения"})

		return
	}

	c.JSON(http.StatusOK, wallets)
}

func (r *WalletRepository) GetWalletId(c *gin.Context, walletId uuid.UUID) {
	var wallet model.Wallet
	sqlxDB := r.getDB(c)
	err := sqlxDB.Get(&wallet, "SELECT * FROM wallets WHERE walletid = $1", walletId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Неверный ID"})

		return
	}

	c.JSON(http.StatusOK, wallet)
}

//запрос на увеличение баланса

//запрос на уменьшение баланса

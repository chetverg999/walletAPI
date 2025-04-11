package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Неверный ID"})
	}

	c.JSON(http.StatusOK, wallet)
}

func (r *WalletRepository) Deposit(c *gin.Context, walletId uuid.UUID, amount decimal.Decimal) {
	tx, err := r.getDB(c).Beginx()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"failed to connect to DB": err.Error()})
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var currentAmount decimal.Decimal
	err = tx.Get(&currentAmount, `SELECT amount FROM wallets WHERE walletid = $1 FOR UPDATE`, walletId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Неверный ID"})
	}

	_, err = tx.Exec(`UPDATE wallets SET amount = $1 WHERE walletid = $2`, currentAmount.Add(amount), walletId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Неверный ID"})
	}

	err = tx.Commit()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

func (r *WalletRepository) Withdraw(c *gin.Context, walletId uuid.UUID, amount decimal.Decimal) {
	tx, err := r.getDB(c).Beginx()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"failed to connect to DB": err.Error()})
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var currentAmount decimal.Decimal
	err = tx.Get(&currentAmount, `SELECT amount FROM wallets WHERE walletid = $1 FOR UPDATE`, walletId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Неверный ID"})
	}

	if currentAmount.LessThan(amount) {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "Недостаточный баланс"})
	}

	_, err = tx.Exec(`UPDATE wallets SET amount = $1 WHERE walletid = $2`, currentAmount.Sub(amount), walletId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	err = tx.Commit()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

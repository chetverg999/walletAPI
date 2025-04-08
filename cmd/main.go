package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"taskAPI/internal/controller"
	"taskAPI/internal/database"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psql := database.NewPsql().Db()

	defer psql.Close()

	router := gin.Default()
	database.PsqlMiddleware(psql)

	walletController := controller.NewWalletController(context.Background())
	router.Use(database.PsqlMiddleware(psql))
	router.GET("/api/v1/wallets/", walletController.Get)
	router.GET("/api/v1/wallets/:WALLET_UUID", walletController.GetWalletID)
	router.POST("/api/v1/wallet", walletController.Post)

	err = router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}

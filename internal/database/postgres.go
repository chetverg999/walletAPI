package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Psql
)

type Psql struct {
	db *sqlx.DB
}

func (p *Psql) Db() *sqlx.DB {
	return p.db
}

func NewPsql() *Psql {
	once.Do(func() {
		db, err := sqlx.Connect(
			"postgres",
			fmt.Sprintf(
				"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
			),
		)

		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		instance = &Psql{db: db}
	})

	return instance
}

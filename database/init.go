package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Postgres implements all database interface to use
type Postgres struct {
	connection *gorm.DB
}

func (db *Postgres) CreateConnection() {
	connStr := "host=localhost port=5432 user=test_guy dbname=go_retro_dev password=passpass sslmode=disable"

	conn, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.connection = conn
}

func (db *Postgres) CloseConnection() {
	fmt.Println("Closing db")
	db.connection.Close()
}

func NewPostgresDB() *Postgres {
	return &Postgres{}
}

package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" //...
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostgresDB create connection a db
func NewPostgresDB() {
	once.Do(func() {
		var err error
		connStr := "postgres://godb:godb@localhost:5432/godb?sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})

}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

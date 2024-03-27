package postgresdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func New(user, password, host, port, dbName string) (*Postgres, error) {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=enabled", host, port, user, password, dbName)

	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatalf("postgres connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("postgres ping failure: %v", err)
	}

	return &Postgres{db: db}, nil
}

func (this Postgres) Close() {
	err := this.db.Close()
	if err != nil {
		log.Fatalf("postgres close failure: %v", err)
	}
}

func (this Postgres) InsertUser(userName string) error {
	this.db.Exec("INSERT...")

	return nil
}

func (this Postgres) SelectSingleUser(userName string) (string, error) {
	this.db.Exec("SELECT...")

	return "user", nil
}

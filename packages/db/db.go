package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// SQLOptions
type SQLOptions struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

func (s SQLOptions) DataSourceString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.User, s.Password, s.DbName)
}

// OpenConnection to create new connection pool
func OpenConnection(sqlOpt *SQLOptions) (*sql.DB, error) {
	db, err := sql.Open("postgres", sqlOpt.DataSourceString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
	return db, nil
}

package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//fmt.Println("DB_CONN =", config.DBConn)

func InitDB(connectionString string) (*sql.DB, error) {
	// open database
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// ping
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// sec connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}

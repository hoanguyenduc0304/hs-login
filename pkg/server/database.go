package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "hslogin"
)

type Database struct {
	Db *sql.DB
}

func (db *Database) Close() {
	if db.Db != nil {
		fmt.Println("close connect database.....!")
		db.Db.Close()
	}
}

func (db *Database) InitPSQL() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db.Db, err = sql.Open("postgres", psqlInfo)
	fmt.Println("Connecting database")
	if err != nil {
		fmt.Println("Error when connect database")
		return err
	}
	err = db.Db.Ping()
	fmt.Println("ping database")
	if err != nil {
		fmt.Println("Error when connect database")
		return err
	}
	fmt.Println("database connected")
	return nil
}

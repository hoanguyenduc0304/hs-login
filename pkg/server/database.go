package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"hs-login/pkg/util"
	"log"
)

type Database struct {
	Db *sql.DB
}

func Close(d *Database) {
	if d.Db != nil {
		fmt.Println("close connect database.....!")
		d.Db.Close()
	}
}

func InitPSQL() (*Database, error) {
	db := &Database{}
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db.Db, err = sql.Open(config.DBDriver, config.DBSource)
	fmt.Println("Connecting database")
	if err != nil {
		fmt.Println("Error when connect database")
		return nil, err
	}
	err = db.Db.Ping()
	fmt.Println("ping database")
	if err != nil {
		fmt.Println("Error when connect database")
		return nil, err
	}
	fmt.Println("database connected")
	return db, nil
}

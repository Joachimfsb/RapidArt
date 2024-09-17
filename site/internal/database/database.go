package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"rapidart/internal/util"
)

var db *sql.DB

func InitializeDatabase() error {
	log.Printf("initializing database...")

	// https://github.com/go-sql-driver/mysql/wiki/Examples
	// This doesn't actually test the connection...
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 util.Config.Database.User,
		Passwd:               util.Config.Database.Password,
		Net:                  "tcp",
		Addr:                 util.Config.Database.Url,
		DBName:               util.Config.Database.Database,
		AllowNativePasswords: true,
	}

	//get a dn handle
	dbOpen, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("could not connect to the database")
	}

	err = dbOpen.Ping()
	if err != nil {
		panic(err.Error())
	}

	db = dbOpen

	return nil
}

// CloseDatabase closes the connection to databases. This should only be called once.
func CloseDatabase() error {
	log.Println("closing database connection...")

	return nil
}

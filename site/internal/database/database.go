package database

import (
	"database/sql"
	"fmt"
	"log"
	config "rapidart/internal/config"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize and sets up a connection to a remote database. Remember to close the connection.
func InitializeDatabase() error {
	log.Printf("initializing database...")

	// https://github.com/go-sql-driver/mysql/wiki/Examples
	// This doesn't actually test the connection...
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 config.Config.Database.Username,
		Passwd:               config.Config.Database.Password,
		Net:                  "tcp",
		Addr:                 config.Config.Database.Url,
		DBName:               config.Config.Database.Db,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	//get a dn handle
	dbOpen, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("could not connect to the database")
	}

	// Ping db to verify connection opened successfully.
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

	db.Close()

	return nil
}

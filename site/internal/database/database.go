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
		User:                 util.Config.Database.Username,
		Passwd:               util.Config.Database.Password,
		Net:                  "tcp",
		Addr:                 util.Config.Database.Url,
		DBName:               util.Config.Database.Db,
		AllowNativePasswords: true,
		ParseTime:            true,
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

/*
func AddTestCanvases() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	canvasNr := 1
	for canvasNr <= 10 {
		fileName := "daily_line_" + strconv.Itoa(canvasNr) + ".png"
		log.Println("Working on: " + fileName)
		// Construct the relative path
		testcanvasPath := filepath.Join(cwd, "internal", "database", "testCanvases", fileName) //path to temporary picture

		canvasPic, err := ioutil.ReadFile(testcanvasPath)
		if err != nil {
			log.Println(glob.PictureNotFound + ": " + testcanvasPath)
			return err
		}

		// Insert the binary data directly as a BLOB
		sqlCanvases := "INSERT INTO `BasisCanvas` (BasisGalleryId, Type, Image) \nVALUES (?, 'curve', ?);"
		_, err = db.Exec(sqlCanvases, 1, canvasPic)
		if err != nil {
			log.Println("Failed to execute SQL:", err)
			return err
		}

		log.Printf("Successfully loaded test canvas nr: %d\n", canvasNr)
		canvasNr++
	}
	log.Println("Successfully added all test canvases to DB")
	return nil
}
*/

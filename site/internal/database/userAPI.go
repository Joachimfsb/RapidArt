package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"rapidart/internal/crypto"
	"rapidart/internal/models"
	"time"
)

func AddUser(newUser models.RapidUser) error {
	//check if there are any attempts at sql injection
	err := SQLCheck(newUser.Email)
	if err != nil {
		log.Println("sql injection attempt discovered")
		return fmt.Errorf("possible sql injection discovered")
	}

	//checks if mail is already registeres
	rows, err := db.Query("SELECT Email FROM `rapidart`.`user`")
	if err != nil {
		fmt.Println(err)
		log.Println("ERROR: %d", err)
		return fmt.Errorf("ERROR: %v", err)
	}
	defer rows.Close()

	//https://stackoverflow.com/questions/25311162/go-how-to-retrieve-multiple-results-from-mysql-with-sql-db-package
	for rows.Next() {
		var email models.RapidUser

		err = rows.Scan(&email.Email)
		if err != nil {
			return err
		}

		if newUser.Email == email.Email {
			log.Println("Email already exist")
			return fmt.Errorf("Email already exist")
		}
	}

	err = SQLCheck(newUser.Username)
	if err != nil {
		log.Println("sql injection attempt discovered")
		return fmt.Errorf("possible sql injection discovered")
	}

	//checks if username is already registeres
	rows, err = db.Query("SELECT Username FROM `rapidart`.`user`")
	if err != nil {
		fmt.Println(err)
		log.Println("ERROR: %d", err)
		return fmt.Errorf("ERROR: %v", err)
	}
	defer rows.Close()

	//https://stackoverflow.com/questions/25311162/go-how-to-retrieve-multiple-results-from-mysql-with-sql-db-package
	for rows.Next() {
		var username models.RapidUser

		err = rows.Scan(&username.Username)
		if err != nil {
			return err
		}

		if newUser.Username == username.Username {
			log.Println("Username already exist")
			return fmt.Errorf("Username already exist")
		}
	}

	err = SQLCheck(newUser.Displayname)
	if err != nil {
		log.Println("sql injection attempt discovered")
		return fmt.Errorf("possible sql injection discovered")
	}

	newUser.Passwordsalt = crypto.GetMD5Hash(time.Now().String() + newUser.Email)
	newUser.Password = crypto.GetMD5Hash(newUser.Password + newUser.Passwordsalt)

	newUser.CreationTime = time.Now()

	newUser.Role = "user"

	err = SQLCheck(newUser.Bio)
	if err != nil {
		log.Println("sql injection attempt discovered")
		return fmt.Errorf("possible sql injection discovered")
	}

	// Specify the relative file name
	fileName := "tmp.png" // Adjust as necessary

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Construct the relative path
	tempPicPath := filepath.Join(cwd, "internal", "database", fileName)

	/*// Print the current working directory
	fmt.Println("Current working directory:", cwd)

	// Print the relative path
	fmt.Println("Relative path to the file:", relativePath)

	// Check if the file exists
	if _, err := os.Stat(relativePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", relativePath)
	} else {
		fmt.Printf("File %s exists\n", relativePath)
	}*/

	profilePic, err := ioutil.ReadFile(tempPicPath)
	if err != nil {
		log.Println("ERROR: cannot find picture")
		return fmt.Errorf("ERROR: cannot find picture")
	}

	sqlInsert := `
INSERT INTO rapidart.user (
    Username,
    Email,
    DisplayName,
    PasswordHash,
    PasswordSalt,
    CreationTimestamp,
    Role,
    Bio,
    ProfilePicture
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(sqlInsert,
		newUser.Username,
		newUser.Email,
		newUser.Displayname,
		newUser.Password,
		newUser.Passwordsalt,
		newUser.CreationTime.Format("2006-01-02 15:04:05"), // Format the time for MySQL
		newUser.Role,
		newUser.Bio,
		profilePic,
	)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	_, err = db.Exec(sqlInsert)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}

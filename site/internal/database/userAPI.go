package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"rapidart/internal/crypto"
	"rapidart/internal/glob"
	"rapidart/internal/models"
	"time"
)

func AddUser(newUser models.RapidUser) error {
	//check if there are any attempts at sql injection
	err := SQLCheck(newUser.Email)
	if err != nil {
		log.Println(glob.SqlAttempt)
		return fmt.Errorf(glob.SqlAttempt)
	}

	//checks if mail is already registeres
	rows, err := db.Query("SELECT Email FROM `rapidart`.`user`")
	if err != nil {
		fmt.Println(err)
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
			log.Println(glob.EmailAlreadyExist)
			return fmt.Errorf(glob.EmailAlreadyExist)
		}
	}

	err = SQLCheck(newUser.Username)
	if err != nil {
		log.Println(glob.SqlAttempt)
		return fmt.Errorf(glob.SqlAttempt)
	}

	//checks if username is already registeres
	rows, err = db.Query("SELECT Username FROM `rapidart`.`user`")
	if err != nil {
		fmt.Println(err)
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
			log.Println(glob.UsernameAlreadyExist)
			return fmt.Errorf(glob.UsernameAlreadyExist)
		}
	}

	err = SQLCheck(newUser.Displayname)
	if err != nil {
		log.Println(glob.SqlAttempt)
		return fmt.Errorf(glob.SqlAttempt)
	}

	newUser.Passwordsalt = crypto.GetMD5Hash(time.Now().String() + newUser.Email)
	newUser.Password = crypto.GetMD5Hash(newUser.Password + newUser.Passwordsalt)

	newUser.CreationTime = time.Now().String()

	newUser.Role = "user"

	err = SQLCheck(newUser.Bio)
	if err != nil {
		log.Println(glob.SqlAttempt)
		return fmt.Errorf(glob.SqlAttempt)
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
	tempPicPath := filepath.Join(cwd, "internal", "database", fileName) //path to temporary picture

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
		log.Println(glob.PictureNotFound)
		return fmt.Errorf(glob.PictureNotFound)
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
		newUser.CreationTime, // Format the time for MySQL
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

func UserLogin(findUser models.UserAuthentication) (models.RapidUser, error) {
	var user models.RapidUser

	err := SQLCheck(findUser.Email)
	if err != nil {
		fmt.Println(err)
		return models.RapidUser{}, err
	}

	row := db.QueryRow("SELECT Email, PasswordHash, PasswordSalt FROM rapidart.user WHERE Email = ?", findUser.Email)

	err = row.Scan(&user.Email, &user.Password, &user.Passwordsalt)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.InvalidNameOrPass)
		return models.RapidUser{}, fmt.Errorf(glob.InvalidNameOrPass)
	}
	if err != nil {
		fmt.Println(err)
		return models.RapidUser{}, err
	}

	return user, nil
}

func UserById(id int) (models.RapidUser, error) {
	var user models.RapidUser

	row := db.QueryRow("SELECT * FROM rapidart.user WHERE UserId = ?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.Passwordsalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.UserNotFound)
		return models.RapidUser{}, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.RapidUser{}, err
	}

	return user, nil
}

func UserByEmail(email string) (models.RapidUser, error) {
	var user models.RapidUser

	row := db.QueryRow("SELECT * FROM rapidart.user WHERE Email = ?", email)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.Passwordsalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.UserNotFound)
		return models.RapidUser{}, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.RapidUser{}, err
	}

	return user, nil
}

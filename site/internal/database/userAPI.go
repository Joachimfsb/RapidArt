package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"rapidart/internal/crypto"
	"rapidart/internal/glob"
	"rapidart/internal/models"
	"time"
)

func AddUser(newUser models.RapidUser) error {
	//checks if mail is already registeres
	rows, err := db.Query("SELECT Email FROM `rapidart`.`user` WHERE Email = ?", newUser.Email)
	if err != nil {
		log.Println("Error på mail")
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	defer rows.Close()

	//https://stackoverflow.com/questions/25311162/go-how-to-retrieve-multiple-results-from-mysql-with-sql-db-package
	for rows.Next() {
		var email models.RapidUser

		err = rows.Scan(&email.Email) //scan Email from db into email
		if err != nil {
			log.Println(glob.ScanFailed)
			return err
		}

		if newUser.Email == email.Email {
			log.Println(glob.EmailAlreadyExist)
			return fmt.Errorf(glob.EmailAlreadyExist)
		}
	}

	//checks if username is already registeres
	rows, err = db.Query("SELECT Username FROM `user` WHERE Username = ?", newUser.Username)
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
			log.Println(glob.ScanFailed)
			return err
		}

		if newUser.Username == username.Username {
			log.Println(glob.UsernameAlreadyExist)
			return fmt.Errorf(glob.UsernameAlreadyExist)
		}
	}

	newUser.PasswordSalt = generatePasswordSalt()
	newUser.Password = crypto.PBDKF2(newUser.Password, newUser.PasswordSalt)

	newUser.CreationTime = time.Now().String()

	newUser.Role = "user"

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
INSERT INTO user (
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
		newUser.PasswordSalt,
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

func UserLogin(newUser models.UserAuthentication) (models.RapidUser, error) {
	var user models.RapidUser

	row := db.QueryRow("SELECT Email, PasswordHash, PasswordSalt FROM `user` WHERE Email = ?", newUser.Email)

	err := row.Scan(&user.Email, &user.Password, &user.PasswordSalt)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.InvalidNameOrPass)
		return models.RapidUser{}, fmt.Errorf(glob.InvalidNameOrPass)
	}
	if err != nil {
		fmt.Println(err)
		return models.RapidUser{}, err
	}

	newUser.Password = crypto.PBDKF2(newUser.Password, user.PasswordSalt)
	if newUser.Password != user.Password {
		log.Println(glob.InvalidNameOrPass)
		return models.RapidUser{}, fmt.Errorf(glob.InvalidNameOrPass)
	}

	return user, nil
}

func UserById(id int) (models.RapidUser, error) {
	var user models.RapidUser

	row := db.QueryRow("SELECT * FROM user WHERE UserId = ?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

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

	row := db.QueryRow("SELECT * FROM user WHERE Email = ?", email)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

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

// This function is inspired from
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func generatePasswordSalt() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator
	abc := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	//adds 5 char random char from abc into the array
	b := make([]byte, 5)
	for i := range b {
		b[i] = byte(abc[rand.Intn(len(abc))])
	}
	return string(b)
}

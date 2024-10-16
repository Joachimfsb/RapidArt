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

func AddUser(newUser models.User) error {
	//checks if mail is already registeres
	rows, err := db.Query("SELECT Email FROM `rapidart`.`User` WHERE Email = ?", newUser.Email)
	if err != nil {
		log.Println("Error på mail")
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	defer rows.Close()

	//https://stackoverflow.com/questions/25311162/go-how-to-retrieve-multiple-results-from-mysql-with-sql-db-package
	for rows.Next() {
		var email models.User

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
	rows, err = db.Query("SELECT Username FROM `User` WHERE Username = ?", newUser.Username)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	defer rows.Close()

	//https://stackoverflow.com/questions/25311162/go-how-to-retrieve-multiple-results-from-mysql-with-sql-db-package
	for rows.Next() {
		var username models.User

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

	newUser.PasswordSalt = crypto.GenerateRandomCharacters(5)
	newUser.Password = crypto.PBDKF2(newUser.Password, newUser.PasswordSalt)

	newUser.CreationTime = time.Now()

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

	profilePic, err := ioutil.ReadFile(tempPicPath)
	if err != nil {
		log.Println(glob.PictureNotFound)
		return fmt.Errorf(glob.PictureNotFound)
	}

	sqlInsert := `
INSERT INTO User (
    Username,
    Email,
    DisplayName,
    PasswordHash,
    PasswordSalt,
    CreationDateTime,
    Role,
    Bio,
    ProfilePicture
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

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

func GetUserById(id int) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE UserId = ?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.UserNotFound)
		return models.User{}, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE Email = ?", email)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE Username = ?", username)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/crypto"
	"rapidart/internal/glob"
	"rapidart/internal/models"
	"time"
)

// Inserts the specified user into the database.
//
// NOTE: displayName, role, bio and profilePic are optional.
func AddUser(
	email string,
	username string,
	password string,
	displayName string,
	role string,
	bio string,
	profilePic []byte,
) error {

	newUser := models.User{
		Email:       email,
		Username:    username,
		Displayname: displayName,
		Bio:         bio,
		Profilepic:  profilePic,
	}

	newUser.PasswordSalt = crypto.GenerateRandomCharacters(5)
	newUser.Password = crypto.PBDKF2(password, newUser.PasswordSalt)
	newUser.CreationTime = time.Now()
	if role == "moderator" || role == "admin" {
		newUser.Role = role
	} else {
		newUser.Role = "User"
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

	_, err := db.Exec(sqlInsert,
		newUser.Username,
		newUser.Email,
		newUser.Displayname,
		newUser.Password,
		newUser.PasswordSalt,
		newUser.CreationTime,
		newUser.Role,
		newUser.Bio,
		newUser.Profilepic,
	)
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

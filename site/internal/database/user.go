package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/glob"
	"rapidart/internal/models"
)

// Inserts the specified user into the database.
func AddUser(newUser models.User) error {

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

func GetUserProfilePic(id int) ([]byte, error) {
	var user models.User

	row := db.QueryRow("SELECT ProfilePicture FROM User WHERE UserId = ?", id)
	err := row.Scan(&user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(glob.UserNotFound)
		return nil, fmt.Errorf(glob.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user.Profilepic, nil
}

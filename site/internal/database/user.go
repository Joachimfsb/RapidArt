package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		newUser.Username,
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

	if user.Profilepic == nil {
		fileName := "tmp.png" // Adjust as necessary

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return []byte{}, err
		}

		// Construct the relative path
		tempPicPath := filepath.Join(cwd, "internal", "database", fileName)

		user.Profilepic, err = ioutil.ReadFile(tempPicPath)
		if err != nil {
			log.Println("ERROR: cannot find picture")
			return []byte{}, fmt.Errorf("ERROR: cannot find picture")
		}

	}

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

// Fetches users and their follower counts
func GetUsersWithFollowerCountSortedByMostFollowers(limit int) ([]models.UserExtended, error) {
	query := `
    SELECT u.UserId, u.Displayname, u.ProfilePicture, COUNT(f.FolloweeUserId) AS FollowerCount
    FROM User u
    LEFT JOIN rapidart.Follow f ON u.UserId = f.FolloweeUserId
    GROUP BY u.UserId
    ORDER BY FollowerCount DESC
    LIMIT ?;
    `

	// Execute the query
	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store the results
	var users []models.UserExtended

	// Iterate through the rows
	for rows.Next() {
		var user models.UserExtended
		err := rows.Scan(&user.UserId, &user.Displayname, &user.Profilepic, &user.FollowerCount)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

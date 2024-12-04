package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/consts"
	"rapidart/internal/crypto"
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

// Get a user by their id
func GetUserById(id int) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE UserId = ?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(consts.UserNotFound)
		return models.User{}, fmt.Errorf(consts.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

// Get a user by their email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE Email = ?", email)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, fmt.Errorf(consts.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

// Get a user by their username
func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	row := db.QueryRow("SELECT * FROM User WHERE Username = ?", username)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, fmt.Errorf(consts.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	// Convert times to local
	user.CreationTime = user.CreationTime.Local()

	return user, nil
}

// Get a users profile picture
func GetUserProfilePic(id int) ([]byte, error) {
	var user models.User

	row := db.QueryRow("SELECT ProfilePicture FROM User WHERE UserId = ?", id)
	err := row.Scan(&user.Profilepic)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(consts.UserNotFound)
		return nil, fmt.Errorf(consts.UserNotFound)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user.Profilepic, nil
}

// Fetches users and their follower counts
//
// The following fields are populated: UserId, Username, DisplayName, ProfilePicture, FollowerCount
func GetUsersWithFollowerCountSortedByMostFollowers(limit int) ([]models.UserExtended, error) {
	query := `
    SELECT u.UserId, u.Username, u.Displayname, u.ProfilePicture, COUNT(f.FolloweeUserId) AS FollowerCount
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
		err := rows.Scan(&user.UserId, &user.Username, &user.Displayname, &user.Profilepic, &user.FollowerCount)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Fetches the users with the most total likes. Sorted by most likes
//
// Returns: User data + like count, error
func GetUsersWithMostTotalLikes(limit int) ([]models.UserExtended, error) {
	query := `
		SELECT u.*, COUNT(l.PostId) AS LikeCount
		FROM User u
		LEFT OUTER JOIN ` + "`Post`" + ` p ON p.UserId = u.UserId
		LEFT OUTER JOIN ` + "`Like`" + ` l ON l.PostId = p.PostId
		GROUP BY p.UserId
		ORDER BY LikeCount DESC
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
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Displayname, &user.Password, &user.PasswordSalt, &user.CreationTime, &user.Role, &user.Bio, &user.Profilepic, &user.TotalLikes)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Updates the user information
func UpdateUser(updatedUser models.User) error {
	user, err := GetUserById(updatedUser.UserId)
	if err != nil {
		log.Println("failed to find user", err)
		return err
	}
	updatedPassword := crypto.PBDKF2(updatedUser.Password, user.PasswordSalt) //hashes new password with original salt

	sqlUpdate := `
UPDATE User SET
    Username = ?,
    Email = ?,
    DisplayName = ?,
    PasswordHash = ?,
    Bio = ?,
    ProfilePicture = ?
	WHERE UserId = ?;
`

	_, err = db.Exec(sqlUpdate,
		updatedUser.Username,
		updatedUser.Email,
		updatedUser.Displayname,
		updatedPassword,
		updatedUser.Bio,
		updatedUser.Profilepic,
		updatedUser.UserId,
	)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}

// searches for a users name/displayname
func SearchUsers(query string) ([]models.User, error) {
	queryPattern := "%" + query + "%"

	sqlQuery := `
    SELECT UserId, Username, DisplayName
    FROM User
    WHERE Username LIKE ? OR DisplayName LIKE ?
    LIMIT 20;
    `

	rows, err := db.Query(sqlQuery, queryPattern, queryPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserId, &user.Username, &user.Displayname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

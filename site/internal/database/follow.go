package database

import (
	"fmt"
	"log"
	"rapidart/internal/models"
)

func NewFollow(follow models.Follow) error {
	sqlInsert := `
		INSERT INTO Follow (
		                  FollowerUserId,
		                  FolloweeUserId
		) VALUES (?, ?);`

	_, err := db.Exec(sqlInsert,
		follow.FollowerUserId,
		follow.FolloweeUserId,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}

// Removes a follow
//
// Returns: Success/Fail, error
func RemoveFollow(followerId int, followeeId int) (bool, error) {
	res, err := db.Exec("DELETE FROM `Follow` WHERE FolloweeUserId = ? AND FollowerUserId = ?;", followeeId, followerId)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	if err == nil && rows != 1 {
		// Should remove 1 row
		return false, nil
	} // If db does not support it, assume success

	return true, err
}

// Gets a list of userIds of a users followers
func GetFollowersForUser(userId int) ([]int, error) {

	var userIds []int

	rows, err := db.Query(""+
		"SELECT f.FollowerUserId "+
		"FROM `User` u "+
		"JOIN `Follow` f ON f.FolloweeUserId = u.UserId "+
		"WHERE u.UserId = ?;", userId)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId int
		err = rows.Scan(&userId)
		if err != nil {
			return []int{}, err
		}

		userIds = append(userIds, userId)
	}

	return userIds, nil
}

// Gets a list of userIds that a users follows
func GetFollowsForUser(userId int) ([]int, error) {

	var userIds []int

	rows, err := db.Query(""+
		"SELECT f.FolloweeUserId "+
		"FROM `User` u "+
		"JOIN `Follow` f ON f.FollowerUserId = u.UserId "+
		"WHERE u.UserId = ?;", userId)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var userId int
		err = rows.Scan(&userId)
		if err != nil {
			return []int{}, err
		}

		userIds = append(userIds, userId)
	}

	return userIds, nil
}

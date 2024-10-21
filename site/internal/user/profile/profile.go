package profile

import (
	"log"
	"rapidart/internal/database"
)

func GetUserProfilePic(userId int) ([]byte, error) {
	profilePic, err := database.GetUserProfilePic(userId)
	if err != nil {
		log.Println("Error fetching profile picture:", err)
		return nil, err
	}
	return profilePic, nil
}

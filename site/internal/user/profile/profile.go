package profile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"rapidart/internal/consts"
	"rapidart/internal/database"
)

// Fetches a user's profile picture by their ID
func GetUserProfilePic(userId int) ([]byte, error) {
	profilePic, err := database.GetUserProfilePic(userId)
	if err != nil {
		log.Println("Error fetching profile picture:", err)
		return nil, err
	}
	if profilePic == nil {
		profilePic, err = TemporaryProfilePic(userId)
		if err != nil {
			fmt.Println("could not insert temp profile pic")
			return nil, err
		}
	}

	return profilePic, nil
}

// Fetch the temporary profile picture
func TemporaryProfilePic(userId int) ([]byte, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("ERROR: " + err.Error())
		return []byte{}, err
	}

	// Construct the relative path
	tempPicPath := filepath.Join(cwd, "web", consts.DefaultProfilePicturePath)
	Profilepic, err := ioutil.ReadFile(tempPicPath)
	if err != nil {
		log.Println("ERROR: cannot find picture")
		return []byte{}, fmt.Errorf("ERROR: cannot find picture")
	}
	return Profilepic, nil
}

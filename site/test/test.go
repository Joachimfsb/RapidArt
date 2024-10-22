package test

import (
	"rapidart/internal/crypto"
	"rapidart/internal/models"
	"time"
)

func GenTestUser() models.User {

	salt := crypto.GenerateRandomCharacters(5)

	return models.User{
		UserId:       1, // Should not matter
		Username:     "test",
		Email:        "test@test.com",
		Displayname:  "Test testesen",
		Password:     crypto.PBDKF2("test", salt),
		PasswordSalt: salt,
		CreationTime: time.Now(),
		Role:         "user",
		Bio:          "My name is user!",
		Profilepic:   nil,
	}
}

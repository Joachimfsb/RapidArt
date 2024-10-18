package test

import (
	"math/rand"
	"rapidart/internal/crypto"
	"rapidart/internal/models"
	"time"
)

func GenTestUser() models.User {

	salt := crypto.GenerateRandomCharacters(5)

	return models.User{
		UserId:       rand.Intn(5000), // Should not matter
		Username:     crypto.GenerateRandomCharacters(5),
		Email:        crypto.GenerateRandomCharacters(5) + "@" + crypto.GenerateRandomCharacters(5) + ".com",
		Displayname:  crypto.GenerateRandomCharacters(5),
		Password:     crypto.PBDKF2(crypto.GenerateRandomCharacters(5), salt),
		PasswordSalt: salt,
		CreationTime: time.Now(),
		Role:         "user",
		Bio:          crypto.GenerateRandomCharacters(50),
		Profilepic:   nil,
	}
}

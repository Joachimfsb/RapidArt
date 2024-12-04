// Contains cryptographic related functions.
package crypto

import (
	"crypto/sha256"
	"fmt"

	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

// https://en.wikipedia.org/wiki/PBKDF2
// Creates a key from an input password + salt.
// Key may also be used as a hash.
func PBDKF2(pwd string, salt string) string {
	dk := pbkdf2.Key([]byte(pwd), []byte(salt), 4096, 128, sha256.New)
	return fmt.Sprintf("%x", dk)
}

// This function is inspired by
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GenerateRandomCharacters(length int) string {
	abc := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")

	//adds 5 char random char from abc into the array
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(abc[rand.Intn(len(abc))])
	}
	return string(b)
}

// Package myCrypto is our own implementation of commonly used cryptographic functions, tailored to our project's needs.
// Cryptographic functions in this file are never implemented in house, but instead repurposed for our needs.
package crypto

/*
// GetMD5Hash hashes an input string with MD5.
//
// Description:
// Generate a new MD5 hashed string based on the input string provided. If the provided string is unique, then the
// output hash will most likely also be unique. Please DO NOT use this algorithm for password hashing, as it is
// considered weak and prone to collisions.
//
// Parameters:
// - text: The returning hash is generated based on this string. To ensure a random string, you can use time.Now().
//
// Returns:
// random hash based on the current timestamp.
//
// Example:
// import time
// hash := myCrypto.GetMD5Hash(time.Now().String())
//
// Credit:
// https://stackoverflow.com/a/25286918
//
// Notes:
// The function outputs a string of numerical values. The reason is this function is currently only used for
// producing IDs used for registrations and notifications. In our implementations it is easier to retrieve ID
// if it is a numerical value, formatted as a string.
//
// Disclaimers:
// Do not use this function to hash passwords. MD5 is notoriously known for collisions.
func GetMD5Hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// BytesToInt converts an array of bytes to a series of numerical integer values
// Currently this function is only used for producing IDs. This is far from optimal as the length of IDs
// vary, but it helps to retrieve IDs from URLs.
//
// Parameters:
// - byteArray: the array to convert to series of integers
//
// Returns:
// An array of integers.
func BytesToInt(byteArray []byte) int {
	var result int
	for index, b := range byteArray {
		offset := math.Pow(10, float64(len(byteArray)-index-1))
		result += (int(b) % 9) * int(offset)
	}

	return int(math.Abs(float64(result)))
}

// https://en.wikipedia.org/wiki/PBKDF2
// https://www.reddit.com/r/golang/comments/ad4ap6/go_implementation_of_pbkdf2/
//
// Convert to hex: https://stackoverflow.com/a/56091798
// Encode hex to string: https://stackoverflow.com/a/24269558
func PBDKF2(password string, salt string) string {
	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 128, sha256.New)
	return fmt.Sprintf("%x", dk)
}

func GetSHA256(str string) []byte {
	// https://gobyexample.com/sha256-hashes
	hash := crypto.SHA256.New()
	hash.Write([]byte(str))

	return hash.Sum(nil)
}

func GetSHA256asInt(str string) int {
	return BytesToInt(GetSHA256(str))
}

// EncodeToBase64 encodes an object to base64 format.
// If the encode process was unsuccessful, then return error
func EncodeToBase64(v interface{}) (string, error) {
	// https://stackoverflow.com/a/63126657
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}

	err = encoder.Close()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// DecodeFromBase64 decodes a string to an object.
func DecodeFromBase64(v interface{}, enc string) error {
	// https://stackoverflow.com/a/63126657
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}
*/

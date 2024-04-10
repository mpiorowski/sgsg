package system

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func GenerateRandomState(length int) (string) {
	// Generate random bytes
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
        panic(err)
	}

	// Encode the random bytes to a base64 URL-safe string
	state := base64.URLEncoding.EncodeToString(randomBytes)

	return state
}

func GenerateRandomString(length int) (string, error) {
	// Generate random bytes
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a hex string
	hex := hex.EncodeToString(randomBytes)
	return hex, nil
}

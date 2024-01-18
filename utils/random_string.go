package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString(length int) (string, error) {
	// Determine the required byte length based on the string length
	byteLength := length
	if length%4 != 0 {
		byteLength = length + (4 - (length % 4))
	}

	// Generate random bytes
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode random bytes to base64 string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Truncate or pad the string to the desired length
	randomString = randomString[:length]

	return randomString, nil
}

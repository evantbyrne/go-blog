package util

import "crypto/sha512"
import "encoding/base64"

const iterations = 17
const salt = "my$supe^rsecre5ta_wesomeco1olsalt"

// Hashed strings are always 88 characters in length
func Hash(value string) string {
	for i := 0; i < iterations; i++ {
		bytes := []byte(value + salt)
		hasher := sha512.New()
		hasher.Write(bytes)
		value = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	}

	return value
}
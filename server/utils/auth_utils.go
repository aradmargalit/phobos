package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// RandomToken returns a random base64 encoded string
func RandomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

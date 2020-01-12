package main

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

func GeneratePlayerToken(username string) string {
	hasher := sha1.New()
	hasher.Write([]byte(username + time.Now().String()))
	token := hex.EncodeToString(hasher.Sum(nil))
	return token
}

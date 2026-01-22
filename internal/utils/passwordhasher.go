package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))

	log.Println("Hashed password:", hex.EncodeToString(hash[:]))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(hashedPassword, password string) bool {
	return hashedPassword == HashPassword(password)
}
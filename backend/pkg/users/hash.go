package users

import (
	"golang.org/x/crypto/bcrypt"
)

const HASH_COST = 14

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
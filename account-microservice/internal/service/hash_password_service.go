package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPassword(password string, hashPasswordFromDb string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasswordFromDb), []byte(password))
	return err == nil //if equal => true otherwise false
}

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(data string) (string, error){
	hash, err :=  bcrypt.GenerateFromPassword([]byte(data), 13)
	return string(hash), err
}

func CheckHash (hash, data string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)) == nil
}
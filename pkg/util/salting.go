package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Gneration the salt sercert
func HashAndSalt(pwd string) string {
	passwordByte := []byte(pwd)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

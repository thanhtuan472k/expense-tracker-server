package auth

import (
	internalconstant "expense-tracker-server/internal/constant"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword ...
func HashedPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), internalconstant.PasswordHashingCost)
	return string(bytes)
}

// CompareHashedPassword ...
func CompareHashedPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

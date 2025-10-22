package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	// Generate hash with cost 14 (good balance of security and speed)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword compares plain password with hashed password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

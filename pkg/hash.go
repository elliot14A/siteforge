package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func VerifyPassword(hashedPassword, password string) error {
	bcErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return bcErr
}

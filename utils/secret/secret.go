package secret

import (
	"crypto/rand"
	"crypto/sha256"
)

// GenerateSalt 生成盐
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err

}

// HashWithSalt 加盐、加密
func HashWithSalt(password string, salt []byte) []byte {
	salted := append(salt, []byte(password)...)
	hashed := sha256.Sum256(salted)
	return hashed[:]
}

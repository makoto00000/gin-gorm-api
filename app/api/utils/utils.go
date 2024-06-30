package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// ハッシュ化したパスワードを生成
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// ハッシュ化したパスワードを検証
func CheckPasswordHash(hashedPassword string, inputPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
    return err == nil
}

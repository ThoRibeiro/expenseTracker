
package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

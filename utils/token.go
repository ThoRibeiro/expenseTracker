
package utils

import (
    "expensetracker/config"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func GenerateToken(userID uint, isAdmin bool) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "is_admin": isAdmin,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.GetEnv("JWT_SECRET", "secret")))
}

package utils

import (
	"expensetracker/config"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateTokenAndParse(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	tokenStr, err := GenerateToken(42, true)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET", "")), nil
	})
	if err != nil {
		t.Fatalf("Parse token error: %v", err)
	}
	if !token.Valid {
		t.Error("Token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("Claims are not MapClaims")
		return
	}

	if uint64(claims["user_id"].(float64)) != 42 {
		t.Errorf("Expected user_id 42, got %v", claims["user_id"])
	}
	if !claims["is_admin"].(bool) {
		t.Error("Expected is_admin to be true")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		t.Error("exp claim missing or not a number")
	} else if time.Unix(int64(exp), 0).Before(time.Now()) {
		t.Error("Token already expired")
	}
}

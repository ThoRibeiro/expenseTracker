package utils

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	pwd := "secret123"
	hash, err := HashPassword(pwd)
	if err != nil {
		t.Fatalf("HashPassword error: %v", err)
	}
	if hash == pwd {
		t.Error("Expected hashed password to differ from the original")
	}
	if !CheckPasswordHash(pwd, hash) {
		t.Error("CheckPasswordHash failed for correct password")
	}
	if CheckPasswordHash("wrong", hash) {
		t.Error("CheckPasswordHash succeeded for incorrect password")
	}
}

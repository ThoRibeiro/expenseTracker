package controllers_test

import (
	"bytes"
	"encoding/json"
	"expensetracker/database"
	"expensetracker/models"
	"expensetracker/routes"
	"expensetracker/seeds"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	database.DB = db
	if err := database.DB.AutoMigrate(&models.User{}, &models.Expense{}); err != nil {
		t.Fatalf("auto migrate error: %v", err)
	}
	seed.SeedAdmin()
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	routes.SetupRoutes(r)
	return r
}

func TestRegisterAndLogin(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	// Test Register
	w := httptest.NewRecorder()
	regBody := map[string]string{"name": "Bob", "email": "bob@example.com", "password": "password1"}
	jsonReg, _ := json.Marshal(regBody)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonReg))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 Created, got %d: %s", w.Code, w.Body.String())
	}
	var regResp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &regResp); err != nil {
		t.Fatalf("Unmarshal register response error: %v", err)
	}
	if _, ok := regResp["token"]; !ok {
		t.Error("Token not returned on registration")
	}

	// Test Login
	w2 := httptest.NewRecorder()
	loginBody := map[string]string{"email": "bob@example.com", "password": "password1"}
	jsonLogin, _ := json.Marshal(loginBody)
	req2, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	req2.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d: %s", w2.Code, w2.Body.String())
	}
	var loginResp map[string]string
	if err := json.Unmarshal(w2.Body.Bytes(), &loginResp); err != nil {
		t.Fatalf("Unmarshal login response error: %v", err)
	}
	if _, ok := loginResp["token"]; !ok {
		t.Error("Token not returned on login")
	}
}

package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)


func LoadCredentials() (string, string) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	stuID := os.Getenv("STU_ID")
	password := os.Getenv("PASSWORD")

	return stuID, password
}

func TestLogin(t *testing.T) {
	as := new(AuthSvc)
	
	t.Run("normal",func(t *testing.T) {
		stuID, password := LoadCredentials()
		if stuID == "" || password == "" {
			t.Fatal("STU_ID or PASSWORD not set in .env file")
		}

		err := as.Login(t.Context(), stuID, password)
		assert.NoError(t, err, "Login should succeed with valid credentials")
	})
	t.Run("err info", func(t *testing.T) {
		err := as.Login(t.Context(), "admin", "12313214")
		assert.Error(t, err, "Login should fail with empty credentials")
	})
}
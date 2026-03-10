package auth

import (
	"fmt"
	"time"

	"github.com/ros-e/sakura/packages/api/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func CreateProfile(username, password, configPath string) error {
	if username == "" {
		return fmt.Errorf("username must be provided")
	}
	if password == "" {
		return fmt.Errorf("password must be provided")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	cfg := config.Default()
	cfg.Profile.Username = username
	cfg.Profile.PasswordHash = string(hash)
	cfg.Profile.CreatedAt = time.Now()

	return config.Save(cfg, configPath)
}

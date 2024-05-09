package database

import (
	"idea-box/models"
	"idea-box/tolog"
)

func Migrate() error {
	err := DbEngine.AutoMigrate(&models.UserTable{})
	if err != nil {
		tolog.Log().Errorf("Migrate table user %e:", err).PrintAndWriteSafe()
		return err
	}
	err = DbEngine.AutoMigrate(&models.IdeaTable{})
	if err != nil {
		tolog.Log().Errorf("Migrate idea user %e:", err).PrintAndWriteSafe()
		return err
	}
	return nil
}

package utils

import (
	"idea-box/tolog"
	"os"
)

func GetProjectDirRoot() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		tolog.Log().Errorf("Error while GetProjectDirRoot %e", err).PrintAndWriteSafe()
		return "", err
	}
	return rootDir, nil
}

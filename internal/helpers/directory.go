package helpers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", errors.New("failed to get home directory: " + err.Error())
		}
		xdgConfigHome = filepath.Join(homeDir, ".config")
	}

	return xdgConfigHome, nil
}

func CreateConfigDir(configDir string, clusterName *string) (string, error) {
	defaultClusterName := "my-ente"

	if clusterName == nil {
		clusterName = &defaultClusterName
	}

	appConfigDir := filepath.Join(configDir, *clusterName)
	err := os.MkdirAll(appConfigDir, 0755)

	if err != nil {
		fmt.Println("Error creating config directory:", err)
		return "", err
	}

	fmt.Println("Created config directory at:", appConfigDir)
	return appConfigDir, nil
}

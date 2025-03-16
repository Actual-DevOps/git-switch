package config

import (
	"fmt"
	"os"
)

const (
	ConfigFileName = ".git-switch.conf"
)

func ReadConfig() (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", os.Getenv("HOME"), ConfigFileName))
	if err != nil {
		return "", fmt.Errorf("Can't read config file: %v", err)
	}

	return string(data), nil
}

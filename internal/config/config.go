package config

import (
	"fmt"
	"os"

	"github.com/Actual-DevOps/git-switch/internal/git"
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

func LoadAndValidateConfig() error {
	conf, err := ReadConfig()
	if err != nil {
		return fmt.Errorf("Error read config file: %v", err)
	}

	validConfig, err := IsValidConfig(conf)
	if err != nil {
		return fmt.Errorf("Error in IsValidConfig: %v", err)
	}

	if !validConfig {
		return fmt.Errorf("Config file is not valid!\n 'user.name', 'user.email' must be set")
	}

	return nil
}

func IsValidConfig(conf string) (bool, error) {
	profiles, err := git.GetProfiles(conf)
	if err != nil {
		return false, fmt.Errorf("Error get profile: %v", err)
	}

	for i := range profiles {
		gitProfiles, ok := profiles[i].(map[string]any)
		if !ok {
			return false, fmt.Errorf("Error parse profiles.git: %v", err)
		}

		gitProfilesValues, ok := gitProfiles["git"].(map[string]any)
		if !ok {
			return false, fmt.Errorf("Error parse profiles.git: %v", err)
		}

		if gitProfilesValues["user.name"] == nil || gitProfilesValues["user.email"] == nil {
			return false, nil
		}
	}

	return true, nil
}

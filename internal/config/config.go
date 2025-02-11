package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Profiles []Profiles `yaml:"profiles"`
}

type Core struct {
	Hookspath string `yaml:"hookspath"`
}

type User struct {
	Username   string `yaml:"username"`
	Email      string `yaml:"email"`
	Signingkey string `yaml:"signingkey"`
}

type Commit struct {
	GPGsign string `yaml:"gpgsign"`
}

type GPG struct {
	Program string `yaml:"program"`
}

type Git struct {
	User   User   `yaml:"user"`
	Commit Commit `yaml:"commit"`
	GPG    GPG    `yaml:"gpg"`
	Core   Core   `yaml:"core"`
}

type Profiles struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Git         Git    `yaml:"git"`
}

func LoadConfig(config *Config) error {
	if os.Getenv("GSW_CONFIG") == "" {
		if err := os.Setenv("GSW_CONFIG", "config.yaml"); err != nil {
			return fmt.Errorf("Can't set path for GSW_CONFIG: %v", err)
		}
	}

	viper.SetConfigFile(os.Getenv("GSW_CONFIG"))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("configuration read error: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("configuration parsing error: %v", err)
	}

	return nil
}

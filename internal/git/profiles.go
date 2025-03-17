package git

import (
	"fmt"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	GreenColor = "\033[32m"
	ResetColor = "\033[0m"
)

func SetProfile(profilesGit map[string]any, global bool) error {
	for key, v := range profilesGit {
		value, ok := v.(string)
		if !ok {
			return fmt.Errorf("Error parse value %v wrong type", v)
		}

		var cmd *exec.Cmd
		if global {
			cmd = exec.Command("git", "config", key, fmt.Sprintf("\"%s\"", value), "--global")
		} else {
			cmd = exec.Command("git", "config", key, fmt.Sprintf("\"%s\"", value))
		}

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error set git config: %w", err)
		}
	}

	return nil
}

func GetProfiles(conf string) (gitProfiles []any, err error) {
	var profiles map[string]any

	if err := yaml.Unmarshal([]byte(conf), &profiles); err != nil {
		return nil, fmt.Errorf("Error unmarshal config file: %w", err)
	}

	gitProfiles, ok := profiles["profiles"].([]any)
	if !ok {
		return nil, fmt.Errorf("Unknow type for profiles")
	}

	return gitProfiles, nil
}

func IsCurrentProfile(userEmail string) (bool, error) {
	cmd := exec.Command("git", "config", "user.email")

	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("Can't run git execute: %w", err)
	}

	return strings.TrimSpace(string(out)) == fmt.Sprintf("\"%s\"", strings.TrimSpace(userEmail)), nil
}

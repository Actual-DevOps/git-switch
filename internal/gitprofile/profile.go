package gitprofile

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Actual-DevOps/git-switch/internal/config"
)

func IsCurrentProfile(email string) (bool, error) {
	cmd := exec.Command("git", "config", "--get", "user.email")

	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("Can't get current profile: %w", err)
	}

	if strings.TrimSuffix(string(output), "\n") == email {
		return true, nil
	}

	return false, nil
}

func SetProfile(profiles []config.Profiles, profile int) {
	fmt.Println(profiles[profile].Git)
}

func IsProfileNumber(profile any) (bool, error) {
	ok, err := regexp.Match("^[0-9]+$", []byte(profile.(string)))
	if err != nil {
		return ok, fmt.Errorf("Error match regex: %w", err)
	}

	return ok, nil
}

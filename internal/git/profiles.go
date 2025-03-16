package git

import (
	"fmt"
	"os/exec"

	"gopkg.in/yaml.v3"
)

func SetProfile(profilesGit  map[string]any, global bool) error {
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
			return fmt.Errorf("Error set git config: %v", err)
		}
	}

	return nil
}

func GetProfiles(conf string) (gitProfiles []any, err error) {
	var profiles map[string]any

	if err := yaml.Unmarshal([]byte(conf), &profiles); err != nil {
		return nil, fmt.Errorf("Error unmarshal config file: %v", err)
	}

	gitProfiles, ok := profiles["profiles"].([]any)
	if !ok {
		return nil, fmt.Errorf("Unknow type for profiles")
	}

	return gitProfiles, nil
}

// func IsCurrentProfile(profiles map[string]any)  {

// }

// func IsValidConfig(conf string) error {
// 	profiles, err := GetProfiles(conf)
// 	if err != nil {
// 		return fmt.Errorf("Error get profile: %v", err)
// 	}

// 	gitProfiles, ok := profiles[i].(map[string]any)
// 	if !ok {
// 		return fmt.Errorf("Error parse profiles.git: %v", err)
// 	}

// 	for i := range gitProfiles {
// 		fmt.Printf("%s\n", i)
// 	}

// 	return nil
// }

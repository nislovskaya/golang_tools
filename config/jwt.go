package config

import (
	"fmt"
)

func GetSecret() (string, error) {
	secretKey, err := GetConfigValue("JWT_SECRET")
	if err != nil {
		return "", fmt.Errorf("failed to get `JWT_SECRET`: %w", err)
	}

	return secretKey, nil
}

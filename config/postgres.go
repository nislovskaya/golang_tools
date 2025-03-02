package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(logger *logrus.Entry) (*gorm.DB, error) {
	host, err := GetConfigValue("DB_HOST")
	if err != nil {
		return nil, fmt.Errorf("failed to get `DB_HOST`: %w", err)
	}

	user, err := GetConfigValue("DB_USER")
	if err != nil {
		return nil, fmt.Errorf("failed to get `DB_USER`: %w", err)
	}

	password, err := GetConfigValue("DB_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("failed to get `DB_PASSWORD`: %w", err)
	}

	dbname, err := GetConfigValue("DB_NAME")
	if err != nil {
		return nil, fmt.Errorf("failed to get `DB_NAME`: %w", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	logger.Infof("Database URL: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect  database: %w", err)
	}
	return db, nil
}

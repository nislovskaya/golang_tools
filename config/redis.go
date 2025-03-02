package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func ConnectRedis(logger *logrus.Entry) (*redis.Client, error) {
	host, err := GetConfigValue("REDIS_HOST")
	if err != nil {
		return nil, fmt.Errorf("failed to get `REDIS_HOST`: %w", err)
	}

	port, err := GetConfigValue("REDIS_PORT")
	if err != nil {
		return nil, fmt.Errorf("failed to get `REDIS_PORT`: %w", err)
	}

	user, err := GetConfigValue("REDIS_USER")
	if err != nil {
		return nil, fmt.Errorf("failed to get `REDIS_USER`: %w", err)
	}

	pass, err := GetConfigValue("REDIS_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("failed to get `REDIS_PASSWORD`: %w", err)
	}

	redisAddr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: user,
		Password: pass,
	})
	logger.Infof("Connected to redis: %s", redisAddr)

	if err = client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return client, nil
}

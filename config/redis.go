package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func ConnectRedis(logger *logrus.Entry) *redis.Client {
	host, err := GetConfigValue("REDIS_HOST")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	port, err := GetConfigValue("REDIS_PORT")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	user, err := GetConfigValue("REDIS_USER")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	pass, err := GetConfigValue("REDIS_PASSWORD")
	if err != nil {
		panic(fmt.Sprintf("failed to get config value %s", err.Error()))
	}

	redisAddr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: user,
		Password: pass,
	})
	logger.Infof("Connected to redis: %s", redisAddr)

	if err = client.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %s", err.Error()))
	}

	return client
}

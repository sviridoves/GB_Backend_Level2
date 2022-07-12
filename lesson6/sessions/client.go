package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	*redis.Client
	TTL time.Duration
}

func NewRedisClient(host, port string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("try to ping to redis: %w", err)
	}

	c := &RedisClient{
		Client: client,
	}

	return c, nil
}

func (c *RedisClient) Close() error {
	return c.Client.Close()
}

func (c *RedisClient) GetRecord(mkey string) ([]byte, error) {
	data, err := c.Get(context.Background(), mkey).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

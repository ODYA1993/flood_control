package config

import "time"

const (
	redisAddr = "localhost:6379"
	window    = 10 * time.Second
	limit     = 10
)

type Config struct {
	RedisAddr string
	Window    time.Duration
	Limit     int
}

func New() *Config {
	return &Config{
		RedisAddr: redisAddr,
		Window:    window,
		Limit:     limit,
	}
}

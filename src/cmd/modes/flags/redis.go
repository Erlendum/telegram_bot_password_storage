package flags

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisFlags struct {
	Host          string `mapstructure:"host"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Port          string `mapstructure:"port"`
	DB            int    `mapstructure:"db"`
	ExpireSeconds int    `mapstructure:"expireSeconds"`
}

func (f *RedisFlags) Init() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", f.Host, f.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: f.Password,
		DB:       f.DB,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}

package config

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Config struct {
	Redis RedisConfig
}

func GetRedisClient(lgr *log.Logger, config *RedisConfig) *redis.Client {
	lgr.Println("info: getting redis client")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})
	return rdb
}

func GetAppConfig(filename, path string) *Config {
	return loadConfig(filename, path)
}

func loadConfig(filename, path string) *Config {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("fatal: error reading config file")
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("fatal: error reading config variable")
	}
	return &conf
}

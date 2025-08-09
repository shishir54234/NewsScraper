package config

import (
	"fmt"
	"os"
	"strconv"
	"gopkg.in/yaml.v3"
)

func Load(configPath string) (*Config, error) {
	config := &Config{}
	if configPath == "" {
		configPath = "config.yaml"
	}else{
		file,err:=os.Open(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %v", err)
		}
		defer file.Close()
		err = yaml.NewDecoder(file).Decode(config)
		if err != nil {
			return nil, fmt.Errorf("failed to decode config file: %v", err)
		}
	}
	config.overrideWithEnv()
	return config, nil
}

func (c* Config) overrideWithEnv() {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		c.Server.Port = port
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		c.Server.Host = host
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		c.Database.Host = dbHost
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if portInt, err := strconv.ParseInt(dbPort,10,64); err == nil {
			c.Database.Port = portInt
		}
	}
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		c.Redis.Host = redisHost
	}
	if redisPort := os.Getenv("REDIS_PORT"); redisPort != "" {
		if redisPortInt, err := strconv.ParseInt(redisPort,10,64); err == nil {
			c.Redis.Port = redisPortInt
		}
		
	}
	if queueType := os.Getenv("QUEUE_TYPE"); queueType != "" {
		c.Queue.Type = queueType
	}
	if aiProvider := os.Getenv("AI_PROVIDER"); aiProvider != "" {
		c.Ai.Provider = aiProvider
	}
}

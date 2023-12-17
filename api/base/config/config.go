package config

import (
	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Web  *WebConfig      `yaml:"server"`
	DB   *DatabaseConfig `yaml:"db"`
	Auth *ServiceAuth    `yaml:"auth"`
}

type WebConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ServiceAuth struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func GetConfig() *Config {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logger.Error(lue.New("Could not get current file path", lue.ErrorWriter))
	}

	path := filepath.Join(filepath.Dir(filename), "yaml", "local.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error(lue.New(err.Error(), lue.ErrorReader))
	}
	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Error(lue.New(err.Error(), lue.ErrorDecoder))
	}
	return &config
}

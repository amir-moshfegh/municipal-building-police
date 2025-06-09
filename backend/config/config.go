package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	Server ServerConfig
}

func GetConfig() (*Config, error) {
	filename, fileType, err := configFilePath()
	if err != nil {
		return nil, err
	}
	if !isFileExists(filename) {
		return nil, errors.New("config file not found")
	}

	v, err := readConfigFile(filename, fileType)
	if err != nil {
		return nil, errors.New("config file read error")
	}

	cfg, err := parsConfigFile(v)

	return cfg, err
}

func parsConfigFile(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, nil
	}
	return &cfg, nil
}

func readConfigFile(fn, ft string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(fn)
	v.SetConfigType(ft)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func configFilePath() (filename, fileType string, err error) {
	filename = os.Getenv("APP_CONFIG_FILE")
	if filename == "" {
		filename = "./config/config.yml"
		fileType = "yml"
	} else {
		fileInfo := strings.Split(filename, ".")
		fileExt := fileInfo[len(fileInfo)-1]
		fileType = fileExt
	}

	if fileType == "" {
		err = errors.New("invalid config file extension")
	}

	return
}

func isFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

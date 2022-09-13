package common

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/spf13/viper"
)

type Config struct {
	Db     Database     `yaml:"Database"`
	Server ServerConfig `yaml:"Server"`
}

type ServerConfig struct {
	ListenAddr string `yaml:"ListenAddr"`
}

type Database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Schema   string `yaml:"Schema"`
}

var globalConfig Config

func GlobalConfig() *Config {
	return &globalConfig
}

func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file error:%w", err)
	}

	configDecoder := func(d *mapstructure.DecoderConfig) {
		d.TagName = "yaml"
	}
	opts := []viper.DecoderConfigOption{configDecoder}
	if err := viper.Unmarshal(&globalConfig, opts...); err != nil {
		return fmt.Errorf("unmarshal config error:%w", err)
	}
	return nil
}

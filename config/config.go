package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BaseDir string `mapstructure:"BASE_DIR"`
}

func LoadConfig() (Config, error) {
	v := viper.New()

	v.SetDefault("PORT", ":3000")
	v.SetDefault("UPLOAD_DIR", "./base_folder")
	v.SetDefault("BODY_LIMIT", 50*1024*1024) // 50 MB

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, err
		}
	}

	v.AutomaticEnv() // real env vars override .env

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil

}

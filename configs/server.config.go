package configs

import (
	"path/filepath"

	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (*conf, error) {
	envFile := filepath.Join(".", ".env")

	viper.SetConfigType("env")
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	checkRequiredEnvs()

	return cfg, nil
}

func checkRequiredEnvs() {
	requiredEnvs := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"SERVER_PORT",
	}

	for _, env := range requiredEnvs {
		if viper.GetString(env) == "" {
			panic(env + " is required")
		}
	}
}

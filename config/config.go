package config

import (
	"errors"

	"github.com/spf13/viper"
)

var ErrReadingConfigFile = errors.New("error reading config file")

type Config struct {
	DbUri string
}

var config = Config{}
var emptyConfig = Config{}

func GetConfig() Config {
	return config
}

func SetConfig(c Config) {
	config = c
}

func ReadConfigFile() (Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return emptyConfig, errors.Join(ErrReadingConfigFile, err)
	}

	return ReadConfig(), nil
}

func ReadConfig() Config {
	config = Config{
		DbUri: viper.GetString("db.uri"),
	}

	return config
}

func init() {
	// viper.SetEnvPrefix("ENVPREFIX")
	viper.SetDefault("db.uri", "")
	viper.AutomaticEnv()

	// viper.SetConfigName("configfilename")
	// viper.SetConfigType("toml")
	// viper.AddConfigPath(".")
}

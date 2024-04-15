package config

import (
	"errors"

	"github.com/adrg/xdg"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func init() {
	// viper.SetEnvPrefix("ENVPREFIX")
	viper.SetDefault("db.uri", "")
	viper.AutomaticEnv()

	// viper.SetConfigName("configfilename")
	viper.SetConfigType("toml")

}

var ErrReadingConfigFile = errors.New("error reading config file")

type Config struct {
	DbUri string
}

// ReadConfigFile Reads the config file from one of the sources
//
// Tries to read the config from:
// 1. $XDG_CONFIG_HOME/application/application.toml
// 2. $HOME/.application
// 3. <curren path>/application.toml
func ReadConfigFile() error {
	ok, err := readFromXdg()
	if ok {
		return nil
	}
	if err != nil {
		return errors.Join(ErrReadingConfigFile, err)
	}

	ok, err = readFromHome()
	if ok {
		return nil
	}
	if err != nil {
		return errors.Join(ErrReadingConfigFile, err)
	}

	ok, err = readFromCurrentDir()
	if ok {
		return nil
	}
	if err != nil {
		return errors.Join(ErrReadingConfigFile, err)
	}

	return nil
}

// Get Creates a configured Config
func Get() Config {
	config := Config{
		DbUri: viper.GetString("db.uri"),
	}

	return config
}

func readFromXdg() (bool, error) {
	xdgConfig, err := xdg.ConfigFile("application/application.toml")
	if err != nil {
		return false, err
	}

	viper.SetConfigFile(xdgConfig)

	return tryToRead()
}

func readFromHome() (bool, error) {
	home, err := homedir.Dir()
	if err != nil {
		return false, err
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".application")

	return tryToRead()
}

func readFromCurrentDir() (bool, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("application")

	return tryToRead()
}

func tryToRead() (bool, error) {
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

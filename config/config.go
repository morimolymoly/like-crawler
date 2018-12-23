package config

import (
	"fmt"
	"os"
	"path/filepath"

	gh "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var singleton *Config

// Config ... config object
type Config struct {
	ConsumerKey    string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	AccessToken    string `yaml:"accessToken"`
	AccessSecret   string `yaml:"accessSecret"`
}

// GetInstance ... Get config object
func GetInstance() *Config {
	return singleton
}

// SetConfigs ... set configs
func (c *Config) SetConfigs(consumerKey string, consumerSecret string, accessToken string, accessSecret string) error {
	c.ConsumerKey = consumerKey
	c.ConsumerSecret = consumerSecret
	c.AccessToken = accessToken
	c.AccessSecret = accessSecret

	viper.Set(cKEY, consumerKey)
	viper.Set(aTOKEN, accessToken)
	viper.Set(cSECRET, consumerKey)
	viper.Set(aSECRET, accessToken)

	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// ReadConfig ... read cofig from configfile
func (c *Config) ReadConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	c.ConsumerKey = viper.GetString(cKEY)
	c.AccessToken = viper.GetString(aTOKEN)
	c.ConsumerSecret = viper.GetString(cSECRET)
	c.AccessSecret = viper.GetString(aSECRET)
	return nil
}

// CheckConfig ... check config
func CheckConfig() error {
	if singleton.ConsumerKey != "" && singleton.AccessToken != "" {
		return nil
	}
	return fmt.Errorf("You must need to set consumer key and access token by set-subcommand")
}

func getConfigfilePath() (string, error) {
	hd, err := gh.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(hd, dPATH), nil
}

func createConfigFile(dp string) error {
	f, err := os.Create(dp)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

// Init ... initialize config
func Init() error {
	dp, err := getConfigfilePath()
	if err != nil {
		return err
	}

	viper.SetConfigFile(dp)
	viper.SetConfigType(dTYPE)

	_, err = os.Stat(dp)
	if err != nil {
		if err := createConfigFile(dp); err != nil {
			return err
		}
	}

	singleton = &Config{}
	return nil
}

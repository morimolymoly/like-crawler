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
	SinceID        int64  `yaml:"sinceID"`
	SavePath       string `yaml:"savePath"`
}

// GetInstance ... Get config object
func GetInstance() *Config {
	return singleton
}

// UpdateSavePath ... update savepath
func (c *Config) UpdateSavePath(savePath string) error {
	c.ReadConfig()
	c.SavePath = savePath
	viper.Set(sPATH, savePath)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// UpdateSinceID ... update sinceID
func (c *Config) UpdateSinceID(sinceID int64) error {
	c.ReadConfig()
	c.SinceID = sinceID
	viper.Set(sID, sinceID)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// UpdateCreds ... Update creds
func (c *Config) UpdateCreds(consumerKey string, consumerSecret string, accessToken string, accessSecret string) error {
	c.ReadConfig()
	c.ConsumerKey = consumerKey
	c.ConsumerSecret = consumerSecret
	c.AccessToken = accessToken
	c.AccessSecret = accessSecret

	viper.Set(cKEY, consumerKey)
	viper.Set(aTOKEN, accessToken)
	viper.Set(cSECRET, consumerSecret)
	viper.Set(aSECRET, accessSecret)

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
	c.SavePath = viper.GetString(sPATH)
	c.SinceID = viper.GetInt64(sID)
	return nil
}

// CheckConfig ... check config
func CheckConfig() error {
	if singleton.ConsumerKey == "" && singleton.AccessToken == "" {
		return fmt.Errorf("You must need to set consumer key and access token by setcreds subcommand")
	}
	if singleton.SavePath == "" {
		return fmt.Errorf("You must need to set savePath by setsavepath subcommand")
	}
	return nil
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
	viper.SetDefault(sID, 0)
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

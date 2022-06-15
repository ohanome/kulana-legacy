package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"kulana/l"
	"kulana/setup"
	"os"
)

type Config struct {
	Mail MailConfig `json:"mail"`
}

type MailConfig struct {
	StatusCodes  string `json:"status_codes"`
	Subject      string `json:"subject"`
	TemplateFile string `json:"template_file"`
}

var defaultConfig = Config{
	Mail: MailConfig{
		StatusCodes:  "4xx,5xx",
		Subject:      "Host $HOST is $STATUS",
		TemplateFile: setup.GetSetupDir() + "/mail.html",
	},
}

func init() {
	viper.SetConfigFile(setup.GetConfigFile())
	SafeDefaults(true)
}

func Get(key string) interface{} {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		l.Emergency(err.Error())
	}
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		l.Emergency(err.Error())
	}
}

func SafeDefaults(warnOnExist bool) {
	configFile := setup.GetConfigFile()
	_, err := os.Stat(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			l.Emergency(err.Error())
		} else {
		}
	} else {
		if warnOnExist {
			l.Warning(fmt.Sprintf("Config file already exists and will not be overwritten. If you want to refresh it, delete the old file at %s.", setup.GetConfigFile()))
		}
		return
	}

	configJson, err := json.Marshal(defaultConfig)
	if err != nil {
		l.Emergency(err.Error())
	}

	err = viper.ReadConfig(bytes.NewBuffer(configJson))
	if err != nil {
		l.Emergency(err.Error())
	}

	viper.SetConfigFile(setup.GetConfigFile())
	err = viper.WriteConfig()
	if err != nil {
		l.Emergency(err.Error())
	}
}

package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Set() {
	configFile, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal("error loading config file")

	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err = viper.ReadConfig(configFile)
	if err != nil {
		log.Fatal("error reading config")
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}

package cmd

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config/local.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")              // type of config file
	viper.AddConfigPath(".")                 // path to look for the config file in
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

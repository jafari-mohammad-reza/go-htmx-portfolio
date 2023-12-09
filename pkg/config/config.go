package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type PrivateConfig struct {
	Port int `mapstructure:"PORT"`
}

type Config struct {
	TextColor string `mapstructure:"text_color"`
}

var (
	privateConfig PrivateConfig
	config        Config
)

func SetupConfigs() {
	setupConfig()
	setupPrivateConfig()
}
func GetConfigs() (*PrivateConfig, *Config) {
	return &privateConfig, &config
}

func setupConfig() {
	viper.SetConfigFile(".config.json")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)

	}
	fmt.Printf("Config: %+v\n", config)
}

func setupPrivateConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&privateConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)

	}

	fmt.Printf("Config: %+v\n", privateConfig)

}

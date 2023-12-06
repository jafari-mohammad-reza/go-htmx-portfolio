package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port int `mapstructure:"PORT"`
}

func SetupConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return nil
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil
	}

	fmt.Printf("Config: %+v\n", config)
	return &config
}

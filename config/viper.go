package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// log.Panic(err)
		return
	}

	err = viper.Unmarshal(&config)
	return

}
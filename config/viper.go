package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

	viper.BindEnv("PUSHOVER_API_KEY")
	viper.BindEnv("PUSHOVER_USER_KEY")

	err = viper.Unmarshal(&config)
	return

}

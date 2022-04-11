package config

type Config struct {
	PushoverApiKey  string `mapstructure:"PUSHOVER_API_KEY"`
	PushoverUserKey string `mapstructure:"PUSHOVER_USER_KEY"`
}

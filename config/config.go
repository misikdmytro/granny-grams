package config

import "github.com/spf13/viper"

type Config struct {
	Environment   string
	OpenAIBaseURL string
	OpenAIToken   string
}

func LoadConfig() (*Config, error) {
	conf := viper.New()

	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")

	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = conf.Unmarshal(&c)

	return &c, err
}

package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (Config, error) {
	var err error
	var config Config

	viper.AddConfigPath(path) // define paths were viper will search for config file
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, err
}

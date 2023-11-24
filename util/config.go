package util

import "github.com/spf13/viper"

// Config is the application config
// Read by ENV variables
type Config struct {
	dbDriver      string `mapstructure:"DB_DRIVER"`
	dbSource      string `mapstructure:"DB_SOURCE"`
	serverAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // env, json, toml, yaml, yml, properties, props, prop

	viper.AutomaticEnv() // read from ENV variables

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

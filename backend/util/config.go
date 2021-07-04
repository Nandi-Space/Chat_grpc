package util

import (

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Host         string `mapstructure:"SERVER_ADDRESS"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	UserName     string `mapstructure:"USERNAME"`
	Password     string `mapstructure:"PASSWORD"`
	InitailCap   int    `mapstructure:"INITIALCAP"`
	MaxOpen      int    `mapstructure:"MAX_OPEN"`
	ReadTimeOut  int    `mapstructure:"READ_TIME_OUT"`
	TimeOut      int    `mapstructure:"TIME_OUT"`
	WriteTimeOut int    `mapstructure:"WRITE_TIME_OUT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

package common

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	MealsDB string `mapstructure:"meals_db"`

	Server struct {
		Port int
	}
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("default")
	v.AddConfigPath("./pkg/conf")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read config file: %s", err.Error())
	}

	var conf *Config
	conf = new(Config)
	if err := v.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config into struct: %s", err.Error())
	}

	return conf, nil
}

package config

import (
	"errors"

	errs "github.com/pkg/errors"

	"github.com/spf13/viper"
)

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, errs.WithStack(err)
	}
	return v, nil
}

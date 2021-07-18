package config

import (
	"github.com/spf13/viper"
)

type Config interface {
}

type config struct {
	v *viper.Viper
}

func New() Config {
	return &config{
		v: viper.New(),
	}
}

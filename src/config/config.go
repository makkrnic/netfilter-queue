package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	v       *viper.Viper
	Address string `json:"udpAddress"`
}

func New() *Config {
	return &Config{
		v: viper.New(),
	}
}

func Load(log *zap.Logger) (*Config, error) {
	log.Info("loading configuration")

	conf := New()
	conf.v.SetConfigType("json")
	conf.v.SetConfigName("config")
	conf.v.AddConfigPath(".")
	conf.v.AddConfigPath("/etc/netfilter-queue")
	conf.v.AddConfigPath("$HOME/netfilter-queue")
	if err := conf.v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := conf.v.Unmarshal(&conf); err != nil {
		return nil, err
	}

	conf.v.WatchConfig()
	conf.v.OnConfigChange(func(e fsnotify.Event) {
		if err := conf.v.Unmarshal(&conf); err != nil {
			log.Error("error unmarshaling config data", zap.Error(err))
		}
		log.Info("config changed, reloading config", zap.Any("config", &conf))
	})

	return conf, nil
}

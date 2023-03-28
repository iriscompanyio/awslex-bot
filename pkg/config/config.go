package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Aws Credentials
	BotAliasId     string `default:"aws"`
	BotId          string `default:"aws"`
	LocaleId       string `default:"aws"`
	SessionId      string `default:"aws"`
	AccesKeyId     string `default:"aws"`
	SecretAccesKey string `default:"aws"`
	SessionToken   string `default:"aws"`
	Region         string `default:"aws"`
	// Server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
}

var Cfg config

func LoadConfig() error {
	err := envconfig.Process("AWS", &Cfg)
	if err != nil {
		return err
	}
	return nil
}

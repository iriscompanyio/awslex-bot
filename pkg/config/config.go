package config

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Aws Credentials
	BotAliasId      string `envconfig:"BOTALIASID" default:"aws"`
	BotId           string `envconfig:"BOTID" default:"aws"`
	LocaleId        string `envconfig:"LOCALEID" default:"aws"`
	SessionId       string `envconfig:"SESSIONID" default:"aws"`
	AccessKeyId     string `envconfig:"ACCESSKEYID" default:"aws"`
	SecretAccessKey string `envconfig:"SECRETACCESSKEY" default:"aws"`
	SessionToken    string `envconfig:"SESSIONTOKEN" default:""`
	Region          string `envconfig:"REGION" default:"us-east-1"`
	// Server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
}

var Cfg config

func LoadConfig() error {
	err := envconfig.Process("aws", &Cfg)
	if err != nil {
		return err
	}
	format := "BotAliasID: %s\nBotId: %s\nLocaleId: %s\nSessionId: %s\nAccesKeyId: %s\nSecretAccesKey: %s\nSessionToken: %s\nRegion: %s\n"
	_, err = fmt.Printf(format, Cfg.BotAliasId, Cfg.BotId, Cfg.LocaleId, Cfg.SessionId, Cfg.AccessKeyId, Cfg.SecretAccessKey, Cfg.SessionToken, Cfg.Region)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

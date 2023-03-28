package bootstrap

import (
	"context"

	"github.com/iriscompanyio/awslex-bot/internal/platform/server"
	"github.com/iriscompanyio/awslex-bot/pkg/config"
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	ctx, srv := server.NewServer(context.Background(), config.Cfg.Host, config.Cfg.Port, config.Cfg.ShutdownTimeout)

	return srv.RunServer(ctx)
}

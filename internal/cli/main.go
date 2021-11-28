package cli

import (
	"context"
	"eth_service/internal/config"
	"eth_service/internal/service"
	"gitlab.com/distributed_lab/kit/kv"
)

func Run() {
	// TODO: add some commands from cli (ENTRYPOINT) if you gonna build with docker file

	cfg := config.NewConfig(kv.MustFromEnv())

	svc := service.New(cfg)
	svc.Run(context.Background())
}

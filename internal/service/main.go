package service

import (
	"context"
	"eth_service/internal/config"
	"eth_service/internal/service/event_listner"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
)

type Service struct {
	log *logan.Entry
	config config.Config
}

func New(cfg config.Config) *Service {
	return &Service {
		log: cfg.Log(),
		config: cfg,
	}
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Starting listening for events")
	s.spawn(ctx)  // go
}

func (s *Service) spawn(ctx context.Context) {
	conf, err := config.ParseConfig(kv.MustGetStringMap(kv.MustFromEnv(), "connection"))
	if err != nil{
		log.Fatal(err)
	}

	eventer := event_listner.New(config.Opts{
		Log: s.log,
		Conf: conf,
	})

	eventer.Run(ctx)  // go
}


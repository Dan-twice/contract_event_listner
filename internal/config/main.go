package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
)

type config struct {
	getter kv.Getter
	comfig.Logger
	pgdb.Databaser
}

func (c *config) Log() *logan.Entry {
	return logan.New()
}

type Config interface {
	Log() *logan.Entry
	pgdb.Databaser
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser: pgdb.NewDatabaser(getter),
	}
}

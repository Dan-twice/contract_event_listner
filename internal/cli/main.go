package cli

import (
	"eth_service/internal/config"
	"eth_service/internal/db"
	"eth_service/internal/service"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"context"
	"gitlab.com/distributed_lab/kit/kv"
)

/*
import (
	"context"
	"eth_service/internal/config"
	"eth_service/internal/db"
	"eth_service/internal/service"
	"fmt"
	"gitlab.com/distributed_lab/kit/kv"
	"gopkg.in/alecthomas/kingpin.v2"
)

func Run(args []string) bool {
	fmt.Println(args)

	cfg := config.NewConfig(kv.MustFromEnv())

	db.MigrateUp(cfg)
	svc := service.New(cfg)
	svc.Run(context.Background())

	return true
}*/

func Run(args []string) bool {
	fmt.Println("Commands input", args)

	app := kingpin.New("eth_service", "")
	runCmd := app.Command("run", "run command")
	eventer := runCmd.Command("eventer", "run eventer service")

	cfg := config.NewConfig(kv.MustFromEnv())
	log := cfg.Log()

	// first arg is the name of app itself
	cmd, err := app.Parse(args[1:])
	if err != nil{
		log.WithError(err).Error("failed to parse arguments")
	}

	switch cmd {
	case eventer.FullCommand():
		db.MigrateUp(cfg)
		svc := service.New(cfg)
		svc.Run(context.Background())
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}

	return true
}

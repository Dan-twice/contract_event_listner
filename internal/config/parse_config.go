package config

import (
	"eth_service/internal/transaction"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3"
)

type EventerInterface interface {
	Add(issuance *transaction.TransactionTestEvent) error
	GetById(id int64) EventerInterface
	Select() ([]Event, error)
}

type Event struct {
	Id        int64            `db:"id"`
	Works string `db:"works"`
}

type Opts struct {
	Eventer EventerInterface
	Log *logan.Entry
	Conf *Conf
}

type Conf struct {
	WebSocket string `fig:"ws"`
	ProjectID string `fig:"project"`
	ContractAddr string `fig:"contract"`
}

func ParseConfig(raw map[string]interface{}) (*Conf, error) {
	config := Conf{}

	err := figure.
		Out(&config).
		From(raw).
		Please()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}

	return &config, nil
}

type DBConfig struct {
	User string `fig:"pguser"`
	Password string `fig:"pgpassword"`
	Host string `fig:"pghost"`
	Port string `fig:"pgport"`
	DB string `fig:"pgdb"`
}

func ParseDB(raw map[string]interface{}) (*DBConfig, error) {
	config := DBConfig{}

	err := figure.
		Out(&config).
		From(raw).
		Please()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}

	return &config, nil
}

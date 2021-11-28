package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3"
)

type Opts struct {
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

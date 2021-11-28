package event_listner

import (
	"eth_service/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
)

type Service struct {
	log *logan.Entry
	client *ethclient.Client
	address common.Address
}

func New(opts config.Opts) *Service{
	infuraWebSocket := opts.Conf.WebSocket + opts.Conf.ProjectID
	client, err := ethclient.Dial(infuraWebSocket)  // NetworkURL + ProjectID
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(opts.Conf.ContractAddr)

	return &Service{
		log: opts.Log,
		client: client,
		address: address,
	}
}

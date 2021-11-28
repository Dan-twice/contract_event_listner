package event_listner

import (
	"context"
	"eth_service/internal/transaction"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"time"
)

// TODO: is this name important?
const runnerName = "transfer_test"
const runnerReal = "transfer"

func (s *Service) Run(ctx context.Context) {
	running.WithBackOff(ctx, s.log, runnerName, s.processTest,
		5*time.Second, 5*time.Second, 5*time.Second)
	running.WithBackOff(ctx, s.log, runnerReal, s.processReal,
		5*time.Second, 5*time.Second, 5*time.Second)
}

func (s *Service) processTest(ctx context.Context) error {
	contract, err := transaction.NewTransaction(s.address, s.client)
	if err != nil {
		return err
	}

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	loggesChannel := make(chan *transaction.TransactionTestEvent, 2)
	defer close(loggesChannel)
	sub, err := contract.WatchTestEvent(watchOpts, loggesChannel)
	if err != nil {
		return errors.Wrap(err, "failed to start watch TestEvent")
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <- ctx.Done():
			return nil
		case err := <- sub.Err():
			return err
		case event := <-loggesChannel:
			if event != nil {
				fmt.Println(event)
			}
		}
	}
}

func (s *Service) processReal(ctx context.Context) error {
	contract, err := transaction.NewTransaction(s.address, s.client)
	if err != nil {
		return err
	}

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	loggesChannel := make(chan *transaction.TransactionValueSender, 2)
	defer close(loggesChannel)
	sub, err := contract.WatchValueSender(watchOpts, loggesChannel)
	if err != nil {
		return errors.Wrap(err, "failed to start watch TestEvent")
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <- ctx.Done():
			return nil
		case err := <- sub.Err():
			return err
		case event := <-loggesChannel:
			if event != nil {
				fmt.Println(event)
			}
		}
	}
}

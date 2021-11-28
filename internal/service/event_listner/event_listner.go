package event_listner

import (
	"crypto/ecdsa"
	//"eth_meet/transaction"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PrivateK = "9194d6dd67829333170dda154e4e4a20e49b3e1612eb6490c48072db0d2132bb" // Account1
	//"dde886375bb3b61b48c7c7d4cf886a35d5fe35a93e42853ede18304eb762772f" // Account2
	ProjectID = "b9397e496e5541489195b3da512b912f" // Infura
	AccountAddr = "0x6d71AE8EE533AB5425f187A56A7CD13BF774cac0" // Account1
	NetworkURL = "https://rinkeby.infura.io/v3/"
	ContractAddr = "0x37cee7ae7309df851e7ba68ac7ebab39bdbafb9a" // Last Transfer
	// "0x4b390ab7831e9c065eb29cab50f4dd06267d8bb8" // Transfer
	//"0x33e0ca9fe03346c37b1f886fc1f3a025288d9f77" // Store
	// "0xce9f8aa7f622155735778115f1f43e1a917fb3a8"
)

func Foo() {
	// Connection (not get things) cross https is deprecated;
	// use WevSockets (ws) or IPC (default)
	infuraWebSocket := `wss://rinkeby.infura.io/ws/v3/` + ProjectID
	client, err := ethclient.Dial(infuraWebSocket)  // NetworkURL + ProjectID
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateK)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	//auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(10000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(ContractAddr)

	//instance, err := transaction.NewTransaction(address, client)
	//if err != nil {
	//	log.Println("NewStore prob")
	//	log.Fatal(err)
	//}

	fmt.Println("Contract is loaded")

	// EventListner1
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	var logs = make(chan types.Log, 2)
	s, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("Subscribe Failed: ", err)
		return
	}
	errChan := s.Err()

	for {
		select {
		case err := <-errChan:
			log.Fatal("Logs subscription error:", err)
		case l := <-logs:
			fmt.Println("new log:", l)
		}
	}

	//// EventListner2
	//// Watch for Watch<EventName> function
	//watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	//// Channel for responses
	//channel := make(chan *transaction.TransactionTestEvent)
	//
	//go func() {
	//	sub, err := instance.WatchTestEvent(watchOpts, channel)
	//	if err != nil {
	//		log.Println("Watch events prob:", err)
	//		return
	//	}
	//
	//	defer sub.Unsubscribe()
	//}()
	//
	//event := <-channel
	//fmt.Println(event)

	//// Test of contract
	//retr, err := instance.TestReturn(nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(retr)
}

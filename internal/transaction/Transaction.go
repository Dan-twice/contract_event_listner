// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transaction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TransactionMetaData contains all meta data concerning the Transaction contract.
var TransactionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"success\",\"type\":\"string\"}],\"name\":\"TestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"ValueSender\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TestEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testReturn\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040518060400160405280600681526020017f576f726b732100000000000000000000000000000000000000000000000000008152506000908051906020019061004f929190610062565b5034801561005c57600080fd5b50610166565b82805461006e90610134565b90600052602060002090601f01602090048101928261009057600085556100d7565b82601f106100a957805160ff19168380011785556100d7565b828001600101855582156100d7579182015b828111156100d65782518255916020019190600101906100bb565b5b5090506100e491906100e8565b5090565b5b808211156101015760008160009055506001016100e9565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061014c57607f821691505b602082108114156101605761015f610105565b5b50919050565b6103e1806101756000396000f3fe60806040526004361061002d5760003560e01c8063e13a771614610072578063f802454a1461009d5761006d565b3661006d577f9f1dd2a0b3bec3db71afa9520af20702e06a9b0d68c36d64a1380667129cf36b33346040516100639291906101d3565b60405180910390a1005b600080fd5b34801561007e57600080fd5b506100876100b4565b6040516100949190610295565b60405180910390f35b3480156100a957600080fd5b506100b2610142565b005b600080546100c1906102e6565b80601f01602080910402602001604051908101604052809291908181526020018280546100ed906102e6565b801561013a5780601f1061010f5761010080835404028352916020019161013a565b820191906000526020600020905b81548152906001019060200180831161011d57829003601f168201915b505050505081565b7fe75028ff36bb6473da3731a30e1aeeae9988e2415dba2c4e91e0357955065fba60405161016f90610364565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101a482610179565b9050919050565b6101b481610199565b82525050565b6000819050919050565b6101cd816101ba565b82525050565b60006040820190506101e860008301856101ab565b6101f560208301846101c4565b9392505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561023657808201518184015260208101905061021b565b83811115610245576000848401525b50505050565b6000601f19601f8301169050919050565b6000610267826101fc565b6102718185610207565b9350610281818560208601610218565b61028a8161024b565b840191505092915050565b600060208201905081810360008301526102af818461025c565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806102fe57607f821691505b60208210811415610312576103116102b7565b5b50919050565b7f4576656e7473206c6973746e657220776f726b73210000000000000000000000600082015250565b600061034e601583610207565b915061035982610318565b602082019050919050565b6000602082019050818103600083015261037d81610341565b905091905056fea2646970667358221220a810e29f1ab868035df49f37147116f1d13e7fa83f1ab0b1f0b34f805ea0151564736f6c637829302e382e31312d646576656c6f702e323032312e31312e32372b636f6d6d69742e6530633835633666005a",
}

// TransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionMetaData.ABI instead.
var TransactionABI = TransactionMetaData.ABI

// TransactionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransactionMetaData.Bin instead.
var TransactionBin = TransactionMetaData.Bin

// DeployTransaction deploys a new Ethereum contract, binding an instance of Transaction to it.
func DeployTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Transaction, error) {
	parsed, err := TransactionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// Transaction is an auto generated Go binding around an Ethereum contract.
type Transaction struct {
	TransactionCaller     // Read-only binding to the contract
	TransactionTransactor // Write-only binding to the contract
	TransactionFilterer   // Log filterer for contract events
}

// TransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionSession struct {
	Contract     *Transaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionCallerSession struct {
	Contract *TransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionTransactorSession struct {
	Contract     *TransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionRaw struct {
	Contract *Transaction // Generic contract binding to access the raw methods on
}

// TransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionCallerRaw struct {
	Contract *TransactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionTransactorRaw struct {
	Contract *TransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransaction creates a new instance of Transaction, bound to a specific deployed contract.
func NewTransaction(address common.Address, backend bind.ContractBackend) (*Transaction, error) {
	contract, err := bindTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// NewTransactionCaller creates a new read-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionCaller(address common.Address, caller bind.ContractCaller) (*TransactionCaller, error) {
	contract, err := bindTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionCaller{contract: contract}, nil
}

// NewTransactionTransactor creates a new write-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionTransactor, error) {
	contract, err := bindTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionTransactor{contract: contract}, nil
}

// NewTransactionFilterer creates a new log filterer instance of Transaction, bound to a specific deployed contract.
func NewTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionFilterer, error) {
	contract, err := bindTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionFilterer{contract: contract}, nil
}

// bindTransaction binds a generic wrapper to an already deployed contract.
func bindTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.TransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transact(opts, method, params...)
}

// TestReturn is a free data retrieval call binding the contract method 0xe13a7716.
//
// Solidity: function testReturn() view returns(string)
func (_Transaction *TransactionCaller) TestReturn(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Transaction.contract.Call(opts, &out, "testReturn")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TestReturn is a free data retrieval call binding the contract method 0xe13a7716.
//
// Solidity: function testReturn() view returns(string)
func (_Transaction *TransactionSession) TestReturn() (string, error) {
	return _Transaction.Contract.TestReturn(&_Transaction.CallOpts)
}

// TestReturn is a free data retrieval call binding the contract method 0xe13a7716.
//
// Solidity: function testReturn() view returns(string)
func (_Transaction *TransactionCallerSession) TestReturn() (string, error) {
	return _Transaction.Contract.TestReturn(&_Transaction.CallOpts)
}

// TestEvents is a paid mutator transaction binding the contract method 0xf802454a.
//
// Solidity: function TestEvents() returns()
func (_Transaction *TransactionTransactor) TestEvents(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "TestEvents")
}

// TestEvents is a paid mutator transaction binding the contract method 0xf802454a.
//
// Solidity: function TestEvents() returns()
func (_Transaction *TransactionSession) TestEvents() (*types.Transaction, error) {
	return _Transaction.Contract.TestEvents(&_Transaction.TransactOpts)
}

// TestEvents is a paid mutator transaction binding the contract method 0xf802454a.
//
// Solidity: function TestEvents() returns()
func (_Transaction *TransactionTransactorSession) TestEvents() (*types.Transaction, error) {
	return _Transaction.Contract.TestEvents(&_Transaction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionSession) Receive() (*types.Transaction, error) {
	return _Transaction.Contract.Receive(&_Transaction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionTransactorSession) Receive() (*types.Transaction, error) {
	return _Transaction.Contract.Receive(&_Transaction.TransactOpts)
}

// TransactionTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the Transaction contract.
type TransactionTestEventIterator struct {
	Event *TransactionTestEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TransactionTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionTestEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TransactionTestEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TransactionTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionTestEvent represents a TestEvent event raised by the Transaction contract.
type TransactionTestEvent struct {
	Success string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0xe75028ff36bb6473da3731a30e1aeeae9988e2415dba2c4e91e0357955065fba.
//
// Solidity: event TestEvent(string success)
func (_Transaction *TransactionFilterer) FilterTestEvent(opts *bind.FilterOpts) (*TransactionTestEventIterator, error) {

	logs, sub, err := _Transaction.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &TransactionTestEventIterator{contract: _Transaction.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0xe75028ff36bb6473da3731a30e1aeeae9988e2415dba2c4e91e0357955065fba.
//
// Solidity: event TestEvent(string success)
func (_Transaction *TransactionFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *TransactionTestEvent) (event.Subscription, error) {

	logs, sub, err := _Transaction.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionTestEvent)
				if err := _Transaction.contract.UnpackLog(event, "TestEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTestEvent is a log parse operation binding the contract event 0xe75028ff36bb6473da3731a30e1aeeae9988e2415dba2c4e91e0357955065fba.
//
// Solidity: event TestEvent(string success)
func (_Transaction *TransactionFilterer) ParseTestEvent(log types.Log) (*TransactionTestEvent, error) {
	event := new(TransactionTestEvent)
	if err := _Transaction.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransactionValueSenderIterator is returned from FilterValueSender and is used to iterate over the raw logs and unpacked data for ValueSender events raised by the Transaction contract.
type TransactionValueSenderIterator struct {
	Event *TransactionValueSender // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TransactionValueSenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionValueSender)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TransactionValueSender)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TransactionValueSenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionValueSenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionValueSender represents a ValueSender event raised by the Transaction contract.
type TransactionValueSender struct {
	From  common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValueSender is a free log retrieval operation binding the contract event 0x9f1dd2a0b3bec3db71afa9520af20702e06a9b0d68c36d64a1380667129cf36b.
//
// Solidity: event ValueSender(address from, uint256 money)
func (_Transaction *TransactionFilterer) FilterValueSender(opts *bind.FilterOpts) (*TransactionValueSenderIterator, error) {

	logs, sub, err := _Transaction.contract.FilterLogs(opts, "ValueSender")
	if err != nil {
		return nil, err
	}
	return &TransactionValueSenderIterator{contract: _Transaction.contract, event: "ValueSender", logs: logs, sub: sub}, nil
}

// WatchValueSender is a free log subscription operation binding the contract event 0x9f1dd2a0b3bec3db71afa9520af20702e06a9b0d68c36d64a1380667129cf36b.
//
// Solidity: event ValueSender(address from, uint256 money)
func (_Transaction *TransactionFilterer) WatchValueSender(opts *bind.WatchOpts, sink chan<- *TransactionValueSender) (event.Subscription, error) {

	logs, sub, err := _Transaction.contract.WatchLogs(opts, "ValueSender")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionValueSender)
				if err := _Transaction.contract.UnpackLog(event, "ValueSender", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValueSender is a log parse operation binding the contract event 0x9f1dd2a0b3bec3db71afa9520af20702e06a9b0d68c36d64a1380667129cf36b.
//
// Solidity: event ValueSender(address from, uint256 money)
func (_Transaction *TransactionFilterer) ParseValueSender(log types.Log) (*TransactionValueSender, error) {
	event := new(TransactionValueSender)
	if err := _Transaction.contract.UnpackLog(event, "ValueSender", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

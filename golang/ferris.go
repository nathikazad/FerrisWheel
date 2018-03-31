// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// FerrisABI is the input ABI used to generate the binding from.
const FerrisABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"chosenBidder\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AcceptedBid\",\"type\":\"event\"}]"

// FerrisBin is the compiled bytecode used for deploying new contracts.
const FerrisBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561034b8061003b6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631998aeef811461007157806338af3eed1461007b5780633ccfd60b146100aa578063c8b342ab146100d1578063cd67571c14610102575b600080fd5b610079610124565b005b341561008657600080fd5b61008e61018c565b604051600160a060020a03909116815260200160405180910390f35b34156100b557600080fd5b6100bd61019b565b604051901515815260200160405180910390f35b34156100dc57600080fd5b6100f0600160a060020a0360043516610226565b60405190815260200160405180910390f35b341561010d57600080fd5b610079600160a060020a0360043516602435610241565b33600160a060020a038116600090815260016020526040908190208054349081019091557fdd0b6c6a77960e2066c96171b4d7ac9e8b4c184011f38544afa36a5bb63ec59f929151600160a060020a03909216825260208201526040908101905180910390a1565b600054600160a060020a031681565b600160a060020a0333166000908152600160205260408120548181111561021d57600160a060020a0333166000818152600160205260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561021d57600160a060020a03331660009081526001602052604081208290559150610222565b600191505b5090565b600160a060020a031660009081526001602052604090205490565b60005433600160a060020a0390811691161461025c57600080fd5b600160a060020a0382166000908152600160205260409020548190101561028257600080fd5b600054600160a060020a031681156108fc0282604051600060405180830381858888f1935050505015156102b557600080fd5b600160a060020a03821660009081526001602052604090819020805483900390557fdefdc4699ae6600934634bec71e3e4081368537fb24e45ea245b17b8c448c391908390839051600160a060020a03909216825260208201526040908101905180910390a150505600a165627a7a723058206c8025a37d0f30134e9decc040709e220e647b30d374bfd693fdb91a6c506b700029`

// DeployFerris deploys a new Ethereum contract, binding an instance of Ferris to it.
func DeployFerris(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ferris, error) {
	parsed, err := abi.JSON(strings.NewReader(FerrisABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FerrisBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ferris{FerrisCaller: FerrisCaller{contract: contract}, FerrisTransactor: FerrisTransactor{contract: contract}, FerrisFilterer: FerrisFilterer{contract: contract}}, nil
}

// Ferris is an auto generated Go binding around an Ethereum contract.
type Ferris struct {
	FerrisCaller     // Read-only binding to the contract
	FerrisTransactor // Write-only binding to the contract
	FerrisFilterer   // Log filterer for contract events
}

// FerrisCaller is an auto generated read-only Go binding around an Ethereum contract.
type FerrisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FerrisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FerrisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FerrisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FerrisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FerrisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FerrisSession struct {
	Contract     *Ferris           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FerrisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FerrisCallerSession struct {
	Contract *FerrisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FerrisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FerrisTransactorSession struct {
	Contract     *FerrisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FerrisRaw is an auto generated low-level Go binding around an Ethereum contract.
type FerrisRaw struct {
	Contract *Ferris // Generic contract binding to access the raw methods on
}

// FerrisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FerrisCallerRaw struct {
	Contract *FerrisCaller // Generic read-only contract binding to access the raw methods on
}

// FerrisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FerrisTransactorRaw struct {
	Contract *FerrisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFerris creates a new instance of Ferris, bound to a specific deployed contract.
func NewFerris(address common.Address, backend bind.ContractBackend) (*Ferris, error) {
	contract, err := bindFerris(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ferris{FerrisCaller: FerrisCaller{contract: contract}, FerrisTransactor: FerrisTransactor{contract: contract}, FerrisFilterer: FerrisFilterer{contract: contract}}, nil
}

// NewFerrisCaller creates a new read-only instance of Ferris, bound to a specific deployed contract.
func NewFerrisCaller(address common.Address, caller bind.ContractCaller) (*FerrisCaller, error) {
	contract, err := bindFerris(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FerrisCaller{contract: contract}, nil
}

// NewFerrisTransactor creates a new write-only instance of Ferris, bound to a specific deployed contract.
func NewFerrisTransactor(address common.Address, transactor bind.ContractTransactor) (*FerrisTransactor, error) {
	contract, err := bindFerris(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FerrisTransactor{contract: contract}, nil
}

// NewFerrisFilterer creates a new log filterer instance of Ferris, bound to a specific deployed contract.
func NewFerrisFilterer(address common.Address, filterer bind.ContractFilterer) (*FerrisFilterer, error) {
	contract, err := bindFerris(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FerrisFilterer{contract: contract}, nil
}

// bindFerris binds a generic wrapper to an already deployed contract.
func bindFerris(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FerrisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ferris *FerrisRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ferris.Contract.FerrisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ferris *FerrisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ferris.Contract.FerrisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ferris *FerrisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ferris.Contract.FerrisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ferris *FerrisCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ferris.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ferris *FerrisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ferris.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ferris *FerrisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ferris.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Ferris *FerrisCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ferris.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Ferris *FerrisSession) Beneficiary() (common.Address, error) {
	return _Ferris.Contract.Beneficiary(&_Ferris.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Ferris *FerrisCallerSession) Beneficiary() (common.Address, error) {
	return _Ferris.Contract.Beneficiary(&_Ferris.CallOpts)
}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(addr address) constant returns(uint256)
func (_Ferris *FerrisCaller) GetBid(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ferris.contract.Call(opts, out, "getBid", addr)
	return *ret0, err
}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(addr address) constant returns(uint256)
func (_Ferris *FerrisSession) GetBid(addr common.Address) (*big.Int, error) {
	return _Ferris.Contract.GetBid(&_Ferris.CallOpts, addr)
}

// GetBid is a free data retrieval call binding the contract method 0xc8b342ab.
//
// Solidity: function getBid(addr address) constant returns(uint256)
func (_Ferris *FerrisCallerSession) GetBid(addr common.Address) (*big.Int, error) {
	return _Ferris.Contract.GetBid(&_Ferris.CallOpts, addr)
}

// Accept is a paid mutator transaction binding the contract method 0xcd67571c.
//
// Solidity: function accept(chosenBidder address, amount uint256) returns()
func (_Ferris *FerrisTransactor) Accept(opts *bind.TransactOpts, chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.contract.Transact(opts, "accept", chosenBidder, amount)
}

// Accept is a paid mutator transaction binding the contract method 0xcd67571c.
//
// Solidity: function accept(chosenBidder address, amount uint256) returns()
func (_Ferris *FerrisSession) Accept(chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.Contract.Accept(&_Ferris.TransactOpts, chosenBidder, amount)
}

// Accept is a paid mutator transaction binding the contract method 0xcd67571c.
//
// Solidity: function accept(chosenBidder address, amount uint256) returns()
func (_Ferris *FerrisTransactorSession) Accept(chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.Contract.Accept(&_Ferris.TransactOpts, chosenBidder, amount)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_Ferris *FerrisTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ferris.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_Ferris *FerrisSession) Bid() (*types.Transaction, error) {
	return _Ferris.Contract.Bid(&_Ferris.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_Ferris *FerrisTransactorSession) Bid() (*types.Transaction, error) {
	return _Ferris.Contract.Bid(&_Ferris.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Ferris *FerrisTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ferris.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Ferris *FerrisSession) Withdraw() (*types.Transaction, error) {
	return _Ferris.Contract.Withdraw(&_Ferris.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Ferris *FerrisTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Ferris.Contract.Withdraw(&_Ferris.TransactOpts)
}

// FerrisAcceptedBidIterator is returned from FilterAcceptedBid and is used to iterate over the raw logs and unpacked data for AcceptedBid events raised by the Ferris contract.
type FerrisAcceptedBidIterator struct {
	Event *FerrisAcceptedBid // Event containing the contract specifics and raw log

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
func (it *FerrisAcceptedBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FerrisAcceptedBid)
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
		it.Event = new(FerrisAcceptedBid)
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
func (it *FerrisAcceptedBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FerrisAcceptedBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FerrisAcceptedBid represents a AcceptedBid event raised by the Ferris contract.
type FerrisAcceptedBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAcceptedBid is a free log retrieval operation binding the contract event 0xdefdc4699ae6600934634bec71e3e4081368537fb24e45ea245b17b8c448c391.
//
// Solidity: event AcceptedBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) FilterAcceptedBid(opts *bind.FilterOpts) (*FerrisAcceptedBidIterator, error) {

	logs, sub, err := _Ferris.contract.FilterLogs(opts, "AcceptedBid")
	if err != nil {
		return nil, err
	}
	return &FerrisAcceptedBidIterator{contract: _Ferris.contract, event: "AcceptedBid", logs: logs, sub: sub}, nil
}

// WatchAcceptedBid is a free log subscription operation binding the contract event 0xdefdc4699ae6600934634bec71e3e4081368537fb24e45ea245b17b8c448c391.
//
// Solidity: event AcceptedBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) WatchAcceptedBid(opts *bind.WatchOpts, sink chan<- *FerrisAcceptedBid) (event.Subscription, error) {

	logs, sub, err := _Ferris.contract.WatchLogs(opts, "AcceptedBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FerrisAcceptedBid)
				if err := _Ferris.contract.UnpackLog(event, "AcceptedBid", log); err != nil {
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

// FerrisNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Ferris contract.
type FerrisNewBidIterator struct {
	Event *FerrisNewBid // Event containing the contract specifics and raw log

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
func (it *FerrisNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FerrisNewBid)
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
		it.Event = new(FerrisNewBid)
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
func (it *FerrisNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FerrisNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FerrisNewBid represents a NewBid event raised by the Ferris contract.
type FerrisNewBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xdd0b6c6a77960e2066c96171b4d7ac9e8b4c184011f38544afa36a5bb63ec59f.
//
// Solidity: event NewBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) FilterNewBid(opts *bind.FilterOpts) (*FerrisNewBidIterator, error) {

	logs, sub, err := _Ferris.contract.FilterLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return &FerrisNewBidIterator{contract: _Ferris.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xdd0b6c6a77960e2066c96171b4d7ac9e8b4c184011f38544afa36a5bb63ec59f.
//
// Solidity: event NewBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *FerrisNewBid) (event.Subscription, error) {

	logs, sub, err := _Ferris.contract.WatchLogs(opts, "NewBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FerrisNewBid)
				if err := _Ferris.contract.UnpackLog(event, "NewBid", log); err != nil {
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

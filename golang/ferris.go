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
const FerrisABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"bid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"chosenBidder\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AcceptedBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrewBid\",\"type\":\"event\"}]"

// FerrisBin is the compiled bytecode used for deploying new contracts.
const FerrisBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556103838061003b6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631998aeef811461007157806338af3eed1461008d5780633ccfd60b146100bc578063c8b342ab146100cf578063cd67571c14610100575b600080fd5b610079610122565b604051901515815260200160405180910390f35b341561009857600080fd5b6100a061018d565b604051600160a060020a03909116815260200160405180910390f35b34156100c757600080fd5b61007961019c565b34156100da57600080fd5b6100ee600160a060020a0360043516610258565b60405190815260200160405180910390f35b341561010b57600080fd5b610079600160a060020a0360043516602435610273565b33600160a060020a03811660009081526001602052604080822080543490810190915591927fdd0b6c6a77960e2066c96171b4d7ac9e8b4c184011f38544afa36a5bb63ec59f92909151600160a060020a03909216825260208201526040908101905180910390a190565b600054600160a060020a031681565b600160a060020a03331660009081526001602052604081205481908190116101c357600080fd5b50600160a060020a033316600081815260016020526040808220805492905590919082156108fc0290839051600060405180830381858888f19350505050151561020c57600080fd5b7f3e801e3bce46799cb1aa4cff37a34f7af26ce7f6da9024b08ec967c8e205f2e73382604051600160a060020a03909216825260208201526040908101905180910390a1600191505090565b600160a060020a031660009081526001602052604090205490565b6000805433600160a060020a0390811691161461028f57600080fd5b600160a060020a038316600090815260016020526040902054829010156102b557600080fd5b600054600160a060020a031682156108fc0283604051600060405180830381858888f1935050505015156102e857600080fd5b600160a060020a03831660009081526001602052604090819020805484900390557fdefdc4699ae6600934634bec71e3e4081368537fb24e45ea245b17b8c448c391908490849051600160a060020a03909216825260208201526040908101905180910390a1506001929150505600a165627a7a72305820b15c38ac21c1d2294b84f062890a25090a7c4a57b82dc231edcb42763c65258a0029`

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
// Solidity: function accept(chosenBidder address, amount uint256) returns(bool)
func (_Ferris *FerrisTransactor) Accept(opts *bind.TransactOpts, chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.contract.Transact(opts, "accept", chosenBidder, amount)
}

// Accept is a paid mutator transaction binding the contract method 0xcd67571c.
//
// Solidity: function accept(chosenBidder address, amount uint256) returns(bool)
func (_Ferris *FerrisSession) Accept(chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.Contract.Accept(&_Ferris.TransactOpts, chosenBidder, amount)
}

// Accept is a paid mutator transaction binding the contract method 0xcd67571c.
//
// Solidity: function accept(chosenBidder address, amount uint256) returns(bool)
func (_Ferris *FerrisTransactorSession) Accept(chosenBidder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ferris.Contract.Accept(&_Ferris.TransactOpts, chosenBidder, amount)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns(bool)
func (_Ferris *FerrisTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ferris.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns(bool)
func (_Ferris *FerrisSession) Bid() (*types.Transaction, error) {
	return _Ferris.Contract.Bid(&_Ferris.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns(bool)
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

// FerrisWithdrewBidIterator is returned from FilterWithdrewBid and is used to iterate over the raw logs and unpacked data for WithdrewBid events raised by the Ferris contract.
type FerrisWithdrewBidIterator struct {
	Event *FerrisWithdrewBid // Event containing the contract specifics and raw log

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
func (it *FerrisWithdrewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FerrisWithdrewBid)
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
		it.Event = new(FerrisWithdrewBid)
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
func (it *FerrisWithdrewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FerrisWithdrewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FerrisWithdrewBid represents a WithdrewBid event raised by the Ferris contract.
type FerrisWithdrewBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrewBid is a free log retrieval operation binding the contract event 0x3e801e3bce46799cb1aa4cff37a34f7af26ce7f6da9024b08ec967c8e205f2e7.
//
// Solidity: event WithdrewBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) FilterWithdrewBid(opts *bind.FilterOpts) (*FerrisWithdrewBidIterator, error) {

	logs, sub, err := _Ferris.contract.FilterLogs(opts, "WithdrewBid")
	if err != nil {
		return nil, err
	}
	return &FerrisWithdrewBidIterator{contract: _Ferris.contract, event: "WithdrewBid", logs: logs, sub: sub}, nil
}

// WatchWithdrewBid is a free log subscription operation binding the contract event 0x3e801e3bce46799cb1aa4cff37a34f7af26ce7f6da9024b08ec967c8e205f2e7.
//
// Solidity: event WithdrewBid(bidder address, amount uint256)
func (_Ferris *FerrisFilterer) WatchWithdrewBid(opts *bind.WatchOpts, sink chan<- *FerrisWithdrewBid) (event.Subscription, error) {

	logs, sub, err := _Ferris.contract.WatchLogs(opts, "WithdrewBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FerrisWithdrewBid)
				if err := _Ferris.contract.UnpackLog(event, "WithdrewBid", log); err != nil {
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/xcareteam/xci/accounts/abi"
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/core/types"
)

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"},{\"name\":\"DIDJson\",\"type\":\"string\"}],\"name\":\"addNewNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"inWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"getDID\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistCaller) GetDID(opts *bind.CallOpts, enode string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "getDID", enode)
	return *ret0, err
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistSession) GetDID(enode string) (string, error) {
	return _Whitelist.Contract.GetDID(&_Whitelist.CallOpts, enode)
}

// GetDID is a free data retrieval call binding the contract method 0xe0783898.
//
// Solidity: function getDID(enode string) constant returns(string)
func (_Whitelist *WhitelistCallerSession) GetDID(enode string) (string, error) {
	return _Whitelist.Contract.GetDID(&_Whitelist.CallOpts, enode)
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistCaller) InWhiteList(opts *bind.CallOpts, enode string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "inWhiteList", enode)
	return *ret0, err
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistSession) InWhiteList(enode string) (bool, error) {
	return _Whitelist.Contract.InWhiteList(&_Whitelist.CallOpts, enode)
}

// InWhiteList is a free data retrieval call binding the contract method 0xaf657f33.
//
// Solidity: function inWhiteList(enode string) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) InWhiteList(enode string) (bool, error) {
	return _Whitelist.Contract.InWhiteList(&_Whitelist.CallOpts, enode)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistTransactor) AddNewNode(opts *bind.TransactOpts, enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "addNewNode", enode, DIDJson)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistSession) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.Contract.AddNewNode(&_Whitelist.TransactOpts, enode, DIDJson)
}

// AddNewNode is a paid mutator transaction binding the contract method 0x760ca6fc.
//
// Solidity: function addNewNode(enode string, DIDJson string) returns()
func (_Whitelist *WhitelistTransactorSession) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return _Whitelist.Contract.AddNewNode(&_Whitelist.TransactOpts, enode, DIDJson)
}

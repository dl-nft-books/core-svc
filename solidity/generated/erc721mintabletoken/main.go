// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721mintabletoken

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

// Erc721mintabletokenMetaData contains all meta data concerning the Erc721mintabletoken contract.
var Erc721mintabletokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"__ERC721MintableToken_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"updateTokenParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc721mintabletokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721mintabletokenMetaData.ABI instead.
var Erc721mintabletokenABI = Erc721mintabletokenMetaData.ABI

// Erc721mintabletoken is an auto generated Go binding around an Ethereum contract.
type Erc721mintabletoken struct {
	Erc721mintabletokenCaller     // Read-only binding to the contract
	Erc721mintabletokenTransactor // Write-only binding to the contract
	Erc721mintabletokenFilterer   // Log filterer for contract events
}

// Erc721mintabletokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721mintabletokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721mintabletokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721mintabletokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721mintabletokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721mintabletokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721mintabletokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721mintabletokenSession struct {
	Contract     *Erc721mintabletoken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Erc721mintabletokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721mintabletokenCallerSession struct {
	Contract *Erc721mintabletokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Erc721mintabletokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721mintabletokenTransactorSession struct {
	Contract     *Erc721mintabletokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Erc721mintabletokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721mintabletokenRaw struct {
	Contract *Erc721mintabletoken // Generic contract binding to access the raw methods on
}

// Erc721mintabletokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721mintabletokenCallerRaw struct {
	Contract *Erc721mintabletokenCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721mintabletokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721mintabletokenTransactorRaw struct {
	Contract *Erc721mintabletokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721mintabletoken creates a new instance of Erc721mintabletoken, bound to a specific deployed contract.
func NewErc721mintabletoken(address common.Address, backend bind.ContractBackend) (*Erc721mintabletoken, error) {
	contract, err := bindErc721mintabletoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721mintabletoken{Erc721mintabletokenCaller: Erc721mintabletokenCaller{contract: contract}, Erc721mintabletokenTransactor: Erc721mintabletokenTransactor{contract: contract}, Erc721mintabletokenFilterer: Erc721mintabletokenFilterer{contract: contract}}, nil
}

// NewErc721mintabletokenCaller creates a new read-only instance of Erc721mintabletoken, bound to a specific deployed contract.
func NewErc721mintabletokenCaller(address common.Address, caller bind.ContractCaller) (*Erc721mintabletokenCaller, error) {
	contract, err := bindErc721mintabletoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721mintabletokenCaller{contract: contract}, nil
}

// NewErc721mintabletokenTransactor creates a new write-only instance of Erc721mintabletoken, bound to a specific deployed contract.
func NewErc721mintabletokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721mintabletokenTransactor, error) {
	contract, err := bindErc721mintabletoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721mintabletokenTransactor{contract: contract}, nil
}

// NewErc721mintabletokenFilterer creates a new log filterer instance of Erc721mintabletoken, bound to a specific deployed contract.
func NewErc721mintabletokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721mintabletokenFilterer, error) {
	contract, err := bindErc721mintabletoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721mintabletokenFilterer{contract: contract}, nil
}

// bindErc721mintabletoken binds a generic wrapper to an already deployed contract.
func bindErc721mintabletoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721mintabletokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721mintabletoken *Erc721mintabletokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721mintabletoken.Contract.Erc721mintabletokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721mintabletoken *Erc721mintabletokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Erc721mintabletokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721mintabletoken *Erc721mintabletokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Erc721mintabletokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721mintabletoken *Erc721mintabletokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721mintabletoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721mintabletoken *Erc721mintabletokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721mintabletoken *Erc721mintabletokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.contract.Transact(opts, method, params...)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Erc721mintabletoken *Erc721mintabletokenCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721mintabletoken.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Erc721mintabletoken *Erc721mintabletokenSession) NextTokenId() (*big.Int, error) {
	return _Erc721mintabletoken.Contract.NextTokenId(&_Erc721mintabletoken.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Erc721mintabletoken *Erc721mintabletokenCallerSession) NextTokenId() (*big.Int, error) {
	return _Erc721mintabletoken.Contract.NextTokenId(&_Erc721mintabletoken.CallOpts)
}

// ERC721MintableTokenInit is a paid mutator transaction binding the contract method 0xcdbbe0fe.
//
// Solidity: function __ERC721MintableToken_init(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactor) ERC721MintableTokenInit(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.contract.Transact(opts, "__ERC721MintableToken_init", name_, symbol_)
}

// ERC721MintableTokenInit is a paid mutator transaction binding the contract method 0xcdbbe0fe.
//
// Solidity: function __ERC721MintableToken_init(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenSession) ERC721MintableTokenInit(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.ERC721MintableTokenInit(&_Erc721mintabletoken.TransactOpts, name_, symbol_)
}

// ERC721MintableTokenInit is a paid mutator transaction binding the contract method 0xcdbbe0fe.
//
// Solidity: function __ERC721MintableToken_init(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactorSession) ERC721MintableTokenInit(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.ERC721MintableTokenInit(&_Erc721mintabletoken.TransactOpts, name_, symbol_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactor) Burn(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _Erc721mintabletoken.contract.Transact(opts, "burn", tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Burn(&_Erc721mintabletoken.TransactOpts, tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactorSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Burn(&_Erc721mintabletoken.TransactOpts, tokenId_)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to_, uint256 tokenId_, string uri_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, tokenId_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.contract.Transact(opts, "mint", to_, tokenId_, uri_)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to_, uint256 tokenId_, string uri_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenSession) Mint(to_ common.Address, tokenId_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Mint(&_Erc721mintabletoken.TransactOpts, to_, tokenId_, uri_)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to_, uint256 tokenId_, string uri_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactorSession) Mint(to_ common.Address, tokenId_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.Mint(&_Erc721mintabletoken.TransactOpts, to_, tokenId_, uri_)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x6dd84235.
//
// Solidity: function updateTokenParams(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactor) UpdateTokenParams(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.contract.Transact(opts, "updateTokenParams", name_, symbol_)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x6dd84235.
//
// Solidity: function updateTokenParams(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenSession) UpdateTokenParams(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.UpdateTokenParams(&_Erc721mintabletoken.TransactOpts, name_, symbol_)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x6dd84235.
//
// Solidity: function updateTokenParams(string name_, string symbol_) returns()
func (_Erc721mintabletoken *Erc721mintabletokenTransactorSession) UpdateTokenParams(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Erc721mintabletoken.Contract.UpdateTokenParams(&_Erc721mintabletoken.TransactOpts, name_, symbol_)
}

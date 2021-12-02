// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package app

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []common.Address
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	Participants      []common.Address
	App               common.Address
	LedgerChannel     bool
	VirtualChannel    bool
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [32]byte
	Balances []*big.Int
	IndexMap []uint16
}

// AppABI is the input ABI used to generate the binding from.
const AppABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actorIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// AppFuncSigs maps the 4-byte function signature to its string representation.
var AppFuncSigs = map[string]string{
	"0d1feb4f": "validTransition((uint256,uint256,address[],address,bool,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),uint256)",
}

// App is an auto generated Go binding around an Ethereum contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppCaller) ValidTransition(opts *bind.CallOpts, params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "validTransition", params, from, to, actorIdx)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _App.Contract.ValidTransition(&_App.CallOpts, params, from, to, actorIdx)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) params, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) from, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_App *AppCallerSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _App.Contract.ValidTransition(&_App.CallOpts, params, from, to, actorIdx)
}

// ArrayABI is the input ABI used to generate the binding from.
const ArrayABI = "[]"

// ArrayBin is the compiled bytecode used for deploying new contracts.
var ArrayBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209f4d25c7e3cd971c189df97503706b5ca3a995a0082abe18c64da43c672f747264736f6c63430007060033"

// DeployArray deploys a new Ethereum contract, binding an instance of Array to it.
func DeployArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Array, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// Array is an auto generated Go binding around an Ethereum contract.
type Array struct {
	ArrayCaller     // Read-only binding to the contract
	ArrayTransactor // Write-only binding to the contract
	ArrayFilterer   // Log filterer for contract events
}

// ArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraySession struct {
	Contract     *Array            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayCallerSession struct {
	Contract *ArrayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayTransactorSession struct {
	Contract     *ArrayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayRaw struct {
	Contract *Array // Generic contract binding to access the raw methods on
}

// ArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayCallerRaw struct {
	Contract *ArrayCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayTransactorRaw struct {
	Contract *ArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArray creates a new instance of Array, bound to a specific deployed contract.
func NewArray(address common.Address, backend bind.ContractBackend) (*Array, error) {
	contract, err := bindArray(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// NewArrayCaller creates a new read-only instance of Array, bound to a specific deployed contract.
func NewArrayCaller(address common.Address, caller bind.ContractCaller) (*ArrayCaller, error) {
	contract, err := bindArray(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayCaller{contract: contract}, nil
}

// NewArrayTransactor creates a new write-only instance of Array, bound to a specific deployed contract.
func NewArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayTransactor, error) {
	contract, err := bindArray(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayTransactor{contract: contract}, nil
}

// NewArrayFilterer creates a new log filterer instance of Array, bound to a specific deployed contract.
func NewArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayFilterer, error) {
	contract, err := bindArray(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayFilterer{contract: contract}, nil
}

// bindArray binds a generic wrapper to an already deployed contract.
func bindArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Array.Contract.ArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Array.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.contract.Transact(opts, method, params...)
}

// ChannelABI is the input ABI used to generate the binding from.
const ChannelABI = "[]"

// ChannelBin is the compiled bytecode used for deploying new contracts.
var ChannelBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ebe3379116f0685b95af4da576c211ec5b18abdb3e8fe45d14a084f0dac865ef64736f6c63430007060033"

// DeployChannel deploys a new Ethereum contract, binding an instance of Channel to it.
func DeployChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Channel, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// Channel is an auto generated Go binding around an Ethereum contract.
type Channel struct {
	ChannelCaller     // Read-only binding to the contract
	ChannelTransactor // Write-only binding to the contract
	ChannelFilterer   // Log filterer for contract events
}

// ChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelSession struct {
	Contract     *Channel          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelCallerSession struct {
	Contract *ChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelTransactorSession struct {
	Contract     *ChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelRaw struct {
	Contract *Channel // Generic contract binding to access the raw methods on
}

// ChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelCallerRaw struct {
	Contract *ChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelTransactorRaw struct {
	Contract *ChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannel creates a new instance of Channel, bound to a specific deployed contract.
func NewChannel(address common.Address, backend bind.ContractBackend) (*Channel, error) {
	contract, err := bindChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// NewChannelCaller creates a new read-only instance of Channel, bound to a specific deployed contract.
func NewChannelCaller(address common.Address, caller bind.ContractCaller) (*ChannelCaller, error) {
	contract, err := bindChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelCaller{contract: contract}, nil
}

// NewChannelTransactor creates a new write-only instance of Channel, bound to a specific deployed contract.
func NewChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelTransactor, error) {
	contract, err := bindChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelTransactor{contract: contract}, nil
}

// NewChannelFilterer creates a new log filterer instance of Channel, bound to a specific deployed contract.
func NewChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFilterer, error) {
	contract, err := bindChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFilterer{contract: contract}, nil
}

// bindChannel binds a generic wrapper to an already deployed contract.
func bindChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.ChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transact(opts, method, params...)
}

// CredentialSwapABI is the input ABI used to generate the binding from.
const CredentialSwapABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"cur\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actor\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// CredentialSwapFuncSigs maps the 4-byte function signature to its string representation.
var CredentialSwapFuncSigs = map[string]string{
	"0d1feb4f": "validTransition((uint256,uint256,address[],address,bool,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),(bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),uint256)",
}

// CredentialSwapBin is the compiled bytecode used for deploying new contracts.
var CredentialSwapBin = "0x608060405234801561001057600080fd5b5061108b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80630d1feb4f14610030575b600080fd5b61004361003e366004610c07565b610045565b005b61004f8383610155565b600061005a84610292565b805190915060ff166001141561008d5760006100798260200151610327565b90506100878186868661034c565b5061014e565b6100978484610516565b60006100a284610292565b805190915060ff166001141561014c5760006100c18260200151610327565b90503660006100d36040880188610f08565b6100e1906020810190610e75565b6040850151919350915082826000816100f657fe5b90506020028101906101089190610e75565b856060015161ffff1681811061011a57fe5b9050602002013510156101485760405162461bcd60e51b815260040161013f90610e49565b60405180910390fd5b5050505b505b5050505050565b6000806101656040850185610f08565b61016f9080610e75565b61017c6040860186610f08565b6101869080610e75565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160208087028281018201909352868252949750959493849350860191508490808284376000920191909152505082519294506001928314915061021090505760405162461bcd60e51b815260040161013f90610d90565b8060ff168251146102335760405162461bcd60e51b815260040161013f90610d59565b81600060ff168151811061024357fe5b60200260200101516001600160a01b031683600060ff168151811061026457fe5b60200260200101516001600160a01b03161461014e5760405162461bcd60e51b815260040161013f90610cc5565b61029a610a27565b60026000816102ac6060860186610ec3565b905003905060006103058580606001906102c69190610ec3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505060ff861690508461056a565b905060008180602001905181019061031d9190610a90565b9695505050505050565b61032f610a3f565b6000828060200190518101906103459190610b88565b9392505050565b60008061035884610673565b91509150806103795760405162461bcd60e51b815260040161013f90610d2e565b602086015182518751859261038f9290916106d2565b6103ab5760405162461bcd60e51b815260040161013f90610c9a565b3660006103bb6040890189610f08565b6103c9906020810190610e75565b90925090503660006103de60408a018a610f08565b6103ec906020810190610e75565b60408d01519193509150848460008161040157fe5b90506020028101906104139190610e75565b8d6060015161ffff1681811061042557fe5b90506020020135038282600060ff1681811061043d57fe5b905060200281019061044f9190610e75565b8d6060015161ffff1681811061046157fe5b90506020020135146104855760405162461bcd60e51b815260040161013f90610dd1565b60408b0151848460008161049557fe5b90506020028101906104a79190610e75565b878181106104b157fe5b90506020020135018282600060ff168181106104c957fe5b90506020028101906104db9190610e75565b878181106104e557fe5b90506020020135146105095760405162461bcd60e51b815260040161013f90610cec565b5050505050505050505050565b6105666105266040840184610f08565b610534906020810190610e75565b61053d91610f69565b61054a6040840184610f08565b610558906020810190610e75565b61056191610f69565b6106f8565b5050565b60608182601f0110156105b5576040805162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015290519081900360640190fd5b81830184511015610601576040805162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015290519081900360640190fd5b606082158015610620576040519150600082526020820160405261066a565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610659578051835260209283019201610641565b5050858452601f01601f1916604052505b50949350505050565b61067b610a66565b60008061068784610292565b805190915060ff166002146106b55750506040805180820190915260006020820181815282529091506106cd565b60408051602080820190925291015181529150600190505b915091565b6000806106df8585610762565b6001600160a01b03908116908416149150509392505050565b80518251146107195760405162461bcd60e51b815260040161013f90610e12565b60005b825181101561075d5761075583828151811061073457fe5b602002602001015183838151811061074857fe5b6020026020010151610943565b60010161071c565b505050565b600081516041146107ba576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561082b5760405162461bcd60e51b81526004018080602001828103825260228152602001806110126022913960400191505060405180910390fd5b8060ff16601b1415801561084357508060ff16601c14155b1561087f5760405162461bcd60e51b81526004018080602001828103825260228152602001806110346022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156108db573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661031d576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b8051825114610999576040805162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e67746800000000000000604482015290519081900360640190fd5b60005b825181101561075d578181815181106109b157fe5b60200260200101518382815181106109c557fe5b602002602001015114610a1f576040805162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d000000000000000000604482015290519081900360640190fd5b60010161099c565b60408051808201909152600081526060602082015290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6040518060200160405280606081525090565b600060a08284031215610a8a578081fd5b50919050565b60006020808385031215610aa2578182fd5b825167ffffffffffffffff80821115610ab9578384fd5b9084019060408287031215610acc578384fd5b604051604081018181108382111715610ae157fe5b604052825160ff81168114610af4578586fd5b81528284015182811115610b06578586fd5b80840193505086601f840112610b1a578485fd5b825182811115610b2657fe5b610b38601f8201601f19168601610f27565b92508083528785828601011115610b4d578586fd5b855b81811015610b6a578481018601518482018701528501610b4f565b81811115610b7a57868683860101525b505092830152509392505050565b600060808284031215610b99578081fd5b6040516080810181811067ffffffffffffffff82111715610bb657fe5b60405282516001600160a01b0381168114610bcf578283fd5b808252506020830151602082015260408301516040820152606083015161ffff81168114610bfb578283fd5b60608201529392505050565b60008060008060808587031215610c1c578283fd5b843567ffffffffffffffff80821115610c33578485fd5b9086019060c08289031215610c46578485fd5b90945060208601359080821115610c5b578485fd5b610c6788838901610a79565b94506040870135915080821115610c7c578384fd5b50610c8987828801610a79565b949793965093946060013593505050565b602080825260119082015270696e76616c6964207369676e617475726560781b604082015260600190565b6020808252600d908201526c1a5b9d985b1a5908185cdcd95d609a1b604082015260600190565b60208082526022908201527f696e76616c696420616d6f756e74207472616e736665727265643a2073656c6c60408201526132b960f11b606082015260800190565b602080825260119082015270696e76616c6964206e657874206d6f646560781b604082015260600190565b6020808252601e908201527f696e76616c6964206e756d626572206f66206173736574733a206e6578740000604082015260600190565b60208082526021908201527f696e76616c6964206e756d626572206f66206173736574733a2063757272656e6040820152601d60fa1b606082015260800190565b60208082526021908201527f696e76616c696420616d6f756e74207472616e736665727265643a20627579656040820152603960f91b606082015260800190565b6020808252601b908201527f75696e743235365b5d5b5d3a20756e657175616c206c656e6774680000000000604082015260600190565b602080825260129082015271696e73756666696369656e742066756e647360701b604082015260600190565b6000808335601e19843603018112610e8b578283fd5b83018035915067ffffffffffffffff821115610ea5578283fd5b6020908101925081023603821315610ebc57600080fd5b9250929050565b6000808335601e19843603018112610ed9578283fd5b83018035915067ffffffffffffffff821115610ef3578283fd5b602001915036819003821315610ebc57600080fd5b60008235605e19833603018112610f1d578182fd5b9190910192915050565b60405181810167ffffffffffffffff81118282101715610f4357fe5b604052919050565b600067ffffffffffffffff821115610f5f57fe5b5060209081020190565b6000610f7c610f7784610f4b565b610f27565b8381526020808201919084845b87811015611005578135870136601f820112610fa3578687fd5b8035610fb1610f7782610f4b565b8181528581019083870136888502860189011115610fcd578a8bfd5b8a94505b83851015610fef578035835260019490940193918701918701610fd1565b5088525050509382019390820190600101610f89565b5091969550505050505056fe45434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c7565a2646970667358221220cfcc15bea829ce7b872ad4fbfbdce0336d3e14594ec9223d90e2040d3fcfa9e064736f6c63430007060033"

// DeployCredentialSwap deploys a new Ethereum contract, binding an instance of CredentialSwap to it.
func DeployCredentialSwap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CredentialSwap, error) {
	parsed, err := abi.JSON(strings.NewReader(CredentialSwapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CredentialSwapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CredentialSwap{CredentialSwapCaller: CredentialSwapCaller{contract: contract}, CredentialSwapTransactor: CredentialSwapTransactor{contract: contract}, CredentialSwapFilterer: CredentialSwapFilterer{contract: contract}}, nil
}

// CredentialSwap is an auto generated Go binding around an Ethereum contract.
type CredentialSwap struct {
	CredentialSwapCaller     // Read-only binding to the contract
	CredentialSwapTransactor // Write-only binding to the contract
	CredentialSwapFilterer   // Log filterer for contract events
}

// CredentialSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type CredentialSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CredentialSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CredentialSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CredentialSwapSession struct {
	Contract     *CredentialSwap   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CredentialSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CredentialSwapCallerSession struct {
	Contract *CredentialSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CredentialSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CredentialSwapTransactorSession struct {
	Contract     *CredentialSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CredentialSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type CredentialSwapRaw struct {
	Contract *CredentialSwap // Generic contract binding to access the raw methods on
}

// CredentialSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CredentialSwapCallerRaw struct {
	Contract *CredentialSwapCaller // Generic read-only contract binding to access the raw methods on
}

// CredentialSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CredentialSwapTransactorRaw struct {
	Contract *CredentialSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredentialSwap creates a new instance of CredentialSwap, bound to a specific deployed contract.
func NewCredentialSwap(address common.Address, backend bind.ContractBackend) (*CredentialSwap, error) {
	contract, err := bindCredentialSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CredentialSwap{CredentialSwapCaller: CredentialSwapCaller{contract: contract}, CredentialSwapTransactor: CredentialSwapTransactor{contract: contract}, CredentialSwapFilterer: CredentialSwapFilterer{contract: contract}}, nil
}

// NewCredentialSwapCaller creates a new read-only instance of CredentialSwap, bound to a specific deployed contract.
func NewCredentialSwapCaller(address common.Address, caller bind.ContractCaller) (*CredentialSwapCaller, error) {
	contract, err := bindCredentialSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialSwapCaller{contract: contract}, nil
}

// NewCredentialSwapTransactor creates a new write-only instance of CredentialSwap, bound to a specific deployed contract.
func NewCredentialSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*CredentialSwapTransactor, error) {
	contract, err := bindCredentialSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialSwapTransactor{contract: contract}, nil
}

// NewCredentialSwapFilterer creates a new log filterer instance of CredentialSwap, bound to a specific deployed contract.
func NewCredentialSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*CredentialSwapFilterer, error) {
	contract, err := bindCredentialSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CredentialSwapFilterer{contract: contract}, nil
}

// bindCredentialSwap binds a generic wrapper to an already deployed contract.
func bindCredentialSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CredentialSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredentialSwap *CredentialSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredentialSwap.Contract.CredentialSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredentialSwap *CredentialSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredentialSwap.Contract.CredentialSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredentialSwap *CredentialSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredentialSwap.Contract.CredentialSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredentialSwap *CredentialSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredentialSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredentialSwap *CredentialSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredentialSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredentialSwap *CredentialSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredentialSwap.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) , (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) cur, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) next, uint256 actor) pure returns()
func (_CredentialSwap *CredentialSwapCaller) ValidTransition(opts *bind.CallOpts, arg0 ChannelParams, cur ChannelState, next ChannelState, actor *big.Int) error {
	var out []interface{}
	err := _CredentialSwap.contract.Call(opts, &out, "validTransition", arg0, cur, next, actor)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) , (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) cur, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) next, uint256 actor) pure returns()
func (_CredentialSwap *CredentialSwapSession) ValidTransition(arg0 ChannelParams, cur ChannelState, next ChannelState, actor *big.Int) error {
	return _CredentialSwap.Contract.ValidTransition(&_CredentialSwap.CallOpts, arg0, cur, next, actor)
}

// ValidTransition is a free data retrieval call binding the contract method 0x0d1feb4f.
//
// Solidity: function validTransition((uint256,uint256,address[],address,bool,bool) , (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) cur, (bytes32,uint64,(address[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) next, uint256 actor) pure returns()
func (_CredentialSwap *CredentialSwapCallerSession) ValidTransition(arg0 ChannelParams, cur ChannelState, next ChannelState, actor *big.Int) error {
	return _CredentialSwap.Contract.ValidTransition(&_CredentialSwap.CallOpts, arg0, cur, next, actor)
}

// DecodeABI is the input ABI used to generate the binding from.
const DecodeABI = "[]"

// DecodeBin is the compiled bytecode used for deploying new contracts.
var DecodeBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084ac21a72f6174ad3b6bf4941f013451d27571b7d7fe44d8534ec9bb7ea7e38264736f6c63430007060033"

// DeployDecode deploys a new Ethereum contract, binding an instance of Decode to it.
func DeployDecode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Decode, error) {
	parsed, err := abi.JSON(strings.NewReader(DecodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DecodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Decode{DecodeCaller: DecodeCaller{contract: contract}, DecodeTransactor: DecodeTransactor{contract: contract}, DecodeFilterer: DecodeFilterer{contract: contract}}, nil
}

// Decode is an auto generated Go binding around an Ethereum contract.
type Decode struct {
	DecodeCaller     // Read-only binding to the contract
	DecodeTransactor // Write-only binding to the contract
	DecodeFilterer   // Log filterer for contract events
}

// DecodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecodeSession struct {
	Contract     *Decode           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecodeCallerSession struct {
	Contract *DecodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DecodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecodeTransactorSession struct {
	Contract     *DecodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecodeRaw struct {
	Contract *Decode // Generic contract binding to access the raw methods on
}

// DecodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecodeCallerRaw struct {
	Contract *DecodeCaller // Generic read-only contract binding to access the raw methods on
}

// DecodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecodeTransactorRaw struct {
	Contract *DecodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecode creates a new instance of Decode, bound to a specific deployed contract.
func NewDecode(address common.Address, backend bind.ContractBackend) (*Decode, error) {
	contract, err := bindDecode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Decode{DecodeCaller: DecodeCaller{contract: contract}, DecodeTransactor: DecodeTransactor{contract: contract}, DecodeFilterer: DecodeFilterer{contract: contract}}, nil
}

// NewDecodeCaller creates a new read-only instance of Decode, bound to a specific deployed contract.
func NewDecodeCaller(address common.Address, caller bind.ContractCaller) (*DecodeCaller, error) {
	contract, err := bindDecode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecodeCaller{contract: contract}, nil
}

// NewDecodeTransactor creates a new write-only instance of Decode, bound to a specific deployed contract.
func NewDecodeTransactor(address common.Address, transactor bind.ContractTransactor) (*DecodeTransactor, error) {
	contract, err := bindDecode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecodeTransactor{contract: contract}, nil
}

// NewDecodeFilterer creates a new log filterer instance of Decode, bound to a specific deployed contract.
func NewDecodeFilterer(address common.Address, filterer bind.ContractFilterer) (*DecodeFilterer, error) {
	contract, err := bindDecode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecodeFilterer{contract: contract}, nil
}

// bindDecode binds a generic wrapper to an already deployed contract.
func bindDecode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DecodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Decode *DecodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Decode.Contract.DecodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Decode *DecodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Decode.Contract.DecodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Decode *DecodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Decode.Contract.DecodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Decode *DecodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Decode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Decode *DecodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Decode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Decode *DecodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Decode.Contract.contract.Transact(opts, method, params...)
}

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206864e8222b45afce4d7d7178eae6e7e5fc2ecf2fb7014b473611c107655370b364736f6c63430007060033"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220711fae5256c1fee8fbc215255c19ada9220a069445547ec5b42ec69d74584b4164736f6c63430007060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SigABI is the input ABI used to generate the binding from.
const SigABI = "[]"

// SigBin is the compiled bytecode used for deploying new contracts.
var SigBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201718b7e8ea73e0a79b0b33b8b17bcb8b4cb4240ba977ae47a506fe8a5075556e64736f6c63430007060033"

// DeploySig deploys a new Ethereum contract, binding an instance of Sig to it.
func DeploySig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sig, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// Sig is an auto generated Go binding around an Ethereum contract.
type Sig struct {
	SigCaller     // Read-only binding to the contract
	SigTransactor // Write-only binding to the contract
	SigFilterer   // Log filterer for contract events
}

// SigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigSession struct {
	Contract     *Sig              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigCallerSession struct {
	Contract *SigCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigTransactorSession struct {
	Contract     *SigTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigRaw struct {
	Contract *Sig // Generic contract binding to access the raw methods on
}

// SigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigCallerRaw struct {
	Contract *SigCaller // Generic read-only contract binding to access the raw methods on
}

// SigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigTransactorRaw struct {
	Contract *SigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSig creates a new instance of Sig, bound to a specific deployed contract.
func NewSig(address common.Address, backend bind.ContractBackend) (*Sig, error) {
	contract, err := bindSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// NewSigCaller creates a new read-only instance of Sig, bound to a specific deployed contract.
func NewSigCaller(address common.Address, caller bind.ContractCaller) (*SigCaller, error) {
	contract, err := bindSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigCaller{contract: contract}, nil
}

// NewSigTransactor creates a new write-only instance of Sig, bound to a specific deployed contract.
func NewSigTransactor(address common.Address, transactor bind.ContractTransactor) (*SigTransactor, error) {
	contract, err := bindSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigTransactor{contract: contract}, nil
}

// NewSigFilterer creates a new log filterer instance of Sig, bound to a specific deployed contract.
func NewSigFilterer(address common.Address, filterer bind.ContractFilterer) (*SigFilterer, error) {
	contract, err := bindSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigFilterer{contract: contract}, nil
}

// bindSig binds a generic wrapper to an already deployed contract.
func bindSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.SigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transact(opts, method, params...)
}

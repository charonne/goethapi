// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CoinFidTokenABI is the input ABI used to generate the binding from.
const CoinFidTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDistributedCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperAdmin\",\"type\":\"address\"}],\"name\":\"nameSuperAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDistributor\",\"type\":\"address\"}],\"name\":\"nameDistributor\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_superAdmin\",\"type\":\"address\"}],\"name\":\"removeSuperAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"generate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"removeDistributor\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"superAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewardsOffered\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTokenSymbol\",\"type\":\"string\"}],\"name\":\"setTokenSymbol\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTokenName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"distributors\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"nameAdmin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"changeCoinsForReward\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"firstSuperAdmin\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// CoinFidTokenBin is the compiled bytecode used for deploying new contracts.
const CoinFidTokenBin = `0x606060405260408051908101604052600981527f546f6b656e20302e310000000000000000000000000000000000000000000000602082015260009080516200004d92916020019062000163565b5034156200005a57600080fd5b604051620010e2380380620010e28339810160405280805191906020018051820191906020018051919060200180518201919060200180519150505b60048590556001848051620000b092916020019062000163565b506002828051620000c692916020019062000163565b506003805460ff851660ff1991821617909155600060058190556006819055600160a060020a0380841682526007602090815260408084208054861660019081179091556008808452828620805488168317905560098085528387208054891684179055339095168652835281852080548716821790559282528084208054909516909217909355600a909252208590555b50505050506200020d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001a657805160ff1916838001178555620001d6565b82800160010185558215620001d6579182015b82811115620001d6578251825591602001919060010190620001b9565b5b50620001e5929150620001e9565b5090565b6200020a91905b80821115620001e55760008155600101620001f0565b5090565b90565b610ec5806200021d6000396000f300606060405236156101305763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146101355780631785f53c146101c057806318160ddd146101e1578063195527ae14610206578063260e84571461022b57806326ed13fc1461024c578063313ce5671461026d578063429b62e5146102965780634902e4aa146102c95780634a7dd523146102ea57806357c1f9e2146103025780635a3b7e421461032357806370a08231146103ae57806395d89b41146103df578063a9059cbb1461046a578063ad25cb421461048e578063b6bbd864146104c1578063ba51b1b4146104e6578063c47f002714610539578063cc6427841461058c578063ebe23006146105bf578063f9328975146105e0578063fb93210814610604575b600080fd5b341561014057600080fd5b610148610628565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101855780820151818401525b60200161016c565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101cb57600080fd5b6101df600160a060020a03600435166106c6565b005b34156101ec57600080fd5b6101f461074a565b60405190815260200160405180910390f35b341561021157600080fd5b6101f4610750565b60405190815260200160405180910390f35b341561023657600080fd5b6101df600160a060020a0360043516610756565b005b341561025757600080fd5b6101df600160a060020a03600435166107ce565b005b341561027857600080fd5b61028061081f565b60405160ff909116815260200160405180910390f35b34156102a157600080fd5b6102b5600160a060020a0360043516610828565b604051901515815260200160405180910390f35b34156102d457600080fd5b6101df600160a060020a036004351661083d565b005b34156102f557600080fd5b6101df6004356108ac565b005b341561030d57600080fd5b6101df600160a060020a0360043516610900565b005b341561032e57600080fd5b610148610991565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101855780820151818401525b60200161016c565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103b957600080fd5b6101f4600160a060020a0360043516610a2f565b60405190815260200160405180910390f35b34156103ea57600080fd5b610148610a4e565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101855780820151818401525b60200161016c565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561047557600080fd5b6101df600160a060020a0360043516602435610aec565b005b341561049957600080fd5b6102b5600160a060020a0360043516610bb8565b604051901515815260200160405180910390f35b34156104cc57600080fd5b6101f4610bcd565b60405190815260200160405180910390f35b34156104f157600080fd5b6101df60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610bd395505050505050565b005b341561054457600080fd5b6101df60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610beb95505050505050565b005b341561059757600080fd5b6102b5600160a060020a0360043516610c03565b604051901515815260200160405180910390f35b34156105ca57600080fd5b6101df600160a060020a0360043516610c18565b005b34156105eb57600080fd5b6101df600160a060020a0360043516602435610c81565b005b341561060f57600080fd5b6101df600160a060020a0360043516602435610d5c565b005b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106be5780601f10610693576101008083540402835291602001916106be565b820191906000526020600020905b8154815290600101906020018083116106a157829003601f168201915b505050505081565b600160a060020a03331660009081526008602052604090205460ff1615156001146106f057600080fd5b80600160a060020a031633600160a060020a03161415151561071157600080fd5b600160a060020a0381166000908152600860209081526040808320805460ff199081169091556007909252909120805490911690555b50565b60045481565b60055481565b600160a060020a03331660009081526007602052604090205460ff16151560011461078057600080fd5b600160a060020a03811660009081526007602090815260408083208054600160ff19918216811790925560088452828520805482168317905560099093529220805490911690911790555b50565b600160a060020a03331660009081526008602052604090205460ff1615156001146107f857600080fd5b600160a060020a0381166000908152600960205260409020805460ff191660011790555b50565b60035460ff1681565b60086020526000908152604090205460ff1681565b600160a060020a03331660009081526007602052604090205460ff16151560011461086757600080fd5b80600160a060020a031633600160a060020a03161415151561088857600080fd5b600160a060020a0381166000908152600760205260409020805460ff191690555b50565b600160a060020a03331660009081526007602052604090205460ff1615156001146108d657600080fd5b600160a060020a0333166000908152600a6020526040902080548201905560048054820190555b50565b600160a060020a03331660009081526008602052604090205460ff16151560011461092a57600080fd5b80600160a060020a031633600160a060020a03161415151561094b57600080fd5b600160a060020a0381166000908152600960209081526040808320805460ff19908116909155600883528184208054821690556007909252909120805490911690555b50565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106be5780601f10610693576101008083540402835291602001916106be565b820191906000526020600020905b8154815290600101906020018083116106a157829003601f168201915b505050505081565b600160a060020a0381166000908152600a60205260409020545b919050565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106be5780601f10610693576101008083540402835291602001916106be565b820191906000526020600020905b8154815290600101906020018083116106a157829003601f168201915b505050505081565b600160a060020a0382161515610b0157600080fd5b600160a060020a0333166000908152600a602052604090205481901015610b2757600080fd5b600160a060020a0382166000908152600a60205260409020548181011015610b4e57600080fd5b600160a060020a033381166000818152600a60205260408082208054869003905592851680825290839020805485019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905190815260200160405180910390a35b5050565b60076020526000908152604090205460ff1681565b60065481565b6002818051610bb4929160200190610df9565b505b50565b6001818051610bb4929160200190610df9565b505b50565b60096020526000908152604090205460ff1681565b600160a060020a03331660009081526008602052604090205460ff161515600114610c4257600080fd5b600160a060020a03811660009081526008602090815260408083208054600160ff19918216811790925560099093529220805490911690911790555b50565b600160a060020a03331660009081526009602052604090205460ff161515600114610cab57600080fd5b600160a060020a0382166000908152600a602052604090205481901015610cd157600080fd5b600160a060020a0382166000908152600a6020908152604080832080548590039055600480548590039055600990915290205460ff161515600114610d1a576006805460010190555b81600160a060020a03167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405190815260200160405180910390a25b5050565b600160a060020a0382161515610d7157600080fd5b600160a060020a03331660009081526009602052604090205460ff161515600114610d9b57600080fd5b33600160a060020a031682600160a060020a031614151515610dbc57600080fd5b600160a060020a0333166000908152600a6020526040902080548201905560048054820190556005805482019055610bb48282610aec565b5b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e3a57805160ff1916838001178555610e67565b82800160010185558215610e67579182015b82811115610e67578251825591602001919060010190610e4c565b5b50610e74929150610e78565b5090565b610e9691905b80821115610e745760008155600101610e7e565b5090565b905600a165627a7a72305820198f87c12385e6eb8102aadbe8cd2eced68520b64050f5decf556eb6c25d818b0029`

// DeployCoinFidToken deploys a new Ethereum contract, binding an instance of CoinFidToken to it.
func DeployCoinFidToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string, firstSuperAdmin common.Address) (common.Address, *types.Transaction, *CoinFidToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinFidTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoinFidTokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol, firstSuperAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoinFidToken{CoinFidTokenCaller: CoinFidTokenCaller{contract: contract}, CoinFidTokenTransactor: CoinFidTokenTransactor{contract: contract}}, nil
}

// CoinFidToken is an auto generated Go binding around an Ethereum contract.
type CoinFidToken struct {
	CoinFidTokenCaller     // Read-only binding to the contract
	CoinFidTokenTransactor // Write-only binding to the contract
}

// CoinFidTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinFidTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFidTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinFidTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFidTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinFidTokenSession struct {
	Contract     *CoinFidToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinFidTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinFidTokenCallerSession struct {
	Contract *CoinFidTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CoinFidTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinFidTokenTransactorSession struct {
	Contract     *CoinFidTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CoinFidTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinFidTokenRaw struct {
	Contract *CoinFidToken // Generic contract binding to access the raw methods on
}

// CoinFidTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinFidTokenCallerRaw struct {
	Contract *CoinFidTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CoinFidTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinFidTokenTransactorRaw struct {
	Contract *CoinFidTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinFidToken creates a new instance of CoinFidToken, bound to a specific deployed contract.
func NewCoinFidToken(address common.Address, backend bind.ContractBackend) (*CoinFidToken, error) {
	contract, err := bindCoinFidToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoinFidToken{CoinFidTokenCaller: CoinFidTokenCaller{contract: contract}, CoinFidTokenTransactor: CoinFidTokenTransactor{contract: contract}}, nil
}

// NewCoinFidTokenCaller creates a new read-only instance of CoinFidToken, bound to a specific deployed contract.
func NewCoinFidTokenCaller(address common.Address, caller bind.ContractCaller) (*CoinFidTokenCaller, error) {
	contract, err := bindCoinFidToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CoinFidTokenCaller{contract: contract}, nil
}

// NewCoinFidTokenTransactor creates a new write-only instance of CoinFidToken, bound to a specific deployed contract.
func NewCoinFidTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinFidTokenTransactor, error) {
	contract, err := bindCoinFidToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CoinFidTokenTransactor{contract: contract}, nil
}

// bindCoinFidToken binds a generic wrapper to an already deployed contract.
func bindCoinFidToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinFidTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinFidToken *CoinFidTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinFidToken.Contract.CoinFidTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinFidToken *CoinFidTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFidToken.Contract.CoinFidTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinFidToken *CoinFidTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinFidToken.Contract.CoinFidTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinFidToken *CoinFidTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinFidToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinFidToken *CoinFidTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFidToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinFidToken *CoinFidTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinFidToken.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "admins", arg0)
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenSession) Admins(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.Admins(&_CoinFidToken.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.Admins(&_CoinFidToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCaller) BalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "balanceOf", _address)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_CoinFidToken *CoinFidTokenSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _CoinFidToken.Contract.BalanceOf(&_CoinFidToken.CallOpts, _address)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCallerSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _CoinFidToken.Contract.BalanceOf(&_CoinFidToken.CallOpts, _address)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CoinFidToken *CoinFidTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CoinFidToken *CoinFidTokenSession) Decimals() (uint8, error) {
	return _CoinFidToken.Contract.Decimals(&_CoinFidToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_CoinFidToken *CoinFidTokenCallerSession) Decimals() (uint8, error) {
	return _CoinFidToken.Contract.Decimals(&_CoinFidToken.CallOpts)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCaller) Distributors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "distributors", arg0)
	return *ret0, err
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenSession) Distributors(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.Distributors(&_CoinFidToken.CallOpts, arg0)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCallerSession) Distributors(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.Distributors(&_CoinFidToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CoinFidToken *CoinFidTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CoinFidToken *CoinFidTokenSession) Name() (string, error) {
	return _CoinFidToken.Contract.Name(&_CoinFidToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CoinFidToken *CoinFidTokenCallerSession) Name() (string, error) {
	return _CoinFidToken.Contract.Name(&_CoinFidToken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_CoinFidToken *CoinFidTokenCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_CoinFidToken *CoinFidTokenSession) Standard() (string, error) {
	return _CoinFidToken.Contract.Standard(&_CoinFidToken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_CoinFidToken *CoinFidTokenCallerSession) Standard() (string, error) {
	return _CoinFidToken.Contract.Standard(&_CoinFidToken.CallOpts)
}

// SuperAdmins is a free data retrieval call binding the contract method 0xad25cb42.
//
// Solidity: function superAdmins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCaller) SuperAdmins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "superAdmins", arg0)
	return *ret0, err
}

// SuperAdmins is a free data retrieval call binding the contract method 0xad25cb42.
//
// Solidity: function superAdmins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenSession) SuperAdmins(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.SuperAdmins(&_CoinFidToken.CallOpts, arg0)
}

// SuperAdmins is a free data retrieval call binding the contract method 0xad25cb42.
//
// Solidity: function superAdmins( address) constant returns(bool)
func (_CoinFidToken *CoinFidTokenCallerSession) SuperAdmins(arg0 common.Address) (bool, error) {
	return _CoinFidToken.Contract.SuperAdmins(&_CoinFidToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CoinFidToken *CoinFidTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CoinFidToken *CoinFidTokenSession) Symbol() (string, error) {
	return _CoinFidToken.Contract.Symbol(&_CoinFidToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CoinFidToken *CoinFidTokenCallerSession) Symbol() (string, error) {
	return _CoinFidToken.Contract.Symbol(&_CoinFidToken.CallOpts)
}

// TotalDistributedCoins is a free data retrieval call binding the contract method 0x195527ae.
//
// Solidity: function totalDistributedCoins() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCaller) TotalDistributedCoins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "totalDistributedCoins")
	return *ret0, err
}

// TotalDistributedCoins is a free data retrieval call binding the contract method 0x195527ae.
//
// Solidity: function totalDistributedCoins() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenSession) TotalDistributedCoins() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalDistributedCoins(&_CoinFidToken.CallOpts)
}

// TotalDistributedCoins is a free data retrieval call binding the contract method 0x195527ae.
//
// Solidity: function totalDistributedCoins() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCallerSession) TotalDistributedCoins() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalDistributedCoins(&_CoinFidToken.CallOpts)
}

// TotalRewardsOffered is a free data retrieval call binding the contract method 0xb6bbd864.
//
// Solidity: function totalRewardsOffered() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCaller) TotalRewardsOffered(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "totalRewardsOffered")
	return *ret0, err
}

// TotalRewardsOffered is a free data retrieval call binding the contract method 0xb6bbd864.
//
// Solidity: function totalRewardsOffered() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenSession) TotalRewardsOffered() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalRewardsOffered(&_CoinFidToken.CallOpts)
}

// TotalRewardsOffered is a free data retrieval call binding the contract method 0xb6bbd864.
//
// Solidity: function totalRewardsOffered() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCallerSession) TotalRewardsOffered() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalRewardsOffered(&_CoinFidToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinFidToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenSession) TotalSupply() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalSupply(&_CoinFidToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CoinFidToken *CoinFidTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CoinFidToken.Contract.TotalSupply(&_CoinFidToken.CallOpts)
}

// ChangeCoinsForReward is a paid mutator transaction binding the contract method 0xf9328975.
//
// Solidity: function changeCoinsForReward(_from address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactor) ChangeCoinsForReward(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "changeCoinsForReward", _from, _value)
}

// ChangeCoinsForReward is a paid mutator transaction binding the contract method 0xf9328975.
//
// Solidity: function changeCoinsForReward(_from address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenSession) ChangeCoinsForReward(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.ChangeCoinsForReward(&_CoinFidToken.TransactOpts, _from, _value)
}

// ChangeCoinsForReward is a paid mutator transaction binding the contract method 0xf9328975.
//
// Solidity: function changeCoinsForReward(_from address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) ChangeCoinsForReward(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.ChangeCoinsForReward(&_CoinFidToken.TransactOpts, _from, _value)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactor) Distribute(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "distribute", _to, _value)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenSession) Distribute(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Distribute(&_CoinFidToken.TransactOpts, _to, _value)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) Distribute(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Distribute(&_CoinFidToken.TransactOpts, _to, _value)
}

// Generate is a paid mutator transaction binding the contract method 0x4a7dd523.
//
// Solidity: function generate(_value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactor) Generate(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "generate", _value)
}

// Generate is a paid mutator transaction binding the contract method 0x4a7dd523.
//
// Solidity: function generate(_value uint256) returns()
func (_CoinFidToken *CoinFidTokenSession) Generate(_value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Generate(&_CoinFidToken.TransactOpts, _value)
}

// Generate is a paid mutator transaction binding the contract method 0x4a7dd523.
//
// Solidity: function generate(_value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) Generate(_value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Generate(&_CoinFidToken.TransactOpts, _value)
}

// NameAdmin is a paid mutator transaction binding the contract method 0xebe23006.
//
// Solidity: function nameAdmin(_newAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) NameAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "nameAdmin", _newAdmin)
}

// NameAdmin is a paid mutator transaction binding the contract method 0xebe23006.
//
// Solidity: function nameAdmin(_newAdmin address) returns()
func (_CoinFidToken *CoinFidTokenSession) NameAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameAdmin(&_CoinFidToken.TransactOpts, _newAdmin)
}

// NameAdmin is a paid mutator transaction binding the contract method 0xebe23006.
//
// Solidity: function nameAdmin(_newAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) NameAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameAdmin(&_CoinFidToken.TransactOpts, _newAdmin)
}

// NameDistributor is a paid mutator transaction binding the contract method 0x26ed13fc.
//
// Solidity: function nameDistributor(_newDistributor address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) NameDistributor(opts *bind.TransactOpts, _newDistributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "nameDistributor", _newDistributor)
}

// NameDistributor is a paid mutator transaction binding the contract method 0x26ed13fc.
//
// Solidity: function nameDistributor(_newDistributor address) returns()
func (_CoinFidToken *CoinFidTokenSession) NameDistributor(_newDistributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameDistributor(&_CoinFidToken.TransactOpts, _newDistributor)
}

// NameDistributor is a paid mutator transaction binding the contract method 0x26ed13fc.
//
// Solidity: function nameDistributor(_newDistributor address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) NameDistributor(_newDistributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameDistributor(&_CoinFidToken.TransactOpts, _newDistributor)
}

// NameSuperAdmin is a paid mutator transaction binding the contract method 0x260e8457.
//
// Solidity: function nameSuperAdmin(_newSuperAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) NameSuperAdmin(opts *bind.TransactOpts, _newSuperAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "nameSuperAdmin", _newSuperAdmin)
}

// NameSuperAdmin is a paid mutator transaction binding the contract method 0x260e8457.
//
// Solidity: function nameSuperAdmin(_newSuperAdmin address) returns()
func (_CoinFidToken *CoinFidTokenSession) NameSuperAdmin(_newSuperAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameSuperAdmin(&_CoinFidToken.TransactOpts, _newSuperAdmin)
}

// NameSuperAdmin is a paid mutator transaction binding the contract method 0x260e8457.
//
// Solidity: function nameSuperAdmin(_newSuperAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) NameSuperAdmin(_newSuperAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.NameSuperAdmin(&_CoinFidToken.TransactOpts, _newSuperAdmin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_admin address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_admin address) returns()
func (_CoinFidToken *CoinFidTokenSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveAdmin(&_CoinFidToken.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_admin address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveAdmin(&_CoinFidToken.TransactOpts, _admin)
}

// RemoveDistributor is a paid mutator transaction binding the contract method 0x57c1f9e2.
//
// Solidity: function removeDistributor(_distributor address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) RemoveDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "removeDistributor", _distributor)
}

// RemoveDistributor is a paid mutator transaction binding the contract method 0x57c1f9e2.
//
// Solidity: function removeDistributor(_distributor address) returns()
func (_CoinFidToken *CoinFidTokenSession) RemoveDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveDistributor(&_CoinFidToken.TransactOpts, _distributor)
}

// RemoveDistributor is a paid mutator transaction binding the contract method 0x57c1f9e2.
//
// Solidity: function removeDistributor(_distributor address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) RemoveDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveDistributor(&_CoinFidToken.TransactOpts, _distributor)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(_superAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactor) RemoveSuperAdmin(opts *bind.TransactOpts, _superAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "removeSuperAdmin", _superAdmin)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(_superAdmin address) returns()
func (_CoinFidToken *CoinFidTokenSession) RemoveSuperAdmin(_superAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveSuperAdmin(&_CoinFidToken.TransactOpts, _superAdmin)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(_superAdmin address) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) RemoveSuperAdmin(_superAdmin common.Address) (*types.Transaction, error) {
	return _CoinFidToken.Contract.RemoveSuperAdmin(&_CoinFidToken.TransactOpts, _superAdmin)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_newTokenName string) returns()
func (_CoinFidToken *CoinFidTokenTransactor) SetName(opts *bind.TransactOpts, _newTokenName string) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "setName", _newTokenName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_newTokenName string) returns()
func (_CoinFidToken *CoinFidTokenSession) SetName(_newTokenName string) (*types.Transaction, error) {
	return _CoinFidToken.Contract.SetName(&_CoinFidToken.TransactOpts, _newTokenName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_newTokenName string) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) SetName(_newTokenName string) (*types.Transaction, error) {
	return _CoinFidToken.Contract.SetName(&_CoinFidToken.TransactOpts, _newTokenName)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xba51b1b4.
//
// Solidity: function setTokenSymbol(_newTokenSymbol string) returns()
func (_CoinFidToken *CoinFidTokenTransactor) SetTokenSymbol(opts *bind.TransactOpts, _newTokenSymbol string) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "setTokenSymbol", _newTokenSymbol)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xba51b1b4.
//
// Solidity: function setTokenSymbol(_newTokenSymbol string) returns()
func (_CoinFidToken *CoinFidTokenSession) SetTokenSymbol(_newTokenSymbol string) (*types.Transaction, error) {
	return _CoinFidToken.Contract.SetTokenSymbol(&_CoinFidToken.TransactOpts, _newTokenSymbol)
}

// SetTokenSymbol is a paid mutator transaction binding the contract method 0xba51b1b4.
//
// Solidity: function setTokenSymbol(_newTokenSymbol string) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) SetTokenSymbol(_newTokenSymbol string) (*types.Transaction, error) {
	return _CoinFidToken.Contract.SetTokenSymbol(&_CoinFidToken.TransactOpts, _newTokenSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Transfer(&_CoinFidToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_CoinFidToken *CoinFidTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CoinFidToken.Contract.Transfer(&_CoinFidToken.TransactOpts, _to, _value)
}

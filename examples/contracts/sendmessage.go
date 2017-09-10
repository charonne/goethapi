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

// SendMessageABI is the input ABI used to generate the binding from.
const SendMessageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"name\":\"text\",\"type\":\"string\"},{\"name\":\"author\",\"type\":\"string\"},{\"name\":\"categories\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"text\",\"type\":\"string\"},{\"name\":\"author\",\"type\":\"string\"},{\"name\":\"categories\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"text\",\"type\":\"string\"},{\"name\":\"author\",\"type\":\"string\"},{\"name\":\"categories\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// SendMessageBin is the compiled bytecode used for deploying new contracts.
const SendMessageBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b610a018061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630d80fefd811461005357806385d01a6b146102dd578063ec2ad1411461043d575b600080fd5b341561005e57600080fd5b6100696004356105a8565b60405160a08082528654600260001961010060018416150201909116049082018190528190602082019060408301906060840190608085019060c08601908c9080156100f65780601f106100cb576101008083540402835291602001916100f6565b820191906000526020600020905b8154815290600101906020018083116100d957829003601f168201915b505086810385528a54600260001961010060018416150201909116048082526020909101908b90801561016a5780601f1061013f5761010080835404028352916020019161016a565b820191906000526020600020905b81548152906001019060200180831161014d57829003601f168201915b505086810384528954600260001961010060018416150201909116048082526020909101908a9080156101de5780601f106101b3576101008083540402835291602001916101de565b820191906000526020600020905b8154815290600101906020018083116101c157829003601f168201915b50508681038352885460026000196101006001841615020190911604808252602090910190899080156102525780601f1061022757610100808354040283529160200191610252565b820191906000526020600020905b81548152906001019060200180831161023557829003601f168201915b50508681038252875460026000196101006001841615020190911604808252602090910190889080156102c65780601f1061029b576101008083540402835291602001916102c6565b820191906000526020600020905b8154815290600101906020018083116102a957829003601f168201915b50509a505050505050505050505060405180910390f35b34156102e857600080fd5b61043b600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506105dd95505050505050565b005b341561044857600080fd5b61059660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506106e695505050505050565b60405190815260200160405180910390f35b60018054829081106105b657fe5b906000526020600020906005020160005b5090506001810160028201600383016004840185565b6105e56107fa565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461060d57600080fd5b60a0604051908101604052808781526020018681526020018581526020018481526020018381525090508060018881548110151561064757fe5b906000526020600020906005020160005b5081518190805161066d929160200190610847565b50602082015181600101908051610688929160200190610847565b506040820151816002019080516106a3929160200190610847565b506060820151816003019080516106be929160200190610847565b506080820151816004019080516106d9929160200190610847565b5050505b50505050505050565b60006106f06107fa565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461071857600080fd5b60a0604051908101604052808881526020018781526020018681526020018581526020018481525090506001805480600101828161075691906108c6565b916000526020600020906005020160005b50829081518190805161077e929160200190610847565b50602082015181600101908051610799929160200190610847565b506040820151816002019080516107b4929160200190610847565b506060820151816003019080516107cf929160200190610847565b506080820151816004019080516107ea929160200190610847565b50505091505b5095945050505050565b60a06040519081016040528061080e6108f8565b815260200161081b6108f8565b81526020016108286108f8565b81526020016108356108f8565b81526020016108426108f8565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061088857805160ff19168380011785556108b5565b828001600101855582156108b5579182015b828111156108b557825182559160200191906001019061089a565b5b506108c292915061090a565b5090565b8154818355818115116108f2576005028160050283600052602060002091820191016108f2919061092b565b5b505050565b60206040519081016040526000815290565b61092891905b808211156108c25760008155600101610910565b5090565b90565b61092891905b808211156108c2576000610945828261098d565b61095360018301600061098d565b61096160028301600061098d565b61096f60038301600061098d565b61097d60048301600061098d565b50600501610931565b5090565b90565b50805460018160011615610100020316600290046000825580601f106109b357506109d1565b601f0160209004906000526020600020908101906109d1919061090a565b5b505600a165627a7a7230582087a22649e49df2843c0f43d8fd31196f68035a9dd754ade42533d3850b9998b80029`

// DeploySendMessage deploys a new Ethereum contract, binding an instance of SendMessage to it.
func DeploySendMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SendMessage, error) {
	parsed, err := abi.JSON(strings.NewReader(SendMessageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SendMessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SendMessage{SendMessageCaller: SendMessageCaller{contract: contract}, SendMessageTransactor: SendMessageTransactor{contract: contract}}, nil
}

// SendMessage is an auto generated Go binding around an Ethereum contract.
type SendMessage struct {
	SendMessageCaller     // Read-only binding to the contract
	SendMessageTransactor // Write-only binding to the contract
}

// SendMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SendMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SendMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SendMessageSession struct {
	Contract     *SendMessage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SendMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SendMessageCallerSession struct {
	Contract *SendMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SendMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SendMessageTransactorSession struct {
	Contract     *SendMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SendMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SendMessageRaw struct {
	Contract *SendMessage // Generic contract binding to access the raw methods on
}

// SendMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SendMessageCallerRaw struct {
	Contract *SendMessageCaller // Generic read-only contract binding to access the raw methods on
}

// SendMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SendMessageTransactorRaw struct {
	Contract *SendMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSendMessage creates a new instance of SendMessage, bound to a specific deployed contract.
func NewSendMessage(address common.Address, backend bind.ContractBackend) (*SendMessage, error) {
	contract, err := bindSendMessage(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SendMessage{SendMessageCaller: SendMessageCaller{contract: contract}, SendMessageTransactor: SendMessageTransactor{contract: contract}}, nil
}

// NewSendMessageCaller creates a new read-only instance of SendMessage, bound to a specific deployed contract.
func NewSendMessageCaller(address common.Address, caller bind.ContractCaller) (*SendMessageCaller, error) {
	contract, err := bindSendMessage(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SendMessageCaller{contract: contract}, nil
}

// NewSendMessageTransactor creates a new write-only instance of SendMessage, bound to a specific deployed contract.
func NewSendMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*SendMessageTransactor, error) {
	contract, err := bindSendMessage(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SendMessageTransactor{contract: contract}, nil
}

// bindSendMessage binds a generic wrapper to an already deployed contract.
func bindSendMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SendMessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendMessage *SendMessageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SendMessage.Contract.SendMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendMessage *SendMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendMessage.Contract.SendMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendMessage *SendMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendMessage.Contract.SendMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendMessage *SendMessageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SendMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendMessage *SendMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendMessage *SendMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendMessage.Contract.contract.Transact(opts, method, params...)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages( uint256) constant returns(text string, author string, categories string, url string, hash string)
func (_SendMessage *SendMessageCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Text       string
	Author     string
	Categories string
	Url        string
	Hash       string
}, error) {
	ret := new(struct {
		Text       string
		Author     string
		Categories string
		Url        string
		Hash       string
	})
	out := ret
	err := _SendMessage.contract.Call(opts, out, "messages", arg0)
	return *ret, err
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages( uint256) constant returns(text string, author string, categories string, url string, hash string)
func (_SendMessage *SendMessageSession) Messages(arg0 *big.Int) (struct {
	Text       string
	Author     string
	Categories string
	Url        string
	Hash       string
}, error) {
	return _SendMessage.Contract.Messages(&_SendMessage.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages( uint256) constant returns(text string, author string, categories string, url string, hash string)
func (_SendMessage *SendMessageCallerSession) Messages(arg0 *big.Int) (struct {
	Text       string
	Author     string
	Categories string
	Url        string
	Hash       string
}, error) {
	return _SendMessage.Contract.Messages(&_SendMessage.CallOpts, arg0)
}

// AddMessage is a paid mutator transaction binding the contract method 0xec2ad141.
//
// Solidity: function addMessage(text string, author string, categories string, url string, hash string) returns(length uint256)
func (_SendMessage *SendMessageTransactor) AddMessage(opts *bind.TransactOpts, text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.contract.Transact(opts, "addMessage", text, author, categories, url, hash)
}

// AddMessage is a paid mutator transaction binding the contract method 0xec2ad141.
//
// Solidity: function addMessage(text string, author string, categories string, url string, hash string) returns(length uint256)
func (_SendMessage *SendMessageSession) AddMessage(text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.Contract.AddMessage(&_SendMessage.TransactOpts, text, author, categories, url, hash)
}

// AddMessage is a paid mutator transaction binding the contract method 0xec2ad141.
//
// Solidity: function addMessage(text string, author string, categories string, url string, hash string) returns(length uint256)
func (_SendMessage *SendMessageTransactorSession) AddMessage(text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.Contract.AddMessage(&_SendMessage.TransactOpts, text, author, categories, url, hash)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x85d01a6b.
//
// Solidity: function updateMessage(id uint256, text string, author string, categories string, url string, hash string) returns()
func (_SendMessage *SendMessageTransactor) UpdateMessage(opts *bind.TransactOpts, id *big.Int, text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.contract.Transact(opts, "updateMessage", id, text, author, categories, url, hash)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x85d01a6b.
//
// Solidity: function updateMessage(id uint256, text string, author string, categories string, url string, hash string) returns()
func (_SendMessage *SendMessageSession) UpdateMessage(id *big.Int, text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.Contract.UpdateMessage(&_SendMessage.TransactOpts, id, text, author, categories, url, hash)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x85d01a6b.
//
// Solidity: function updateMessage(id uint256, text string, author string, categories string, url string, hash string) returns()
func (_SendMessage *SendMessageTransactorSession) UpdateMessage(id *big.Int, text string, author string, categories string, url string, hash string) (*types.Transaction, error) {
	return _SendMessage.Contract.UpdateMessage(&_SendMessage.TransactOpts, id, text, author, categories, url, hash)
}

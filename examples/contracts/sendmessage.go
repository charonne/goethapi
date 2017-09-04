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
const SendMessageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"name\":\"text\",\"type\":\"string\"},{\"name\":\"author\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"text\",\"type\":\"string\"},{\"name\":\"author\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// SendMessageBin is the compiled bytecode used for deploying new contracts.
const SendMessageBin = `0x6060604052341561000f57600080fd5b5b6104538061001f6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630d80fefd8114610048578063631e0c6914610161575b600080fd5b341561005357600080fd5b61005e6004356101f6565b6040516040808252835460026000196101006001841615020190911604908201819052819060208201906060830190869080156100dc5780601f106100b1576101008083540402835291602001916100dc565b820191906000526020600020905b8154815290600101906020018083116100bf57829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156101505780601f1061012557610100808354040283529160200191610150565b820191906000526020600020905b81548152906001019060200180831161013357829003601f168201915b505094505050505060405180910390f35b341561016c57600080fd5b6101f460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061021f95505050505050565b005b600080548290811061020457fe5b906000526020600020906002020160005b5090506001810182565b61022761029e565b604080519081016040528381526020810183905260008054919250906001810161025183826102c3565b916000526020600020906002020160005b5082908151819080516102799291602001906102f5565b506020820151816001019080516102949291602001906102f5565b505050505b505050565b60408051908101604052806102b1610374565b81526020016102be610374565b905290565b815481835581811511610299576002028160020283600052602060002091820191016102999190610386565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061033657805160ff1916838001178555610363565b82800160010185558215610363579182015b82811115610363578251825591602001919060010190610348565b5b506103709291506103be565b5090565b60206040519081016040526000815290565b6103bb91905b808211156103705760006103a082826103df565b6103ae6001830160006103df565b5060020161038c565b5090565b90565b6103bb91905b8082111561037057600081556001016103c4565b5090565b90565b50805460018160011615610100020316600290046000825580601f106104055750610423565b601f01602090049060005260206000209081019061042391906103be565b5b505600a165627a7a7230582027f377642590ab02d4798e1f8d22528e716e4f593bd614d005d97bf62cdff10b0029`

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
// Solidity: function messages( uint256) constant returns(text string, author string)
func (_SendMessage *SendMessageCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Text   string
	Author string
}, error) {
	ret := new(struct {
		Text   string
		Author string
	})
	out := ret
	err := _SendMessage.contract.Call(opts, out, "messages", arg0)
	return *ret, err
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages( uint256) constant returns(text string, author string)
func (_SendMessage *SendMessageSession) Messages(arg0 *big.Int) (struct {
	Text   string
	Author string
}, error) {
	return _SendMessage.Contract.Messages(&_SendMessage.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages( uint256) constant returns(text string, author string)
func (_SendMessage *SendMessageCallerSession) Messages(arg0 *big.Int) (struct {
	Text   string
	Author string
}, error) {
	return _SendMessage.Contract.Messages(&_SendMessage.CallOpts, arg0)
}

// AddMessage is a paid mutator transaction binding the contract method 0x631e0c69.
//
// Solidity: function addMessage(text string, author string) returns()
func (_SendMessage *SendMessageTransactor) AddMessage(opts *bind.TransactOpts, text string, author string) (*types.Transaction, error) {
	return _SendMessage.contract.Transact(opts, "addMessage", text, author)
}

// AddMessage is a paid mutator transaction binding the contract method 0x631e0c69.
//
// Solidity: function addMessage(text string, author string) returns()
func (_SendMessage *SendMessageSession) AddMessage(text string, author string) (*types.Transaction, error) {
	return _SendMessage.Contract.AddMessage(&_SendMessage.TransactOpts, text, author)
}

// AddMessage is a paid mutator transaction binding the contract method 0x631e0c69.
//
// Solidity: function addMessage(text string, author string) returns()
func (_SendMessage *SendMessageTransactorSession) AddMessage(text string, author string) (*types.Transaction, error) {
	return _SendMessage.Contract.AddMessage(&_SendMessage.TransactOpts, text, author)
}

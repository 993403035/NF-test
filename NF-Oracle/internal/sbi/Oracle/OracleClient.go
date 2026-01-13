
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Oracle

import (
"math/big"
"strings"

"github.com/FISCO-BCOS/go-sdk/abi"
"github.com/FISCO-BCOS/go-sdk/abi/bind"
"github.com/FISCO-BCOS/go-sdk/core/types"
ethereum "github.com/ethereum/go-ethereum"
"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// OracleClientABI is the input ABI used to generate the binding from.
const OracleClientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"oraclenode\",\"type\":\"address\"}],\"name\":\"RequestRandomValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oraclenode\",\"type\":\"address\"},{\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"RequestHttp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"random\",\"type\":\"uint256\"},{\"name\":\"pk_str\",\"type\":\"string\"},{\"name\":\"agg_sig_str\",\"type\":\"string\"},{\"name\":\"num_user\",\"type\":\"uint256\"}],\"name\":\"_callbackRandom\",\"outputs\":[{\"name\":\"result\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"httpresponse\",\"type\":\"string\"},{\"name\":\"pk_str\",\"type\":\"string\"},{\"name\":\"agg_sig_str\",\"type\":\"string\"},{\"name\":\"num_user\",\"type\":\"uint256\"}],\"name\":\"_callbackHttp\",\"outputs\":[{\"name\":\"result\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"Randomcallbackevent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oraclenode\",\"type\":\"address\"}],\"name\":\"RequestRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oraclenode\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"RequestHttpevent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"httpresponse\",\"type\":\"string\"}],\"name\":\"Httpcallbackevent\",\"type\":\"event\"}]"

// OracleClientBin is the compiled bytecode used for deploying new contracts.
var OracleClientBin = "0x608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304df0ccc1461006757806377763693146100aa57806381e30b2a14610133578063d069134a1461026f575b600080fd5b34801561007357600080fd5b506100a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103e7565b005b3480156100b657600080fd5b50610131600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061044d565b005b34801561013f57600080fd5b506101f460048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050610521565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610234578082015181840152602081019050610219565b50505050905090810190601f1680156102615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027b57600080fd5b5061036c600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919050505061075f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103ac578082015181840152602081019050610391565b50505050905090810190601f1680156103d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b7fc8c4f6fdb4b0d03c6e5dfeb57792ace90c48d2911129d7f5642442a56d95d68081604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b7fb66bc29ea7db670d9a6c02689ca6b44bc234eab37ce9360fa4c8460115eed1778282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104e25780820151818401526020810190506104c7565b50505050905090810190601f16801561050f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b60607f44b8ce27103e293d017ea63b904b17f0435e6fbe1c332f01e018d5715c49dae3856040518082815260200191505060405180910390a16000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663817040148585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561060f5780820151818401526020810190506105f4565b50505050905090810190601f16801561063c5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561067557808201518184015260208101905061065a565b50505050905090810190601f1680156106a25780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156106c457600080fd5b505af11580156106d8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561070257600080fd5b81019080805164010000000081111561071a57600080fd5b8281019050602081018481111561073057600080fd5b815185600182028301116401000000008211171561074d57600080fd5b50509291905050509050949350505050565b60607f1b85a8b0172d4b653eb02dfa8b0491d39cbdc71727186a2b64ddbb6208ea0806856040518080602001828103825283818151815260200191508051906020019080838360005b838110156107c35780820151818401526020810190506107a8565b50505050905090810190601f1680156107f05780820380516001836020036101000a031916815260200191505b509250505060405180910390a16000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663817040148585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156108b2578082015181840152602081019050610897565b50505050905090810190601f1680156108df5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156109185780820151818401526020810190506108fd565b50505050905090810190601f1680156109455780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561096757600080fd5b505af115801561097b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156109a557600080fd5b8101908080516401000000008111156109bd57600080fd5b828101905060208101848111156109d357600080fd5b81518560018202830111640100000000821117156109f057600080fd5b505092919050505090509493505050505600a165627a7a72305820a0e7722a579ce90412be20819eec0c405a4dcd6554482d6679f1093cb3c2aa040029"

// DeployOracleClient deploys a new contract, binding an instance of OracleClient to it.
func DeployOracleClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OracleClient, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleClientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleClient{OracleClientCaller: OracleClientCaller{contract: contract}, OracleClientTransactor: OracleClientTransactor{contract: contract}, OracleClientFilterer: OracleClientFilterer{contract: contract}}, nil
}

func AsyncDeployOracleClient(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleClientABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(OracleClientBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// OracleClient is an auto generated Go binding around a Solidity contract.
type OracleClient struct {
	OracleClientCaller     // Read-only binding to the contract
	OracleClientTransactor // Write-only binding to the contract
	OracleClientFilterer   // Log filterer for contract events
}

// OracleClientCaller is an auto generated read-only Go binding around a Solidity contract.
type OracleClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleClientTransactor is an auto generated write-only Go binding around a Solidity contract.
type OracleClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleClientFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type OracleClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleClientSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type OracleClientSession struct {
	Contract     *OracleClient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleClientCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type OracleClientCallerSession struct {
	Contract *OracleClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OracleClientTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type OracleClientTransactorSession struct {
	Contract     *OracleClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OracleClientRaw is an auto generated low-level Go binding around a Solidity contract.
type OracleClientRaw struct {
	Contract *OracleClient // Generic contract binding to access the raw methods on
}

// OracleClientCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type OracleClientCallerRaw struct {
	Contract *OracleClientCaller // Generic read-only contract binding to access the raw methods on
}

// OracleClientTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type OracleClientTransactorRaw struct {
	Contract *OracleClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleClient creates a new instance of OracleClient, bound to a specific deployed contract.
func NewOracleClient(address common.Address, backend bind.ContractBackend) (*OracleClient, error) {
	contract, err := bindOracleClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleClient{OracleClientCaller: OracleClientCaller{contract: contract}, OracleClientTransactor: OracleClientTransactor{contract: contract}, OracleClientFilterer: OracleClientFilterer{contract: contract}}, nil
}

// NewOracleClientCaller creates a new read-only instance of OracleClient, bound to a specific deployed contract.
func NewOracleClientCaller(address common.Address, caller bind.ContractCaller) (*OracleClientCaller, error) {
	contract, err := bindOracleClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleClientCaller{contract: contract}, nil
}

// NewOracleClientTransactor creates a new write-only instance of OracleClient, bound to a specific deployed contract.
func NewOracleClientTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleClientTransactor, error) {
	contract, err := bindOracleClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleClientTransactor{contract: contract}, nil
}

// NewOracleClientFilterer creates a new log filterer instance of OracleClient, bound to a specific deployed contract.
func NewOracleClientFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleClientFilterer, error) {
	contract, err := bindOracleClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleClientFilterer{contract: contract}, nil
}

// bindOracleClient binds a generic wrapper to an already deployed contract.
func bindOracleClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleClient *OracleClientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleClient.Contract.OracleClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleClient *OracleClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.OracleClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleClient *OracleClientRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.OracleClientTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleClient *OracleClientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleClient *OracleClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleClient *OracleClientTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// RequestHttp is a paid mutator transaction binding the contract method 0x77763693.
//
// Solidity: function RequestHttp(address oraclenode, string domain) returns()
func (_OracleClient *OracleClientTransactor) RequestHttp(opts *bind.TransactOpts, oraclenode common.Address, domain string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _OracleClient.contract.TransactWithResult(opts, out, "RequestHttp", oraclenode, domain)
	return transaction, receipt, err
}

func (_OracleClient *OracleClientTransactor) AsyncRequestHttp(handler func(*types.Receipt, error), opts *bind.TransactOpts, oraclenode common.Address, domain string) (*types.Transaction, error) {
	return _OracleClient.contract.AsyncTransact(opts, handler, "RequestHttp", oraclenode, domain)
}

// RequestHttp is a paid mutator transaction binding the contract method 0x77763693.
//
// Solidity: function RequestHttp(address oraclenode, string domain) returns()
func (_OracleClient *OracleClientSession) RequestHttp(oraclenode common.Address, domain string) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.RequestHttp(&_OracleClient.TransactOpts, oraclenode, domain)
}

func (_OracleClient *OracleClientSession) AsyncRequestHttp(handler func(*types.Receipt, error), oraclenode common.Address, domain string) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncRequestHttp(handler, &_OracleClient.TransactOpts, oraclenode, domain)
}

// RequestHttp is a paid mutator transaction binding the contract method 0x77763693.
//
// Solidity: function RequestHttp(address oraclenode, string domain) returns()
func (_OracleClient *OracleClientTransactorSession) RequestHttp(oraclenode common.Address, domain string) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.RequestHttp(&_OracleClient.TransactOpts, oraclenode, domain)
}

func (_OracleClient *OracleClientTransactorSession) AsyncRequestHttp(handler func(*types.Receipt, error), oraclenode common.Address, domain string) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncRequestHttp(handler, &_OracleClient.TransactOpts, oraclenode, domain)
}

// RequestRandomValue is a paid mutator transaction binding the contract method 0x04df0ccc.
//
// Solidity: function RequestRandomValue(address oraclenode) returns()
func (_OracleClient *OracleClientTransactor) RequestRandomValue(opts *bind.TransactOpts, oraclenode common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _OracleClient.contract.TransactWithResult(opts, out, "RequestRandomValue", oraclenode)
	return transaction, receipt, err
}

func (_OracleClient *OracleClientTransactor) AsyncRequestRandomValue(handler func(*types.Receipt, error), opts *bind.TransactOpts, oraclenode common.Address) (*types.Transaction, error) {
	return _OracleClient.contract.AsyncTransact(opts, handler, "RequestRandomValue", oraclenode)
}

// RequestRandomValue is a paid mutator transaction binding the contract method 0x04df0ccc.
//
// Solidity: function RequestRandomValue(address oraclenode) returns()
func (_OracleClient *OracleClientSession) RequestRandomValue(oraclenode common.Address) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.RequestRandomValue(&_OracleClient.TransactOpts, oraclenode)
}

func (_OracleClient *OracleClientSession) AsyncRequestRandomValue(handler func(*types.Receipt, error), oraclenode common.Address) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncRequestRandomValue(handler, &_OracleClient.TransactOpts, oraclenode)
}

// RequestRandomValue is a paid mutator transaction binding the contract method 0x04df0ccc.
//
// Solidity: function RequestRandomValue(address oraclenode) returns()
func (_OracleClient *OracleClientTransactorSession) RequestRandomValue(oraclenode common.Address) (*types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.RequestRandomValue(&_OracleClient.TransactOpts, oraclenode)
}

func (_OracleClient *OracleClientTransactorSession) AsyncRequestRandomValue(handler func(*types.Receipt, error), oraclenode common.Address) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncRequestRandomValue(handler, &_OracleClient.TransactOpts, oraclenode)
}

// CallbackHttp is a paid mutator transaction binding the contract method 0xd069134a.
//
// Solidity: function _callbackHttp(string httpresponse, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientTransactor) CallbackHttp(opts *bind.TransactOpts, httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	transaction, receipt, err := _OracleClient.contract.TransactWithResult(opts, out, "_callbackHttp", httpresponse, pk_str, agg_sig_str, num_user)
	return *ret0, transaction, receipt, err
}

func (_OracleClient *OracleClientTransactor) AsyncCallbackHttp(handler func(*types.Receipt, error), opts *bind.TransactOpts, httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.contract.AsyncTransact(opts, handler, "_callbackHttp", httpresponse, pk_str, agg_sig_str, num_user)
}

// CallbackHttp is a paid mutator transaction binding the contract method 0xd069134a.
//
// Solidity: function _callbackHttp(string httpresponse, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientSession) CallbackHttp(httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.CallbackHttp(&_OracleClient.TransactOpts, httpresponse, pk_str, agg_sig_str, num_user)
}

func (_OracleClient *OracleClientSession) AsyncCallbackHttp(handler func(*types.Receipt, error), httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncCallbackHttp(handler, &_OracleClient.TransactOpts, httpresponse, pk_str, agg_sig_str, num_user)
}

// CallbackHttp is a paid mutator transaction binding the contract method 0xd069134a.
//
// Solidity: function _callbackHttp(string httpresponse, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientTransactorSession) CallbackHttp(httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.CallbackHttp(&_OracleClient.TransactOpts, httpresponse, pk_str, agg_sig_str, num_user)
}

func (_OracleClient *OracleClientTransactorSession) AsyncCallbackHttp(handler func(*types.Receipt, error), httpresponse string, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncCallbackHttp(handler, &_OracleClient.TransactOpts, httpresponse, pk_str, agg_sig_str, num_user)
}

// CallbackRandom is a paid mutator transaction binding the contract method 0x81e30b2a.
//
// Solidity: function _callbackRandom(uint256 random, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientTransactor) CallbackRandom(opts *bind.TransactOpts, random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	transaction, receipt, err := _OracleClient.contract.TransactWithResult(opts, out, "_callbackRandom", random, pk_str, agg_sig_str, num_user)
	return *ret0, transaction, receipt, err
}

func (_OracleClient *OracleClientTransactor) AsyncCallbackRandom(handler func(*types.Receipt, error), opts *bind.TransactOpts, random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.contract.AsyncTransact(opts, handler, "_callbackRandom", random, pk_str, agg_sig_str, num_user)
}

// CallbackRandom is a paid mutator transaction binding the contract method 0x81e30b2a.
//
// Solidity: function _callbackRandom(uint256 random, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientSession) CallbackRandom(random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.CallbackRandom(&_OracleClient.TransactOpts, random, pk_str, agg_sig_str, num_user)
}

func (_OracleClient *OracleClientSession) AsyncCallbackRandom(handler func(*types.Receipt, error), random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncCallbackRandom(handler, &_OracleClient.TransactOpts, random, pk_str, agg_sig_str, num_user)
}

// CallbackRandom is a paid mutator transaction binding the contract method 0x81e30b2a.
//
// Solidity: function _callbackRandom(uint256 random, string pk_str, string agg_sig_str, uint256 num_user) returns(string result)
func (_OracleClient *OracleClientTransactorSession) CallbackRandom(random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (string, *types.Transaction, *types.Receipt, error) {
	return _OracleClient.Contract.CallbackRandom(&_OracleClient.TransactOpts, random, pk_str, agg_sig_str, num_user)
}

func (_OracleClient *OracleClientTransactorSession) AsyncCallbackRandom(handler func(*types.Receipt, error), random *big.Int, pk_str string, agg_sig_str string, num_user *big.Int) (*types.Transaction, error) {
	return _OracleClient.Contract.AsyncCallbackRandom(handler, &_OracleClient.TransactOpts, random, pk_str, agg_sig_str, num_user)
}

// OracleClientHttpcallbackevent represents a Httpcallbackevent event raised by the OracleClient contract.
type OracleClientHttpcallbackevent struct {
	Httpresponse string
	Raw          types.Log // Blockchain specific contextual infos
}

// WatchHttpcallbackevent is a free log subscription operation binding the contract event 0x1b85a8b0172d4b653eb02dfa8b0491d39cbdc71727186a2b64ddbb6208ea0806.
//
// Solidity: event Httpcallbackevent(string httpresponse)
func (_OracleClient *OracleClientFilterer) WatchHttpcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "Httpcallbackevent")
}

func (_OracleClient *OracleClientFilterer) WatchAllHttpcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "Httpcallbackevent")
}

// ParseHttpcallbackevent is a log parse operation binding the contract event 0x1b85a8b0172d4b653eb02dfa8b0491d39cbdc71727186a2b64ddbb6208ea0806.
//
// Solidity: event Httpcallbackevent(string httpresponse)
func (_OracleClient *OracleClientFilterer) ParseHttpcallbackevent(log types.Log) (*OracleClientHttpcallbackevent, error) {
	event := new(OracleClientHttpcallbackevent)
	if err := _OracleClient.contract.UnpackLog(event, "Httpcallbackevent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchHttpcallbackevent is a free log subscription operation binding the contract event 0x1b85a8b0172d4b653eb02dfa8b0491d39cbdc71727186a2b64ddbb6208ea0806.
//
// Solidity: event Httpcallbackevent(string httpresponse)
func (_OracleClient *OracleClientSession) WatchHttpcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchHttpcallbackevent(fromBlock, handler)
}

func (_OracleClient *OracleClientSession) WatchAllHttpcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchAllHttpcallbackevent(fromBlock, handler)
}

// ParseHttpcallbackevent is a log parse operation binding the contract event 0x1b85a8b0172d4b653eb02dfa8b0491d39cbdc71727186a2b64ddbb6208ea0806.
//
// Solidity: event Httpcallbackevent(string httpresponse)
func (_OracleClient *OracleClientSession) ParseHttpcallbackevent(log types.Log) (*OracleClientHttpcallbackevent, error) {
	return _OracleClient.Contract.ParseHttpcallbackevent(log)
}

// OracleClientRandomcallbackevent represents a Randomcallbackevent event raised by the OracleClient contract.
type OracleClientRandomcallbackevent struct {
	Random *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// WatchRandomcallbackevent is a free log subscription operation binding the contract event 0x44b8ce27103e293d017ea63b904b17f0435e6fbe1c332f01e018d5715c49dae3.
//
// Solidity: event Randomcallbackevent(uint256 random)
func (_OracleClient *OracleClientFilterer) WatchRandomcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "Randomcallbackevent")
}

func (_OracleClient *OracleClientFilterer) WatchAllRandomcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "Randomcallbackevent")
}

// ParseRandomcallbackevent is a log parse operation binding the contract event 0x44b8ce27103e293d017ea63b904b17f0435e6fbe1c332f01e018d5715c49dae3.
//
// Solidity: event Randomcallbackevent(uint256 random)
func (_OracleClient *OracleClientFilterer) ParseRandomcallbackevent(log types.Log) (*OracleClientRandomcallbackevent, error) {
	event := new(OracleClientRandomcallbackevent)
	if err := _OracleClient.contract.UnpackLog(event, "Randomcallbackevent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRandomcallbackevent is a free log subscription operation binding the contract event 0x44b8ce27103e293d017ea63b904b17f0435e6fbe1c332f01e018d5715c49dae3.
//
// Solidity: event Randomcallbackevent(uint256 random)
func (_OracleClient *OracleClientSession) WatchRandomcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchRandomcallbackevent(fromBlock, handler)
}

func (_OracleClient *OracleClientSession) WatchAllRandomcallbackevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchAllRandomcallbackevent(fromBlock, handler)
}

// ParseRandomcallbackevent is a log parse operation binding the contract event 0x44b8ce27103e293d017ea63b904b17f0435e6fbe1c332f01e018d5715c49dae3.
//
// Solidity: event Randomcallbackevent(uint256 random)
func (_OracleClient *OracleClientSession) ParseRandomcallbackevent(log types.Log) (*OracleClientRandomcallbackevent, error) {
	return _OracleClient.Contract.ParseRandomcallbackevent(log)
}

// OracleClientRequestHttpevent represents a RequestHttpevent event raised by the OracleClient contract.
type OracleClientRequestHttpevent struct {
	Oraclenode common.Address
	Domain     string
	Raw        types.Log // Blockchain specific contextual infos
}

// WatchRequestHttpevent is a free log subscription operation binding the contract event 0xb66bc29ea7db670d9a6c02689ca6b44bc234eab37ce9360fa4c8460115eed177.
//
// Solidity: event RequestHttpevent(address oraclenode, string domain)
func (_OracleClient *OracleClientFilterer) WatchRequestHttpevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "RequestHttpevent")
}

func (_OracleClient *OracleClientFilterer) WatchAllRequestHttpevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "RequestHttpevent")
}

// ParseRequestHttpevent is a log parse operation binding the contract event 0xb66bc29ea7db670d9a6c02689ca6b44bc234eab37ce9360fa4c8460115eed177.
//
// Solidity: event RequestHttpevent(address oraclenode, string domain)
func (_OracleClient *OracleClientFilterer) ParseRequestHttpevent(log types.Log) (*OracleClientRequestHttpevent, error) {
	event := new(OracleClientRequestHttpevent)
	if err := _OracleClient.contract.UnpackLog(event, "RequestHttpevent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRequestHttpevent is a free log subscription operation binding the contract event 0xb66bc29ea7db670d9a6c02689ca6b44bc234eab37ce9360fa4c8460115eed177.
//
// Solidity: event RequestHttpevent(address oraclenode, string domain)
func (_OracleClient *OracleClientSession) WatchRequestHttpevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchRequestHttpevent(fromBlock, handler)
}

func (_OracleClient *OracleClientSession) WatchAllRequestHttpevent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchAllRequestHttpevent(fromBlock, handler)
}

// ParseRequestHttpevent is a log parse operation binding the contract event 0xb66bc29ea7db670d9a6c02689ca6b44bc234eab37ce9360fa4c8460115eed177.
//
// Solidity: event RequestHttpevent(address oraclenode, string domain)
func (_OracleClient *OracleClientSession) ParseRequestHttpevent(log types.Log) (*OracleClientRequestHttpevent, error) {
	return _OracleClient.Contract.ParseRequestHttpevent(log)
}

// OracleClientRequestRandom represents a RequestRandom event raised by the OracleClient contract.
type OracleClientRequestRandom struct {
	Oraclenode common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// WatchRequestRandom is a free log subscription operation binding the contract event 0xc8c4f6fdb4b0d03c6e5dfeb57792ace90c48d2911129d7f5642442a56d95d680.
//
// Solidity: event RequestRandom(address oraclenode)
func (_OracleClient *OracleClientFilterer) WatchRequestRandom(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "RequestRandom")
}

func (_OracleClient *OracleClientFilterer) WatchAllRequestRandom(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.contract.WatchLogs(fromBlock, handler, "RequestRandom")
}

// ParseRequestRandom is a log parse operation binding the contract event 0xc8c4f6fdb4b0d03c6e5dfeb57792ace90c48d2911129d7f5642442a56d95d680.
//
// Solidity: event RequestRandom(address oraclenode)
func (_OracleClient *OracleClientFilterer) ParseRequestRandom(log types.Log) (*OracleClientRequestRandom, error) {
	event := new(OracleClientRequestRandom)
	if err := _OracleClient.contract.UnpackLog(event, "RequestRandom", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRequestRandom is a free log subscription operation binding the contract event 0xc8c4f6fdb4b0d03c6e5dfeb57792ace90c48d2911129d7f5642442a56d95d680.
//
// Solidity: event RequestRandom(address oraclenode)
func (_OracleClient *OracleClientSession) WatchRequestRandom(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchRequestRandom(fromBlock, handler)
}

func (_OracleClient *OracleClientSession) WatchAllRequestRandom(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _OracleClient.Contract.WatchAllRequestRandom(fromBlock, handler)
}

// ParseRequestRandom is a log parse operation binding the contract event 0xc8c4f6fdb4b0d03c6e5dfeb57792ace90c48d2911129d7f5642442a56d95d680.
//
// Solidity: event RequestRandom(address oraclenode)
func (_OracleClient *OracleClientSession) ParseRequestRandom(log types.Log) (*OracleClientRequestRandom, error) {
	return _OracleClient.Contract.ParseRequestRandom(log)
}

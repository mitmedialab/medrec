// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AgentABI is the input ABI used to generate the binding from.
const AgentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeCustodian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"custodianEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"custodians\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"enableAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumPermissions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"enableCustodian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"name\":\"addRelationship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addCustodian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relationships\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"agentEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumEnabledOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumRelationships\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"agent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumCustodians\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"bytes32\"}],\"name\":\"addPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AgentBin is the compiled bytecode used for deploying new contracts.
const AgentBin = `0x6060604052341561000f57600080fd5b6000805474010000000000000000000000000000000000000000600160a060020a031990911633600160a060020a03161760a060020a60ff021916178155610cf290819061005d90396000f3006060604052600436106100f05763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630f12a66a81146100f5578063108fd0cc146101165780632299bda6146101405780632788aa711461017257806358886dba1461018a5780636121785c146101b25780636ca9677a146101c557806379bda937146101e957806385d95b811461020857806394aa27181461022757806399d29e711461023d578063bcf685ed14610250578063c23a0d281461026f578063e968177e14610282578063f5ff5c7614610295578063faea7b2c146102a8578063fd17ec56146102bb575b600080fd5b341561010057600080fd5b610114600160a060020a03600435166102d1565b005b341561012157600080fd5b61012c60043561057a565b604051901515815260200160405180910390f35b341561014b57600080fd5b6101566004356105ac565b604051600160a060020a03909116815260200160405180910390f35b341561017d57600080fd5b61011460043515156105d4565b341561019557600080fd5b6101a06004356106d2565b60405190815260200160405180910390f35b34156101bd57600080fd5b6101a06106f1565b34156101d057600080fd5b610114600160a060020a036004351660243515156106f8565b34156101f457600080fd5b610114600160a060020a0360043516610789565b341561021357600080fd5b610114600160a060020a0360043516610896565b341561023257600080fd5b6101566004356109d6565b341561024857600080fd5b61012c6109e4565b341561025b57600080fd5b610114600160a060020a03600435166109f4565b341561027a57600080fd5b6101a0610ae4565b341561028d57600080fd5b6101a0610b54565b34156102a057600080fd5b610156610b5a565b34156102b357600080fd5b6101a0610b69565b34156102c657600080fd5b610114600435610b6f565b600080600080600060149054906101000a900460ff168015610301575060005433600160a060020a039081169116145b1561030b57600191505b5060005b60015481101561038a57600280548290811061032757fe5b600091825260209182902082820401549190066101000a900460ff1680156103745750600180548290811061035857fe5b60009182526020909120015433600160a060020a039081169116145b15610382576001915061038a565b60010161030f565b81151561039657600080fd5b60009350600092505b6001548310156104cb5783156104875760018054849081106103bd57fe5b60009182526020909120015460018054600160a060020a039092169160001986019081106103e757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600280548490811061042b57fe5b90600052602060002090602091828204019190069054906101000a900460ff1660026001850381548110151561045d57fe5b90600052602060002090602091828204019190066101000a81548160ff0219169083151502179055505b84600160a060020a03166001848154811015156104a057fe5b600091825260209091200154600160a060020a031614156104c057600193505b60019092019161039f565b6001805460001981019081106104dd57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916905560028054600019810190811061051557fe5b6000918252602091829020828204018054929091066101000a60ff0219909116905560018054600019019061054a9082610c54565b5060028054600019019061055e9082610c78565b506000610569610ae4565b1161057357600080fd5b5050505050565b600280548290811061058857fe5b9060005260206000209060209182820401919006915054906101000a900460ff1681565b60018054829081106105ba57fe5b600091825260209091200154600160a060020a0316905081565b60008054819060a060020a900460ff1680156105fe575060005433600160a060020a039081169116145b1561060857600191505b5060005b60015481101561068757600280548290811061062457fe5b600091825260209182902082820401549190066101000a900460ff1680156106715750600180548290811061065557fe5b60009182526020909120015433600160a060020a039081169116145b1561067f5760019150610687565b60010161060c565b81151561069357600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a851515021781556106c3610ae4565b116106cd57600080fd5b505050565b60038054829081106106e057fe5b600091825260209091200154905081565b6003545b90565b60005b60015481101561077f5782600160a060020a031660018281548110151561071e57fe5b600091825260209091200154600160a060020a03161415610777578160028281548110151561074957fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555061077f565b6001016106fb565b60006106c3610ae4565b60008054819060a060020a900460ff1680156107b3575060005433600160a060020a039081169116145b156107bd57600191505b5060005b60015481101561083c5760028054829081106107d957fe5b600091825260209182902082820401549190066101000a900460ff1680156108265750600180548290811061080a57fe5b60009182526020909120015433600160a060020a039081169116145b15610834576001915061083c565b6001016107c1565b81151561084857600080fd5b600480546001810161085a8382610c54565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0394909416939093179092555050565b60008054819060a060020a900460ff1680156108c0575060005433600160a060020a039081169116145b156108ca57600191505b5060005b6001548110156109495760028054829081106108e657fe5b600091825260209182902082820401549190066101000a900460ff1680156109335750600180548290811061091757fe5b60009182526020909120015433600160a060020a039081169116145b156109415760019150610949565b6001016108ce565b81151561095557600080fd5b600180548082016109668382610c54565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851617905560028054600181016109aa8382610c78565b506000918252602091829020828204018054929091066101000a60ff8102199092169091179055505050565b60048054829081106105ba57fe5b60005460a060020a900460ff1681565b60008054819060a060020a900460ff168015610a1e575060005433600160a060020a039081169116145b15610a2857600191505b5060005b600154811015610aa7576002805482908110610a4457fe5b600091825260209182902082820401549190066101000a900460ff168015610a9157506001805482908110610a7557fe5b60009182526020909120015433600160a060020a039081169116145b15610a9f5760019150610aa7565b600101610a2c565b811515610ab357600080fd5b50506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600080548190819060a060020a900460ff1615610b02576001909101905b5060005b600254811015610b4e576002805482908110610b1e57fe5b600091825260209182902082820401549190066101000a900460ff1615610b46576001909101905b600101610b06565b50919050565b60045490565b600054600160a060020a031681565b60015490565b60008054819060a060020a900460ff168015610b99575060005433600160a060020a039081169116145b15610ba357600191505b5060005b600154811015610c22576002805482908110610bbf57fe5b600091825260209182902082820401549190066101000a900460ff168015610c0c57506001805482908110610bf057fe5b60009182526020909120015433600160a060020a039081169116145b15610c1a5760019150610c22565b600101610ba7565b811515610c2e57600080fd5b6003805460018101610c408382610c54565b506000918252602090912001929092555050565b8154818355818115116106cd576000838152602090206106cd918101908301610ca8565b8154818355818115116106cd57601f016020900481601f016020900483600052602060002091820191016106cd91905b6106f591905b80821115610cc25760008155600101610cae565b50905600a165627a7a723058204f506fe73078dd5b3943b45ff11b8571fef2822675e6d7703c99cde42bfab5080029`

// DeployAgent deploys a new Ethereum contract, binding an instance of Agent to it.
func DeployAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agent, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agent{AgentCaller: AgentCaller{contract: contract}, AgentTransactor: AgentTransactor{contract: contract}}, nil
}

// Agent is an auto generated Go binding around an Ethereum contract.
type Agent struct {
	AgentCaller     // Read-only binding to the contract
	AgentTransactor // Write-only binding to the contract
}

// AgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSession struct {
	Contract     *Agent            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentCallerSession struct {
	Contract *AgentCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentTransactorSession struct {
	Contract     *AgentTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRaw struct {
	Contract *Agent // Generic contract binding to access the raw methods on
}

// AgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentCallerRaw struct {
	Contract *AgentCaller // Generic read-only contract binding to access the raw methods on
}

// AgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentTransactorRaw struct {
	Contract *AgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgent creates a new instance of Agent, bound to a specific deployed contract.
func NewAgent(address common.Address, backend bind.ContractBackend) (*Agent, error) {
	contract, err := bindAgent(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agent{AgentCaller: AgentCaller{contract: contract}, AgentTransactor: AgentTransactor{contract: contract}}, nil
}

// NewAgentCaller creates a new read-only instance of Agent, bound to a specific deployed contract.
func NewAgentCaller(address common.Address, caller bind.ContractCaller) (*AgentCaller, error) {
	contract, err := bindAgent(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCaller{contract: contract}, nil
}

// NewAgentTransactor creates a new write-only instance of Agent, bound to a specific deployed contract.
func NewAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentTransactor, error) {
	contract, err := bindAgent(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AgentTransactor{contract: contract}, nil
}

// bindAgent binds a generic wrapper to an already deployed contract.
func bindAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.AgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transact(opts, method, params...)
}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() constant returns(address)
func (_Agent *AgentCaller) Agent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "agent")
	return *ret0, err
}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() constant returns(address)
func (_Agent *AgentSession) Agent() (common.Address, error) {
	return _Agent.Contract.Agent(&_Agent.CallOpts)
}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() constant returns(address)
func (_Agent *AgentCallerSession) Agent() (common.Address, error) {
	return _Agent.Contract.Agent(&_Agent.CallOpts)
}

// AgentEnabled is a free data retrieval call binding the contract method 0x99d29e71.
//
// Solidity: function agentEnabled() constant returns(bool)
func (_Agent *AgentCaller) AgentEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "agentEnabled")
	return *ret0, err
}

// AgentEnabled is a free data retrieval call binding the contract method 0x99d29e71.
//
// Solidity: function agentEnabled() constant returns(bool)
func (_Agent *AgentSession) AgentEnabled() (bool, error) {
	return _Agent.Contract.AgentEnabled(&_Agent.CallOpts)
}

// AgentEnabled is a free data retrieval call binding the contract method 0x99d29e71.
//
// Solidity: function agentEnabled() constant returns(bool)
func (_Agent *AgentCallerSession) AgentEnabled() (bool, error) {
	return _Agent.Contract.AgentEnabled(&_Agent.CallOpts)
}

// CustodianEnabled is a free data retrieval call binding the contract method 0x108fd0cc.
//
// Solidity: function custodianEnabled( uint256) constant returns(bool)
func (_Agent *AgentCaller) CustodianEnabled(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "custodianEnabled", arg0)
	return *ret0, err
}

// CustodianEnabled is a free data retrieval call binding the contract method 0x108fd0cc.
//
// Solidity: function custodianEnabled( uint256) constant returns(bool)
func (_Agent *AgentSession) CustodianEnabled(arg0 *big.Int) (bool, error) {
	return _Agent.Contract.CustodianEnabled(&_Agent.CallOpts, arg0)
}

// CustodianEnabled is a free data retrieval call binding the contract method 0x108fd0cc.
//
// Solidity: function custodianEnabled( uint256) constant returns(bool)
func (_Agent *AgentCallerSession) CustodianEnabled(arg0 *big.Int) (bool, error) {
	return _Agent.Contract.CustodianEnabled(&_Agent.CallOpts, arg0)
}

// Custodians is a free data retrieval call binding the contract method 0x2299bda6.
//
// Solidity: function custodians( uint256) constant returns(address)
func (_Agent *AgentCaller) Custodians(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "custodians", arg0)
	return *ret0, err
}

// Custodians is a free data retrieval call binding the contract method 0x2299bda6.
//
// Solidity: function custodians( uint256) constant returns(address)
func (_Agent *AgentSession) Custodians(arg0 *big.Int) (common.Address, error) {
	return _Agent.Contract.Custodians(&_Agent.CallOpts, arg0)
}

// Custodians is a free data retrieval call binding the contract method 0x2299bda6.
//
// Solidity: function custodians( uint256) constant returns(address)
func (_Agent *AgentCallerSession) Custodians(arg0 *big.Int) (common.Address, error) {
	return _Agent.Contract.Custodians(&_Agent.CallOpts, arg0)
}

// GetNumCustodians is a free data retrieval call binding the contract method 0xfaea7b2c.
//
// Solidity: function getNumCustodians() constant returns(uint256)
func (_Agent *AgentCaller) GetNumCustodians(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "getNumCustodians")
	return *ret0, err
}

// GetNumCustodians is a free data retrieval call binding the contract method 0xfaea7b2c.
//
// Solidity: function getNumCustodians() constant returns(uint256)
func (_Agent *AgentSession) GetNumCustodians() (*big.Int, error) {
	return _Agent.Contract.GetNumCustodians(&_Agent.CallOpts)
}

// GetNumCustodians is a free data retrieval call binding the contract method 0xfaea7b2c.
//
// Solidity: function getNumCustodians() constant returns(uint256)
func (_Agent *AgentCallerSession) GetNumCustodians() (*big.Int, error) {
	return _Agent.Contract.GetNumCustodians(&_Agent.CallOpts)
}

// GetNumEnabledOwners is a free data retrieval call binding the contract method 0xc23a0d28.
//
// Solidity: function getNumEnabledOwners() constant returns(uint256)
func (_Agent *AgentCaller) GetNumEnabledOwners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "getNumEnabledOwners")
	return *ret0, err
}

// GetNumEnabledOwners is a free data retrieval call binding the contract method 0xc23a0d28.
//
// Solidity: function getNumEnabledOwners() constant returns(uint256)
func (_Agent *AgentSession) GetNumEnabledOwners() (*big.Int, error) {
	return _Agent.Contract.GetNumEnabledOwners(&_Agent.CallOpts)
}

// GetNumEnabledOwners is a free data retrieval call binding the contract method 0xc23a0d28.
//
// Solidity: function getNumEnabledOwners() constant returns(uint256)
func (_Agent *AgentCallerSession) GetNumEnabledOwners() (*big.Int, error) {
	return _Agent.Contract.GetNumEnabledOwners(&_Agent.CallOpts)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_Agent *AgentCaller) GetNumPermissions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "getNumPermissions")
	return *ret0, err
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_Agent *AgentSession) GetNumPermissions() (*big.Int, error) {
	return _Agent.Contract.GetNumPermissions(&_Agent.CallOpts)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_Agent *AgentCallerSession) GetNumPermissions() (*big.Int, error) {
	return _Agent.Contract.GetNumPermissions(&_Agent.CallOpts)
}

// GetNumRelationships is a free data retrieval call binding the contract method 0xe968177e.
//
// Solidity: function getNumRelationships() constant returns(uint256)
func (_Agent *AgentCaller) GetNumRelationships(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "getNumRelationships")
	return *ret0, err
}

// GetNumRelationships is a free data retrieval call binding the contract method 0xe968177e.
//
// Solidity: function getNumRelationships() constant returns(uint256)
func (_Agent *AgentSession) GetNumRelationships() (*big.Int, error) {
	return _Agent.Contract.GetNumRelationships(&_Agent.CallOpts)
}

// GetNumRelationships is a free data retrieval call binding the contract method 0xe968177e.
//
// Solidity: function getNumRelationships() constant returns(uint256)
func (_Agent *AgentCallerSession) GetNumRelationships() (*big.Int, error) {
	return _Agent.Contract.GetNumRelationships(&_Agent.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_Agent *AgentCaller) Permissions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "permissions", arg0)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_Agent *AgentSession) Permissions(arg0 *big.Int) ([32]byte, error) {
	return _Agent.Contract.Permissions(&_Agent.CallOpts, arg0)
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_Agent *AgentCallerSession) Permissions(arg0 *big.Int) ([32]byte, error) {
	return _Agent.Contract.Permissions(&_Agent.CallOpts, arg0)
}

// Relationships is a free data retrieval call binding the contract method 0x94aa2718.
//
// Solidity: function relationships( uint256) constant returns(address)
func (_Agent *AgentCaller) Relationships(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Agent.contract.Call(opts, out, "relationships", arg0)
	return *ret0, err
}

// Relationships is a free data retrieval call binding the contract method 0x94aa2718.
//
// Solidity: function relationships( uint256) constant returns(address)
func (_Agent *AgentSession) Relationships(arg0 *big.Int) (common.Address, error) {
	return _Agent.Contract.Relationships(&_Agent.CallOpts, arg0)
}

// Relationships is a free data retrieval call binding the contract method 0x94aa2718.
//
// Solidity: function relationships( uint256) constant returns(address)
func (_Agent *AgentCallerSession) Relationships(arg0 *big.Int) (common.Address, error) {
	return _Agent.Contract.Relationships(&_Agent.CallOpts, arg0)
}

// AddCustodian is a paid mutator transaction binding the contract method 0x85d95b81.
//
// Solidity: function addCustodian(addr address) returns()
func (_Agent *AgentTransactor) AddCustodian(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "addCustodian", addr)
}

// AddCustodian is a paid mutator transaction binding the contract method 0x85d95b81.
//
// Solidity: function addCustodian(addr address) returns()
func (_Agent *AgentSession) AddCustodian(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.AddCustodian(&_Agent.TransactOpts, addr)
}

// AddCustodian is a paid mutator transaction binding the contract method 0x85d95b81.
//
// Solidity: function addCustodian(addr address) returns()
func (_Agent *AgentTransactorSession) AddCustodian(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.AddCustodian(&_Agent.TransactOpts, addr)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_Agent *AgentTransactor) AddPermission(opts *bind.TransactOpts, p [32]byte) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "addPermission", p)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_Agent *AgentSession) AddPermission(p [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.AddPermission(&_Agent.TransactOpts, p)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_Agent *AgentTransactorSession) AddPermission(p [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.AddPermission(&_Agent.TransactOpts, p)
}

// AddRelationship is a paid mutator transaction binding the contract method 0x79bda937.
//
// Solidity: function addRelationship(r address) returns()
func (_Agent *AgentTransactor) AddRelationship(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "addRelationship", r)
}

// AddRelationship is a paid mutator transaction binding the contract method 0x79bda937.
//
// Solidity: function addRelationship(r address) returns()
func (_Agent *AgentSession) AddRelationship(r common.Address) (*types.Transaction, error) {
	return _Agent.Contract.AddRelationship(&_Agent.TransactOpts, r)
}

// AddRelationship is a paid mutator transaction binding the contract method 0x79bda937.
//
// Solidity: function addRelationship(r address) returns()
func (_Agent *AgentTransactorSession) AddRelationship(r common.Address) (*types.Transaction, error) {
	return _Agent.Contract.AddRelationship(&_Agent.TransactOpts, r)
}

// EnableAgent is a paid mutator transaction binding the contract method 0x2788aa71.
//
// Solidity: function enableAgent(enable bool) returns()
func (_Agent *AgentTransactor) EnableAgent(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "enableAgent", enable)
}

// EnableAgent is a paid mutator transaction binding the contract method 0x2788aa71.
//
// Solidity: function enableAgent(enable bool) returns()
func (_Agent *AgentSession) EnableAgent(enable bool) (*types.Transaction, error) {
	return _Agent.Contract.EnableAgent(&_Agent.TransactOpts, enable)
}

// EnableAgent is a paid mutator transaction binding the contract method 0x2788aa71.
//
// Solidity: function enableAgent(enable bool) returns()
func (_Agent *AgentTransactorSession) EnableAgent(enable bool) (*types.Transaction, error) {
	return _Agent.Contract.EnableAgent(&_Agent.TransactOpts, enable)
}

// EnableCustodian is a paid mutator transaction binding the contract method 0x6ca9677a.
//
// Solidity: function enableCustodian(addr address, enable bool) returns()
func (_Agent *AgentTransactor) EnableCustodian(opts *bind.TransactOpts, addr common.Address, enable bool) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "enableCustodian", addr, enable)
}

// EnableCustodian is a paid mutator transaction binding the contract method 0x6ca9677a.
//
// Solidity: function enableCustodian(addr address, enable bool) returns()
func (_Agent *AgentSession) EnableCustodian(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Agent.Contract.EnableCustodian(&_Agent.TransactOpts, addr, enable)
}

// EnableCustodian is a paid mutator transaction binding the contract method 0x6ca9677a.
//
// Solidity: function enableCustodian(addr address, enable bool) returns()
func (_Agent *AgentTransactorSession) EnableCustodian(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Agent.Contract.EnableCustodian(&_Agent.TransactOpts, addr, enable)
}

// RemoveCustodian is a paid mutator transaction binding the contract method 0x0f12a66a.
//
// Solidity: function removeCustodian(addr address) returns()
func (_Agent *AgentTransactor) RemoveCustodian(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "removeCustodian", addr)
}

// RemoveCustodian is a paid mutator transaction binding the contract method 0x0f12a66a.
//
// Solidity: function removeCustodian(addr address) returns()
func (_Agent *AgentSession) RemoveCustodian(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.RemoveCustodian(&_Agent.TransactOpts, addr)
}

// RemoveCustodian is a paid mutator transaction binding the contract method 0x0f12a66a.
//
// Solidity: function removeCustodian(addr address) returns()
func (_Agent *AgentTransactorSession) RemoveCustodian(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.RemoveCustodian(&_Agent.TransactOpts, addr)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(addr address) returns()
func (_Agent *AgentTransactor) SetAgent(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setAgent", addr)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(addr address) returns()
func (_Agent *AgentSession) SetAgent(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAgent(&_Agent.TransactOpts, addr)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(addr address) returns()
func (_Agent *AgentTransactorSession) SetAgent(addr common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAgent(&_Agent.TransactOpts, addr)
}

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

// AgentGroupABI is the input ABI used to generate the binding from.
const AgentGroupABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"agents\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumPermissions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumAgents\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"p\",\"type\":\"bytes32\"}],\"name\":\"addPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AgentGroupBin is the compiled bytecode used for deploying new contracts.
const AgentGroupBin = `0x6060604052341561000f57600080fd5b6000805460018101610021838261004b565b5060009182526020909120018054600160a060020a03191633600160a060020a0316179055610095565b81548183558181151161006f5760008381526020902061006f918101908301610074565b505050565b61009291905b8082111561008e576000815560010161007a565b5090565b90565b6104d2806100a46000396000f3006060604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663513856c8811461008757806358886dba146100b95780636121785c146100e157806384e79842146100f457806397a6278e14610115578063dc79264614610134578063fd17ec5614610147575b600080fd5b341561009257600080fd5b61009d60043561015d565b604051600160a060020a03909116815260200160405180910390f35b34156100c457600080fd5b6100cf600435610185565b60405190815260200160405180910390f35b34156100ec57600080fd5b6100cf6101a4565b34156100ff57600080fd5b610113600160a060020a03600435166101ab565b005b341561012057600080fd5b610113600160a060020a0360043516610252565b341561013f57600080fd5b6100cf6103db565b341561015257600080fd5b6101136004356103e1565b600080548290811061016b57fe5b600091825260209091200154600160a060020a0316905081565b600180548290811061019357fe5b600091825260209091200154905081565b6001545b90565b6000805b6000548110156101f85760008054829081106101c757fe5b60009182526020909120015433600160a060020a03908116911614156101f057600191506101f8565b6001016101af565b81151561020457600080fd5b6000805460018101610216838261045f565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0394909416939093179092555050565b60008080805b6000548110156102a157600080548290811061027057fe5b60009182526020909120015433600160a060020a039081169116141561029957600191506102a1565b600101610258565b8115156102ad57600080fd5b60009350600092505b6000548310156103775783156103335760008054849081106102d457fe5b60009182526020822001548154600160a060020a03909116919060001986019081106102fc57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555b84600160a060020a031660008481548110151561034c57fe5b600091825260209091200154600160a060020a0316141561036c57600193505b6001909201916102b6565b60008054600019810190811061038957fe5b60009182526020822001805473ffffffffffffffffffffffffffffffffffffffff19169055805460001901906103bf908261045f565b5060006103ca6103db565b116103d457600080fd5b5050505050565b60005490565b6000805b60005481101561042e5760008054829081106103fd57fe5b60009182526020909120015433600160a060020a0390811691161415610426576001915061042e565b6001016103e5565b81151561043a57600080fd5b6001805480820161044b838261045f565b506000918252602090912001929092555050565b81548183558181151161048357600083815260209020610483918101908301610488565b505050565b6101a891905b808211156104a2576000815560010161048e565b50905600a165627a7a7230582055fce46fd259e1058997bf05cd0f0dbdc6ff2ab87b045ca9a715019c98ab6fc20029`

// DeployAgentGroup deploys a new Ethereum contract, binding an instance of AgentGroup to it.
func DeployAgentGroup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgentGroup, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentGroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AgentGroupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentGroup{AgentGroupCaller: AgentGroupCaller{contract: contract}, AgentGroupTransactor: AgentGroupTransactor{contract: contract}}, nil
}

// AgentGroup is an auto generated Go binding around an Ethereum contract.
type AgentGroup struct {
	AgentGroupCaller     // Read-only binding to the contract
	AgentGroupTransactor // Write-only binding to the contract
}

// AgentGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentGroupSession struct {
	Contract     *AgentGroup       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentGroupCallerSession struct {
	Contract *AgentGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AgentGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentGroupTransactorSession struct {
	Contract     *AgentGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AgentGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentGroupRaw struct {
	Contract *AgentGroup // Generic contract binding to access the raw methods on
}

// AgentGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentGroupCallerRaw struct {
	Contract *AgentGroupCaller // Generic read-only contract binding to access the raw methods on
}

// AgentGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentGroupTransactorRaw struct {
	Contract *AgentGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentGroup creates a new instance of AgentGroup, bound to a specific deployed contract.
func NewAgentGroup(address common.Address, backend bind.ContractBackend) (*AgentGroup, error) {
	contract, err := bindAgentGroup(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentGroup{AgentGroupCaller: AgentGroupCaller{contract: contract}, AgentGroupTransactor: AgentGroupTransactor{contract: contract}}, nil
}

// NewAgentGroupCaller creates a new read-only instance of AgentGroup, bound to a specific deployed contract.
func NewAgentGroupCaller(address common.Address, caller bind.ContractCaller) (*AgentGroupCaller, error) {
	contract, err := bindAgentGroup(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AgentGroupCaller{contract: contract}, nil
}

// NewAgentGroupTransactor creates a new write-only instance of AgentGroup, bound to a specific deployed contract.
func NewAgentGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentGroupTransactor, error) {
	contract, err := bindAgentGroup(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AgentGroupTransactor{contract: contract}, nil
}

// bindAgentGroup binds a generic wrapper to an already deployed contract.
func bindAgentGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentGroup *AgentGroupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AgentGroup.Contract.AgentGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentGroup *AgentGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentGroup.Contract.AgentGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentGroup *AgentGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentGroup.Contract.AgentGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentGroup *AgentGroupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AgentGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentGroup *AgentGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentGroup *AgentGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentGroup.Contract.contract.Transact(opts, method, params...)
}

// Agents is a free data retrieval call binding the contract method 0x513856c8.
//
// Solidity: function agents( uint256) constant returns(address)
func (_AgentGroup *AgentGroupCaller) Agents(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AgentGroup.contract.Call(opts, out, "agents", arg0)
	return *ret0, err
}

// Agents is a free data retrieval call binding the contract method 0x513856c8.
//
// Solidity: function agents( uint256) constant returns(address)
func (_AgentGroup *AgentGroupSession) Agents(arg0 *big.Int) (common.Address, error) {
	return _AgentGroup.Contract.Agents(&_AgentGroup.CallOpts, arg0)
}

// Agents is a free data retrieval call binding the contract method 0x513856c8.
//
// Solidity: function agents( uint256) constant returns(address)
func (_AgentGroup *AgentGroupCallerSession) Agents(arg0 *big.Int) (common.Address, error) {
	return _AgentGroup.Contract.Agents(&_AgentGroup.CallOpts, arg0)
}

// GetNumAgents is a free data retrieval call binding the contract method 0xdc792646.
//
// Solidity: function getNumAgents() constant returns(uint256)
func (_AgentGroup *AgentGroupCaller) GetNumAgents(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AgentGroup.contract.Call(opts, out, "getNumAgents")
	return *ret0, err
}

// GetNumAgents is a free data retrieval call binding the contract method 0xdc792646.
//
// Solidity: function getNumAgents() constant returns(uint256)
func (_AgentGroup *AgentGroupSession) GetNumAgents() (*big.Int, error) {
	return _AgentGroup.Contract.GetNumAgents(&_AgentGroup.CallOpts)
}

// GetNumAgents is a free data retrieval call binding the contract method 0xdc792646.
//
// Solidity: function getNumAgents() constant returns(uint256)
func (_AgentGroup *AgentGroupCallerSession) GetNumAgents() (*big.Int, error) {
	return _AgentGroup.Contract.GetNumAgents(&_AgentGroup.CallOpts)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_AgentGroup *AgentGroupCaller) GetNumPermissions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AgentGroup.contract.Call(opts, out, "getNumPermissions")
	return *ret0, err
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_AgentGroup *AgentGroupSession) GetNumPermissions() (*big.Int, error) {
	return _AgentGroup.Contract.GetNumPermissions(&_AgentGroup.CallOpts)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x6121785c.
//
// Solidity: function getNumPermissions() constant returns(uint256)
func (_AgentGroup *AgentGroupCallerSession) GetNumPermissions() (*big.Int, error) {
	return _AgentGroup.Contract.GetNumPermissions(&_AgentGroup.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_AgentGroup *AgentGroupCaller) Permissions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AgentGroup.contract.Call(opts, out, "permissions", arg0)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_AgentGroup *AgentGroupSession) Permissions(arg0 *big.Int) ([32]byte, error) {
	return _AgentGroup.Contract.Permissions(&_AgentGroup.CallOpts, arg0)
}

// Permissions is a free data retrieval call binding the contract method 0x58886dba.
//
// Solidity: function permissions( uint256) constant returns(bytes32)
func (_AgentGroup *AgentGroupCallerSession) Permissions(arg0 *big.Int) ([32]byte, error) {
	return _AgentGroup.Contract.Permissions(&_AgentGroup.CallOpts, arg0)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(addr address) returns()
func (_AgentGroup *AgentGroupTransactor) AddAgent(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.contract.Transact(opts, "addAgent", addr)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(addr address) returns()
func (_AgentGroup *AgentGroupSession) AddAgent(addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.Contract.AddAgent(&_AgentGroup.TransactOpts, addr)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(addr address) returns()
func (_AgentGroup *AgentGroupTransactorSession) AddAgent(addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.Contract.AddAgent(&_AgentGroup.TransactOpts, addr)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_AgentGroup *AgentGroupTransactor) AddPermission(opts *bind.TransactOpts, p [32]byte) (*types.Transaction, error) {
	return _AgentGroup.contract.Transact(opts, "addPermission", p)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_AgentGroup *AgentGroupSession) AddPermission(p [32]byte) (*types.Transaction, error) {
	return _AgentGroup.Contract.AddPermission(&_AgentGroup.TransactOpts, p)
}

// AddPermission is a paid mutator transaction binding the contract method 0xfd17ec56.
//
// Solidity: function addPermission(p bytes32) returns()
func (_AgentGroup *AgentGroupTransactorSession) AddPermission(p [32]byte) (*types.Transaction, error) {
	return _AgentGroup.Contract.AddPermission(&_AgentGroup.TransactOpts, p)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(addr address) returns()
func (_AgentGroup *AgentGroupTransactor) RemoveAgent(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.contract.Transact(opts, "removeAgent", addr)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(addr address) returns()
func (_AgentGroup *AgentGroupSession) RemoveAgent(addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.Contract.RemoveAgent(&_AgentGroup.TransactOpts, addr)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(addr address) returns()
func (_AgentGroup *AgentGroupTransactorSession) RemoveAgent(addr common.Address) (*types.Transaction, error) {
	return _AgentGroup.Contract.RemoveAgent(&_AgentGroup.TransactOpts, addr)
}

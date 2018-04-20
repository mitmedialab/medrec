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

// RelationshipABI is the input ABI used to generate the binding from.
const RelationshipABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"permissionInfo\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"canRead\",\"type\":\"bool\"},{\"name\":\"canWrite\",\"type\":\"bool\"},{\"name\":\"startDay\",\"type\":\"int256\"},{\"name\":\"durationDays\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"provider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"patron\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"removePermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"canWrite\",\"type\":\"bool\"}],\"name\":\"updatePermissionWriteAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"canRead\",\"type\":\"bool\"},{\"name\":\"canWrite\",\"type\":\"bool\"},{\"name\":\"duration\",\"type\":\"int256\"}],\"name\":\"addPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"addViewer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"getNumPermissions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getViewers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"canRead\",\"type\":\"bool\"}],\"name\":\"updatePermissionReadAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"duration\",\"type\":\"int256\"}],\"name\":\"updatePermissionDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"viewersList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"removeViewer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"viewer\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"checkPermission\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_provider\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// RelationshipBin is the compiled bytecode used for deploying new contracts.
const RelationshipBin = `0x6060604052341561000f57600080fd5b604051602080610d2b8339810160405280805160008054600160a060020a03338116600160a060020a031992831617909255600180549290931691161790555050610ccc8061005f6000396000f3006060604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663069e820b81146100d4578063085d48831461012d5780630ba32b271461015c578063186ddd851461016f57806333c85fe3146101935780635fa164a5146101ba578063601f8cf5146101e95780637c85e18b146102085780638dd137d9146102395780639090a0a71461029f5780639367a7b3146102c6578063b4d66407146102eb578063d1d9891414610301578063d639205714610320575b600080fd5b34156100df57600080fd5b6100f6600160a060020a036004351660243561035e565b60405194855292151560208501529015156040808501919091526060840191909152608083019190915260a0909101905180910390f35b341561013857600080fd5b61014061039b565b604051600160a060020a03909116815260200160405180910390f35b341561016757600080fd5b6101406103aa565b341561017a57600080fd5b610191600160a060020a03600435166024356103b9565b005b341561019e57600080fd5b610191600160a060020a03600435166024356044351515610576565b34156101c557600080fd5b610191600160a060020a0360043516602435604435151560643515156084356105f6565b34156101f457600080fd5b610191600160a060020a0360043516610717565b341561021357600080fd5b610227600160a060020a03600435166107c9565b60405190815260200160405180910390f35b341561024457600080fd5b61024c6107e4565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561028b578082015183820152602001610273565b505050509050019250505060405180910390f35b34156102aa57600080fd5b610191600160a060020a0360043516602435604435151561084d565b34156102d157600080fd5b610191600160a060020a03600435166024356044356108c7565b34156102f657600080fd5b610140600435610932565b341561030c57600080fd5b610191600160a060020a036004351661095a565b341561032b57600080fd5b610342600160a060020a0360043516602435610b8d565b6040519115158252151560208201526040908101905180910390f35b600260208181526000938452604080852090915291835291208054600182015492820154600390920154909260ff80821693610100909204169185565b600154600160a060020a031681565b600054600160a060020a031681565b60008054819033600160a060020a039081169116146103d757600080fd5b600160a060020a03841660009081526005602052604090205460ff1615156103fe57600080fd5b5050600160a060020a0382166000908152600260208181526040808420858552909152822082815560018101805461ffff19169055908101829055600301819055805b600160a060020a0384166000908152600360205260409020548110156105125781156104cc57600160a060020a038416600090815260036020526040902080548290811061048b57fe5b6000918252602080832090910154600160a060020a038716835260039091526040909120805460001984019081106104bf57fe5b6000918252602090912001555b600160a060020a03841660009081526003602052604090208054849190839081106104f357fe5b600091825260209091200154141561050a57600191505b600101610441565b600160a060020a03841660009081526003602052604090208054600019810190811061053a57fe5b60009182526020808320909101829055600160a060020a038616825260039052604090208054600019019061056f9082610c2b565b5050505050565b60005433600160a060020a0390811691161461059157600080fd5b600160a060020a03831660009081526005602052604090205460ff1615156105b857600080fd5b600160a060020a0390921660009081526002602090815260408083209383529290522060010180549115156101000261ff0019909216919091179055565b60005433600160a060020a0390811691161461061157600080fd5b600160a060020a03851660009081526005602052604090205460ff16151561063857600080fd5b60a060405190810160409081528582528415156020808401919091528415158284015242606084015260808301849052600160a060020a0388166000908152600282528281208882529091522081518155602082015160018201805460ff191691151591909117905560408201516001820180549115156101000261ff0019909216919091179055606082015181600201556080820151600391820155600160a060020a0387166000908152602091909152604090208054909150600181016107018382610c2b565b5060009182526020909120019390935550505050565b60005433600160a060020a0390811691161461073257600080fd5b600160a060020a03811660009081526005602052604090205460ff161561075857600080fd5b600160a060020a0381166000908152600560205260409020805460ff1916600190811790915560048054909181016107908382610c2b565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a031660009081526003602052604090205490565b6107ec610c4f565b600480548060200260200160405190810160405280929190818152602001828054801561084257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610824575b505050505090505b90565b60005433600160a060020a0390811691161461086857600080fd5b600160a060020a03831660009081526005602052604090205460ff16151561088f57600080fd5b600160a060020a0392909216600090815260026020908152604080832093835292905220600101805460ff1916911515919091179055565b60005433600160a060020a039081169116146108e257600080fd5b600160a060020a03831660009081526005602052604090205460ff16151561090957600080fd5b600160a060020a0390921660009081526002602090815260408083209383529290522060030155565b600480548290811061094057fe5b600091825260209091200154600160a060020a0316905081565b60008054819033600160a060020a0390811691161461097857600080fd5b600160a060020a03831660009081526005602052604090205460ff16151561099f57600080fd5b5050600160a060020a0381166000908152600560205260408120805460ff19169055805b600454811015610a83578115610a425760048054829081106109e157fe5b60009182526020909120015460048054600160a060020a03909216916000198401908110610a0b57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555b82600160a060020a0316600482815481101515610a5b57fe5b600091825260209091200154600160a060020a03161415610a7b57600191505b6001016109c3565b600480546000198101908110610a9557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19169055600480546000190190610ace9082610c2b565b50600090505b600160a060020a038316600090815260036020526040902054811015610b6757600160a060020a038316600090815260026020908152604080832060039092528220805491929184908110610b2557fe5b6000918252602080832090910154835282019290925260400181208181556001808201805461ffff191690556002820183905560039091019190915501610ad4565b600160a060020a0383166000908152600360205260408120610b8891610c61565b505050565b600160a060020a038216600090815260026020818152604080842085855290915282206003810154918101548392908101901515610bd15760008093509350610c22565b600082600301541215610bf857600182015460ff8082169550610100909104169250610c22565b42819013610c1a57600182015460ff8082169550610100909104169250610c22565b600080935093505b50509250929050565b815481835581811511610b8857600083815260209020610b88918101908301610c82565b60206040519081016040526000815290565b5080546000825590600052602060002090810190610c7f9190610c82565b50565b61084a91905b80821115610c9c5760008155600101610c88565b50905600a165627a7a723058203a76ca6d08e4c95c4c0039e82eaadc77a7f0df9b5380e6786c44840c0a705fa40029`

// DeployRelationship deploys a new Ethereum contract, binding an instance of Relationship to it.
func DeployRelationship(auth *bind.TransactOpts, backend bind.ContractBackend, _provider common.Address) (common.Address, *types.Transaction, *Relationship, error) {
	parsed, err := abi.JSON(strings.NewReader(RelationshipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelationshipBin), backend, _provider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Relationship{RelationshipCaller: RelationshipCaller{contract: contract}, RelationshipTransactor: RelationshipTransactor{contract: contract}}, nil
}

// Relationship is an auto generated Go binding around an Ethereum contract.
type Relationship struct {
	RelationshipCaller     // Read-only binding to the contract
	RelationshipTransactor // Write-only binding to the contract
}

// RelationshipCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelationshipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelationshipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelationshipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelationshipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelationshipSession struct {
	Contract     *Relationship     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelationshipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelationshipCallerSession struct {
	Contract *RelationshipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RelationshipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelationshipTransactorSession struct {
	Contract     *RelationshipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RelationshipRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelationshipRaw struct {
	Contract *Relationship // Generic contract binding to access the raw methods on
}

// RelationshipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelationshipCallerRaw struct {
	Contract *RelationshipCaller // Generic read-only contract binding to access the raw methods on
}

// RelationshipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelationshipTransactorRaw struct {
	Contract *RelationshipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelationship creates a new instance of Relationship, bound to a specific deployed contract.
func NewRelationship(address common.Address, backend bind.ContractBackend) (*Relationship, error) {
	contract, err := bindRelationship(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relationship{RelationshipCaller: RelationshipCaller{contract: contract}, RelationshipTransactor: RelationshipTransactor{contract: contract}}, nil
}

// NewRelationshipCaller creates a new read-only instance of Relationship, bound to a specific deployed contract.
func NewRelationshipCaller(address common.Address, caller bind.ContractCaller) (*RelationshipCaller, error) {
	contract, err := bindRelationship(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RelationshipCaller{contract: contract}, nil
}

// NewRelationshipTransactor creates a new write-only instance of Relationship, bound to a specific deployed contract.
func NewRelationshipTransactor(address common.Address, transactor bind.ContractTransactor) (*RelationshipTransactor, error) {
	contract, err := bindRelationship(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RelationshipTransactor{contract: contract}, nil
}

// bindRelationship binds a generic wrapper to an already deployed contract.
func bindRelationship(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelationshipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relationship *RelationshipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relationship.Contract.RelationshipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relationship *RelationshipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relationship *RelationshipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relationship *RelationshipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relationship.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relationship *RelationshipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relationship.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relationship *RelationshipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relationship.Contract.contract.Transact(opts, method, params...)
}

// CheckPermission is a free data retrieval call binding the contract method 0xd6392057.
//
// Solidity: function checkPermission(viewer address, name bytes32) constant returns(bool, bool)
func (_Relationship *RelationshipCaller) CheckPermission(opts *bind.CallOpts, viewer common.Address, name [32]byte) (bool, bool, error) {
	var (
		ret0 = new(bool)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Relationship.contract.Call(opts, out, "checkPermission", viewer, name)
	return *ret0, *ret1, err
}

// CheckPermission is a free data retrieval call binding the contract method 0xd6392057.
//
// Solidity: function checkPermission(viewer address, name bytes32) constant returns(bool, bool)
func (_Relationship *RelationshipSession) CheckPermission(viewer common.Address, name [32]byte) (bool, bool, error) {
	return _Relationship.Contract.CheckPermission(&_Relationship.CallOpts, viewer, name)
}

// CheckPermission is a free data retrieval call binding the contract method 0xd6392057.
//
// Solidity: function checkPermission(viewer address, name bytes32) constant returns(bool, bool)
func (_Relationship *RelationshipCallerSession) CheckPermission(viewer common.Address, name [32]byte) (bool, bool, error) {
	return _Relationship.Contract.CheckPermission(&_Relationship.CallOpts, viewer, name)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x7c85e18b.
//
// Solidity: function getNumPermissions(viewer address) constant returns(uint256)
func (_Relationship *RelationshipCaller) GetNumPermissions(opts *bind.CallOpts, viewer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relationship.contract.Call(opts, out, "getNumPermissions", viewer)
	return *ret0, err
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x7c85e18b.
//
// Solidity: function getNumPermissions(viewer address) constant returns(uint256)
func (_Relationship *RelationshipSession) GetNumPermissions(viewer common.Address) (*big.Int, error) {
	return _Relationship.Contract.GetNumPermissions(&_Relationship.CallOpts, viewer)
}

// GetNumPermissions is a free data retrieval call binding the contract method 0x7c85e18b.
//
// Solidity: function getNumPermissions(viewer address) constant returns(uint256)
func (_Relationship *RelationshipCallerSession) GetNumPermissions(viewer common.Address) (*big.Int, error) {
	return _Relationship.Contract.GetNumPermissions(&_Relationship.CallOpts, viewer)
}

// GetViewers is a free data retrieval call binding the contract method 0x8dd137d9.
//
// Solidity: function getViewers() constant returns(address[])
func (_Relationship *RelationshipCaller) GetViewers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Relationship.contract.Call(opts, out, "getViewers")
	return *ret0, err
}

// GetViewers is a free data retrieval call binding the contract method 0x8dd137d9.
//
// Solidity: function getViewers() constant returns(address[])
func (_Relationship *RelationshipSession) GetViewers() ([]common.Address, error) {
	return _Relationship.Contract.GetViewers(&_Relationship.CallOpts)
}

// GetViewers is a free data retrieval call binding the contract method 0x8dd137d9.
//
// Solidity: function getViewers() constant returns(address[])
func (_Relationship *RelationshipCallerSession) GetViewers() ([]common.Address, error) {
	return _Relationship.Contract.GetViewers(&_Relationship.CallOpts)
}

// Patron is a free data retrieval call binding the contract method 0x0ba32b27.
//
// Solidity: function patron() constant returns(address)
func (_Relationship *RelationshipCaller) Patron(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relationship.contract.Call(opts, out, "patron")
	return *ret0, err
}

// Patron is a free data retrieval call binding the contract method 0x0ba32b27.
//
// Solidity: function patron() constant returns(address)
func (_Relationship *RelationshipSession) Patron() (common.Address, error) {
	return _Relationship.Contract.Patron(&_Relationship.CallOpts)
}

// Patron is a free data retrieval call binding the contract method 0x0ba32b27.
//
// Solidity: function patron() constant returns(address)
func (_Relationship *RelationshipCallerSession) Patron() (common.Address, error) {
	return _Relationship.Contract.Patron(&_Relationship.CallOpts)
}

// PermissionInfo is a free data retrieval call binding the contract method 0x069e820b.
//
// Solidity: function permissionInfo( address,  bytes32) constant returns(name bytes32, canRead bool, canWrite bool, startDay int256, durationDays int256)
func (_Relationship *RelationshipCaller) PermissionInfo(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Name         [32]byte
	CanRead      bool
	CanWrite     bool
	StartDay     *big.Int
	DurationDays *big.Int
}, error) {
	ret := new(struct {
		Name         [32]byte
		CanRead      bool
		CanWrite     bool
		StartDay     *big.Int
		DurationDays *big.Int
	})
	out := ret
	err := _Relationship.contract.Call(opts, out, "permissionInfo", arg0, arg1)
	return *ret, err
}

// PermissionInfo is a free data retrieval call binding the contract method 0x069e820b.
//
// Solidity: function permissionInfo( address,  bytes32) constant returns(name bytes32, canRead bool, canWrite bool, startDay int256, durationDays int256)
func (_Relationship *RelationshipSession) PermissionInfo(arg0 common.Address, arg1 [32]byte) (struct {
	Name         [32]byte
	CanRead      bool
	CanWrite     bool
	StartDay     *big.Int
	DurationDays *big.Int
}, error) {
	return _Relationship.Contract.PermissionInfo(&_Relationship.CallOpts, arg0, arg1)
}

// PermissionInfo is a free data retrieval call binding the contract method 0x069e820b.
//
// Solidity: function permissionInfo( address,  bytes32) constant returns(name bytes32, canRead bool, canWrite bool, startDay int256, durationDays int256)
func (_Relationship *RelationshipCallerSession) PermissionInfo(arg0 common.Address, arg1 [32]byte) (struct {
	Name         [32]byte
	CanRead      bool
	CanWrite     bool
	StartDay     *big.Int
	DurationDays *big.Int
}, error) {
	return _Relationship.Contract.PermissionInfo(&_Relationship.CallOpts, arg0, arg1)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() constant returns(address)
func (_Relationship *RelationshipCaller) Provider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relationship.contract.Call(opts, out, "provider")
	return *ret0, err
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() constant returns(address)
func (_Relationship *RelationshipSession) Provider() (common.Address, error) {
	return _Relationship.Contract.Provider(&_Relationship.CallOpts)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() constant returns(address)
func (_Relationship *RelationshipCallerSession) Provider() (common.Address, error) {
	return _Relationship.Contract.Provider(&_Relationship.CallOpts)
}

// ViewersList is a free data retrieval call binding the contract method 0xb4d66407.
//
// Solidity: function viewersList( uint256) constant returns(address)
func (_Relationship *RelationshipCaller) ViewersList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relationship.contract.Call(opts, out, "viewersList", arg0)
	return *ret0, err
}

// ViewersList is a free data retrieval call binding the contract method 0xb4d66407.
//
// Solidity: function viewersList( uint256) constant returns(address)
func (_Relationship *RelationshipSession) ViewersList(arg0 *big.Int) (common.Address, error) {
	return _Relationship.Contract.ViewersList(&_Relationship.CallOpts, arg0)
}

// ViewersList is a free data retrieval call binding the contract method 0xb4d66407.
//
// Solidity: function viewersList( uint256) constant returns(address)
func (_Relationship *RelationshipCallerSession) ViewersList(arg0 *big.Int) (common.Address, error) {
	return _Relationship.Contract.ViewersList(&_Relationship.CallOpts, arg0)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5fa164a5.
//
// Solidity: function addPermission(viewer address, name bytes32, canRead bool, canWrite bool, duration int256) returns()
func (_Relationship *RelationshipTransactor) AddPermission(opts *bind.TransactOpts, viewer common.Address, name [32]byte, canRead bool, canWrite bool, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "addPermission", viewer, name, canRead, canWrite, duration)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5fa164a5.
//
// Solidity: function addPermission(viewer address, name bytes32, canRead bool, canWrite bool, duration int256) returns()
func (_Relationship *RelationshipSession) AddPermission(viewer common.Address, name [32]byte, canRead bool, canWrite bool, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.Contract.AddPermission(&_Relationship.TransactOpts, viewer, name, canRead, canWrite, duration)
}

// AddPermission is a paid mutator transaction binding the contract method 0x5fa164a5.
//
// Solidity: function addPermission(viewer address, name bytes32, canRead bool, canWrite bool, duration int256) returns()
func (_Relationship *RelationshipTransactorSession) AddPermission(viewer common.Address, name [32]byte, canRead bool, canWrite bool, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.Contract.AddPermission(&_Relationship.TransactOpts, viewer, name, canRead, canWrite, duration)
}

// AddViewer is a paid mutator transaction binding the contract method 0x601f8cf5.
//
// Solidity: function addViewer(viewer address) returns()
func (_Relationship *RelationshipTransactor) AddViewer(opts *bind.TransactOpts, viewer common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "addViewer", viewer)
}

// AddViewer is a paid mutator transaction binding the contract method 0x601f8cf5.
//
// Solidity: function addViewer(viewer address) returns()
func (_Relationship *RelationshipSession) AddViewer(viewer common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.AddViewer(&_Relationship.TransactOpts, viewer)
}

// AddViewer is a paid mutator transaction binding the contract method 0x601f8cf5.
//
// Solidity: function addViewer(viewer address) returns()
func (_Relationship *RelationshipTransactorSession) AddViewer(viewer common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.AddViewer(&_Relationship.TransactOpts, viewer)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x186ddd85.
//
// Solidity: function removePermission(viewer address, name bytes32) returns()
func (_Relationship *RelationshipTransactor) RemovePermission(opts *bind.TransactOpts, viewer common.Address, name [32]byte) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "removePermission", viewer, name)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x186ddd85.
//
// Solidity: function removePermission(viewer address, name bytes32) returns()
func (_Relationship *RelationshipSession) RemovePermission(viewer common.Address, name [32]byte) (*types.Transaction, error) {
	return _Relationship.Contract.RemovePermission(&_Relationship.TransactOpts, viewer, name)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x186ddd85.
//
// Solidity: function removePermission(viewer address, name bytes32) returns()
func (_Relationship *RelationshipTransactorSession) RemovePermission(viewer common.Address, name [32]byte) (*types.Transaction, error) {
	return _Relationship.Contract.RemovePermission(&_Relationship.TransactOpts, viewer, name)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(viewer address) returns()
func (_Relationship *RelationshipTransactor) RemoveViewer(opts *bind.TransactOpts, viewer common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "removeViewer", viewer)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(viewer address) returns()
func (_Relationship *RelationshipSession) RemoveViewer(viewer common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.RemoveViewer(&_Relationship.TransactOpts, viewer)
}

// RemoveViewer is a paid mutator transaction binding the contract method 0xd1d98914.
//
// Solidity: function removeViewer(viewer address) returns()
func (_Relationship *RelationshipTransactorSession) RemoveViewer(viewer common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.RemoveViewer(&_Relationship.TransactOpts, viewer)
}

// UpdatePermissionDuration is a paid mutator transaction binding the contract method 0x9367a7b3.
//
// Solidity: function updatePermissionDuration(viewer address, name bytes32, duration int256) returns()
func (_Relationship *RelationshipTransactor) UpdatePermissionDuration(opts *bind.TransactOpts, viewer common.Address, name [32]byte, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "updatePermissionDuration", viewer, name, duration)
}

// UpdatePermissionDuration is a paid mutator transaction binding the contract method 0x9367a7b3.
//
// Solidity: function updatePermissionDuration(viewer address, name bytes32, duration int256) returns()
func (_Relationship *RelationshipSession) UpdatePermissionDuration(viewer common.Address, name [32]byte, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionDuration(&_Relationship.TransactOpts, viewer, name, duration)
}

// UpdatePermissionDuration is a paid mutator transaction binding the contract method 0x9367a7b3.
//
// Solidity: function updatePermissionDuration(viewer address, name bytes32, duration int256) returns()
func (_Relationship *RelationshipTransactorSession) UpdatePermissionDuration(viewer common.Address, name [32]byte, duration *big.Int) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionDuration(&_Relationship.TransactOpts, viewer, name, duration)
}

// UpdatePermissionReadAccess is a paid mutator transaction binding the contract method 0x9090a0a7.
//
// Solidity: function updatePermissionReadAccess(viewer address, name bytes32, canRead bool) returns()
func (_Relationship *RelationshipTransactor) UpdatePermissionReadAccess(opts *bind.TransactOpts, viewer common.Address, name [32]byte, canRead bool) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "updatePermissionReadAccess", viewer, name, canRead)
}

// UpdatePermissionReadAccess is a paid mutator transaction binding the contract method 0x9090a0a7.
//
// Solidity: function updatePermissionReadAccess(viewer address, name bytes32, canRead bool) returns()
func (_Relationship *RelationshipSession) UpdatePermissionReadAccess(viewer common.Address, name [32]byte, canRead bool) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionReadAccess(&_Relationship.TransactOpts, viewer, name, canRead)
}

// UpdatePermissionReadAccess is a paid mutator transaction binding the contract method 0x9090a0a7.
//
// Solidity: function updatePermissionReadAccess(viewer address, name bytes32, canRead bool) returns()
func (_Relationship *RelationshipTransactorSession) UpdatePermissionReadAccess(viewer common.Address, name [32]byte, canRead bool) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionReadAccess(&_Relationship.TransactOpts, viewer, name, canRead)
}

// UpdatePermissionWriteAccess is a paid mutator transaction binding the contract method 0x33c85fe3.
//
// Solidity: function updatePermissionWriteAccess(viewer address, name bytes32, canWrite bool) returns()
func (_Relationship *RelationshipTransactor) UpdatePermissionWriteAccess(opts *bind.TransactOpts, viewer common.Address, name [32]byte, canWrite bool) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "updatePermissionWriteAccess", viewer, name, canWrite)
}

// UpdatePermissionWriteAccess is a paid mutator transaction binding the contract method 0x33c85fe3.
//
// Solidity: function updatePermissionWriteAccess(viewer address, name bytes32, canWrite bool) returns()
func (_Relationship *RelationshipSession) UpdatePermissionWriteAccess(viewer common.Address, name [32]byte, canWrite bool) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionWriteAccess(&_Relationship.TransactOpts, viewer, name, canWrite)
}

// UpdatePermissionWriteAccess is a paid mutator transaction binding the contract method 0x33c85fe3.
//
// Solidity: function updatePermissionWriteAccess(viewer address, name bytes32, canWrite bool) returns()
func (_Relationship *RelationshipTransactorSession) UpdatePermissionWriteAccess(viewer common.Address, name [32]byte, canWrite bool) (*types.Transaction, error) {
	return _Relationship.Contract.UpdatePermissionWriteAccess(&_Relationship.TransactOpts, viewer, name, canWrite)
}

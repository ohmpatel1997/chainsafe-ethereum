// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// IPFSABI is the input ABI used to generate the binding from.
const IPFSABI = "[{\"inputs\":[],\"name\":\"printIPFSIdentifiers\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"setIPFSIdentifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPFSFuncSigs maps the 4-byte function signature to its string representation.
var IPFSFuncSigs = map[string]string{
	"62e27a91": "printIPFSIdentifiers()",
	"5e5a404a": "setIPFSIdentifier(string)",
}

// IPFSBin is the compiled bytecode used for deploying new contracts.
var IPFSBin = "0x608060405234801561001057600080fd5b50610527806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80635e5a404a1461003b57806362e27a9114610050575b600080fd5b61004e61004936600461029c565b61006e565b005b61005861012a565b60405161006591906103e1565b60405180910390f35b60008054905b818110156100e3578280519060200120600082815481106100a557634e487b7160e01b600052603260045260246000fd5b906000526020600020016040516100bc9190610346565b604051809103902014156100d1575050610127565b806100db816104b4565b915050610074565b50600080546001810182559080528251610124917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301906020850190610203565b50505b50565b60606000805480602002602001604051908101604052809291908181526020016000905b828210156101fa57838290600052602060002001805461016d90610479565b80601f016020809104026020016040519081016040528092919081815260200182805461019990610479565b80156101e65780601f106101bb576101008083540402835291602001916101e6565b820191906000526020600020905b8154815290600101906020018083116101c957829003601f168201915b50505050508152602001906001019061014e565b50505050905090565b82805461020f90610479565b90600052602060002090601f0160209004810192826102315760008555610277565b82601f1061024a57805160ff1916838001178555610277565b82800160010185558215610277579182015b8281111561027757825182559160200191906001019061025c565b50610283929150610287565b5090565b5b808211156102835760008155600101610288565b6000602082840312156102ad578081fd5b813567ffffffffffffffff808211156102c4578283fd5b818401915084601f8301126102d7578283fd5b8135818111156102e9576102e96104db565b604051601f8201601f19908116603f01168101908382118183101715610311576103116104db565b81604052828152876020848701011115610329578586fd5b826020860160208301379182016020019490945295945050505050565b600080835482600182811c91508083168061036257607f831692505b602080841082141561038257634e487b7160e01b87526022600452602487fd5b81801561039657600181146103a7576103d3565b60ff198616895284890196506103d3565b60008a815260209020885b868110156103cb5781548b8201529085019083016103b2565b505084890196505b509498975050505050505050565b6000602080830181845280855180835260408601915060408160051b8701019250838701855b8281101561046c57878503603f1901845281518051808752885b8181101561043c578281018901518882018a01528801610421565b8181111561044c578989838a0101525b50601f01601f191695909501860194509285019290850190600101610407565b5092979650505050505050565b600181811c9082168061048d57607f821691505b602082108114156104ae57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156104d457634e487b7160e01b81526011600452602481fd5b5060010190565b634e487b7160e01b600052604160045260246000fdfea26469706673582212204b7065b5577d44d5ecbb7e15ed71510baa208ac8ef3271ea6895a1428f453d9664736f6c63430008030033"

// DeployIPFS deploys a new Ethereum contract, binding an instance of IPFS to it.
func DeployIPFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IPFS, error) {
	parsed, err := abi.JSON(strings.NewReader(IPFSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IPFSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IPFS{IPFSCaller: IPFSCaller{contract: contract}, IPFSTransactor: IPFSTransactor{contract: contract}, IPFSFilterer: IPFSFilterer{contract: contract}}, nil
}

// IPFS is an auto generated Go binding around an Ethereum contract.
type IPFS struct {
	IPFSCaller     // Read-only binding to the contract
	IPFSTransactor // Write-only binding to the contract
	IPFSFilterer   // Log filterer for contract events
}

// IPFSCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPFSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPFSSession struct {
	Contract     *IPFS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPFSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPFSCallerSession struct {
	Contract *IPFSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPFSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPFSTransactorSession struct {
	Contract     *IPFSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPFSRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPFSRaw struct {
	Contract *IPFS // Generic contract binding to access the raw methods on
}

// IPFSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPFSCallerRaw struct {
	Contract *IPFSCaller // Generic read-only contract binding to access the raw methods on
}

// IPFSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPFSTransactorRaw struct {
	Contract *IPFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPFS creates a new instance of IPFS, bound to a specific deployed contract.
func NewIPFS(address common.Address, backend bind.ContractBackend) (*IPFS, error) {
	contract, err := bindIPFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPFS{IPFSCaller: IPFSCaller{contract: contract}, IPFSTransactor: IPFSTransactor{contract: contract}, IPFSFilterer: IPFSFilterer{contract: contract}}, nil
}

// NewIPFSCaller creates a new read-only instance of IPFS, bound to a specific deployed contract.
func NewIPFSCaller(address common.Address, caller bind.ContractCaller) (*IPFSCaller, error) {
	contract, err := bindIPFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPFSCaller{contract: contract}, nil
}

// NewIPFSTransactor creates a new write-only instance of IPFS, bound to a specific deployed contract.
func NewIPFSTransactor(address common.Address, transactor bind.ContractTransactor) (*IPFSTransactor, error) {
	contract, err := bindIPFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPFSTransactor{contract: contract}, nil
}

// NewIPFSFilterer creates a new log filterer instance of IPFS, bound to a specific deployed contract.
func NewIPFSFilterer(address common.Address, filterer bind.ContractFilterer) (*IPFSFilterer, error) {
	contract, err := bindIPFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPFSFilterer{contract: contract}, nil
}

// bindIPFS binds a generic wrapper to an already deployed contract.
func bindIPFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPFS *IPFSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPFS.Contract.IPFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPFS *IPFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPFS.Contract.IPFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPFS *IPFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPFS.Contract.IPFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPFS *IPFSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPFS *IPFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPFS *IPFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPFS.Contract.contract.Transact(opts, method, params...)
}

// PrintIPFSIdentifiers is a free data retrieval call binding the contract method 0x62e27a91.
//
// Solidity: function printIPFSIdentifiers() view returns(string[])
func (_IPFS *IPFSCaller) PrintIPFSIdentifiers(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _IPFS.contract.Call(opts, &out, "printIPFSIdentifiers")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// PrintIPFSIdentifiers is a free data retrieval call binding the contract method 0x62e27a91.
//
// Solidity: function printIPFSIdentifiers() view returns(string[])
func (_IPFS *IPFSSession) PrintIPFSIdentifiers() ([]string, error) {
	return _IPFS.Contract.PrintIPFSIdentifiers(&_IPFS.CallOpts)
}

// PrintIPFSIdentifiers is a free data retrieval call binding the contract method 0x62e27a91.
//
// Solidity: function printIPFSIdentifiers() view returns(string[])
func (_IPFS *IPFSCallerSession) PrintIPFSIdentifiers() ([]string, error) {
	return _IPFS.Contract.PrintIPFSIdentifiers(&_IPFS.CallOpts)
}

// SetIPFSIdentifier is a paid mutator transaction binding the contract method 0x5e5a404a.
//
// Solidity: function setIPFSIdentifier(string _cid) returns()
func (_IPFS *IPFSTransactor) SetIPFSIdentifier(opts *bind.TransactOpts, _cid string) (*types.Transaction, error) {
	return _IPFS.contract.Transact(opts, "setIPFSIdentifier", _cid)
}

// SetIPFSIdentifier is a paid mutator transaction binding the contract method 0x5e5a404a.
//
// Solidity: function setIPFSIdentifier(string _cid) returns()
func (_IPFS *IPFSSession) SetIPFSIdentifier(_cid string) (*types.Transaction, error) {
	return _IPFS.Contract.SetIPFSIdentifier(&_IPFS.TransactOpts, _cid)
}

// SetIPFSIdentifier is a paid mutator transaction binding the contract method 0x5e5a404a.
//
// Solidity: function setIPFSIdentifier(string _cid) returns()
func (_IPFS *IPFSTransactorSession) SetIPFSIdentifier(_cid string) (*types.Transaction, error) {
	return _IPFS.Contract.SetIPFSIdentifier(&_IPFS.TransactOpts, _cid)
}

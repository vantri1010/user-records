// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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
	_ = abi.ConvertType
)

// EmployeeUser is an auto generated low-level Go binding around an user-defined struct.
type EmployeeUser struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"LogDeleteUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"LogNewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"LogUpdateUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deleteAddress\",\"type\":\"address\"}],\"name\":\"deleteUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structEmployee.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"insertUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isIndeed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listUsers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"}],\"name\":\"updateUserEmail\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"updateUserTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a3f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636f77926b116100715780636f77926b1461016757806381f1d75c1461017a578063b5cb15f71461018d578063c37dfd6e14610195578063e2842d79146101c0578063ffcc7bbf146101d557600080fd5b80631e75c2ec146100ae57806329434e33146100d65780634209fff1146100f75780635bc9bba61461010a5780635c60f22614610154575b600080fd5b6100c16100bc366004610866565b6101e8565b60405190151581526020015b60405180910390f35b6100e96100e4366004610890565b610289565b6040519081526020016100cd565b6100c16101053660046108c3565b610366565b6101396101183660046108c3565b60006020819052908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016100cd565b6100e96101623660046108c3565b610407565b6101396101753660046108c3565b6105b3565b6100c1610188366004610866565b610608565b6001546100e9565b6101a86101a33660046108e5565b610690565b6040516001600160a01b0390911681526020016100cd565b6101c86106ba565b6040516100cd91906108fe565b6101a86101e33660046108e5565b61081a565b60006101f383610366565b6102185760405162461bcd60e51b815260040161020f90610957565b60405180910390fd5b6001600160a01b0383166000818152602081815260409182902060018101869055600281015490548351918252918101919091529081018490527f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee906060015b60405180910390a250600192915050565b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b03861690811790915560009081526020819052604081208481558201839055815490916102f391610997565b6001600160a01b03851660008181526020818152604091829020600201849055815193845283018690528201849052907f57f7301fe8e6f5b36f73ccce215bfaee0be1dda153d32992eb9ac05da86686119060600160405180910390a26001805461035e9190610997565b949350505050565b60015460009081036103ba5760405162461bcd60e51b815260206004820152601c60248201527f4c6973742075736572206164647265737320697320656d707479202100000000604482015260640161020f565b6001600160a01b0382166000818152602081905260409020600201546001805490919081106103eb576103eb6109ae565b6000918252602090912001546001600160a01b03161492915050565b600061041282610366565b61042e5760405162461bcd60e51b815260040161020f90610957565b6001600160a01b0382166000908152602081905260408120600201546001805491929161045c908290610997565b8154811061046c5761046c6109ae565b600091825260209091200154600180546001600160a01b03909216925082918490811061049b5761049b6109ae565b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559183168152908190526040902060020182905560018054806104e7576104e76109c4565b600082815260209020810160001990810180546001600160a01b03191690550190556040516001600160a01b038516907f9a67af531bfbfd34f63109e83cfab08e30f92a1827d5c3f50e8fa815469ff9f2906105469085815260200190565b60405180910390a26001600160a01b0381166000818152602081815260409182902080546001909101548351878152928301919091528183015290517f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee9181900360600190a25092915050565b60008060006105c184610366565b6105dd5760405162461bcd60e51b815260040161020f90610957565b5050506001600160a01b03166000908152602081905260409020805460018201546002909201549092565b600061061383610366565b61062f5760405162461bcd60e51b815260040161020f90610957565b6001600160a01b0383166000818152602081815260409182902085815560028101546001909101548351918252918101869052918201527f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee90606001610278565b600181815481106106a057600080fd5b6000918252602090912001546001600160a01b0316905081565b60015460609061070c5760405162461bcd60e51b815260206004820152601960248201527f4c697374206f6620757365727320697320656d70747920212100000000000000604482015260640161020f565b60015460009067ffffffffffffffff81111561072a5761072a6109da565b60405190808252806020026020018201604052801561077557816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816107485790505b50905060005b600154811015610814576000806001838154811061079b5761079b6109ae565b60009182526020808320909101546001600160a01b031683528281019390935260409182019020815160608101835281548152600182015493810193909352600201549082015282518390839081106107f6576107f66109ae565b6020026020010181905250808061080c906109f0565b91505061077b565b50919050565b60006001828154811061082f5761082f6109ae565b6000918252602090912001546001600160a01b031692915050565b80356001600160a01b038116811461086157600080fd5b919050565b6000806040838503121561087957600080fd5b6108828361084a565b946020939093013593505050565b6000806000606084860312156108a557600080fd5b6108ae8461084a565b95602085013595506040909401359392505050565b6000602082840312156108d557600080fd5b6108de8261084a565b9392505050565b6000602082840312156108f757600080fd5b5035919050565b602080825282518282018190526000919060409081850190868401855b8281101561094a578151805185528681015187860152850151858501526060909301929085019060010161091b565b5091979650505050505050565b60208082526010908201526f55736572206e6f7420666f756e64202160801b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000828210156109a9576109a9610981565b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060018201610a0257610a02610981565b506001019056fea2646970667358221220b75bcfe49ff5d5a4d5b11c0a044bc5c28b01ac5dad24ca4a75da5eeb11a1aa5864736f6c634300080f0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[])
func (_Api *ApiCaller) GetAllUsers(opts *bind.CallOpts) ([]EmployeeUser, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllUsers")

	if err != nil {
		return *new([]EmployeeUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]EmployeeUser)).(*[]EmployeeUser)

	return out0, err

}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[])
func (_Api *ApiSession) GetAllUsers() ([]EmployeeUser, error) {
	return _Api.Contract.GetAllUsers(&_Api.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[])
func (_Api *ApiCallerSession) GetAllUsers() ([]EmployeeUser, error) {
	return _Api.Contract.GetAllUsers(&_Api.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getUser", userAddress)

	outstruct := new(struct {
		UserEmail [32]byte
		UserTime  *big.Int
		Index     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserEmail = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.UserTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiSession) GetUser(userAddress common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	return _Api.Contract.GetUser(&_Api.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiCallerSession) GetUser(userAddress common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	return _Api.Contract.GetUser(&_Api.CallOpts, userAddress)
}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 index) view returns(address userAddress)
func (_Api *ApiCaller) GetUserAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getUserAtIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 index) view returns(address userAddress)
func (_Api *ApiSession) GetUserAtIndex(index *big.Int) (common.Address, error) {
	return _Api.Contract.GetUserAtIndex(&_Api.CallOpts, index)
}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 index) view returns(address userAddress)
func (_Api *ApiCallerSession) GetUserAtIndex(index *big.Int) (common.Address, error) {
	return _Api.Contract.GetUserAtIndex(&_Api.CallOpts, index)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256 count)
func (_Api *ApiCaller) GetUserCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getUserCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256 count)
func (_Api *ApiSession) GetUserCount() (*big.Int, error) {
	return _Api.Contract.GetUserCount(&_Api.CallOpts)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256 count)
func (_Api *ApiCallerSession) GetUserCount() (*big.Int, error) {
	return _Api.Contract.GetUserCount(&_Api.CallOpts)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) view returns(bool isIndeed)
func (_Api *ApiCaller) IsUser(opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isUser", userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) view returns(bool isIndeed)
func (_Api *ApiSession) IsUser(userAddress common.Address) (bool, error) {
	return _Api.Contract.IsUser(&_Api.CallOpts, userAddress)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) view returns(bool isIndeed)
func (_Api *ApiCallerSession) IsUser(userAddress common.Address) (bool, error) {
	return _Api.Contract.IsUser(&_Api.CallOpts, userAddress)
}

// ListUsers is a free data retrieval call binding the contract method 0x5bc9bba6.
//
// Solidity: function listUsers(address ) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiCaller) ListUsers(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "listUsers", arg0)

	outstruct := new(struct {
		UserEmail [32]byte
		UserTime  *big.Int
		Index     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserEmail = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.UserTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ListUsers is a free data retrieval call binding the contract method 0x5bc9bba6.
//
// Solidity: function listUsers(address ) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiSession) ListUsers(arg0 common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	return _Api.Contract.ListUsers(&_Api.CallOpts, arg0)
}

// ListUsers is a free data retrieval call binding the contract method 0x5bc9bba6.
//
// Solidity: function listUsers(address ) view returns(bytes32 userEmail, uint256 userTime, uint256 index)
func (_Api *ApiCallerSession) ListUsers(arg0 common.Address) (struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}, error) {
	return _Api.Contract.ListUsers(&_Api.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc37dfd6e.
//
// Solidity: function userIndex(uint256 ) view returns(address)
func (_Api *ApiCaller) UserIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "userIndex", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserIndex is a free data retrieval call binding the contract method 0xc37dfd6e.
//
// Solidity: function userIndex(uint256 ) view returns(address)
func (_Api *ApiSession) UserIndex(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.UserIndex(&_Api.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc37dfd6e.
//
// Solidity: function userIndex(uint256 ) view returns(address)
func (_Api *ApiCallerSession) UserIndex(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.UserIndex(&_Api.CallOpts, arg0)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address deleteAddress) returns(uint256 index)
func (_Api *ApiTransactor) DeleteUser(opts *bind.TransactOpts, deleteAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deleteUser", deleteAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address deleteAddress) returns(uint256 index)
func (_Api *ApiSession) DeleteUser(deleteAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DeleteUser(&_Api.TransactOpts, deleteAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address deleteAddress) returns(uint256 index)
func (_Api *ApiTransactorSession) DeleteUser(deleteAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DeleteUser(&_Api.TransactOpts, deleteAddress)
}

// InsertUser is a paid mutator transaction binding the contract method 0x29434e33.
//
// Solidity: function insertUser(address userAddress, bytes32 userEmail, uint256 userTime) returns(uint256 index)
func (_Api *ApiTransactor) InsertUser(opts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte, userTime *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "insertUser", userAddress, userEmail, userTime)
}

// InsertUser is a paid mutator transaction binding the contract method 0x29434e33.
//
// Solidity: function insertUser(address userAddress, bytes32 userEmail, uint256 userTime) returns(uint256 index)
func (_Api *ApiSession) InsertUser(userAddress common.Address, userEmail [32]byte, userTime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.InsertUser(&_Api.TransactOpts, userAddress, userEmail, userTime)
}

// InsertUser is a paid mutator transaction binding the contract method 0x29434e33.
//
// Solidity: function insertUser(address userAddress, bytes32 userEmail, uint256 userTime) returns(uint256 index)
func (_Api *ApiTransactorSession) InsertUser(userAddress common.Address, userEmail [32]byte, userTime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.InsertUser(&_Api.TransactOpts, userAddress, userEmail, userTime)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x81f1d75c.
//
// Solidity: function updateUserEmail(address userAddress, bytes32 userEmail) returns(bool success)
func (_Api *ApiTransactor) UpdateUserEmail(opts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateUserEmail", userAddress, userEmail)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x81f1d75c.
//
// Solidity: function updateUserEmail(address userAddress, bytes32 userEmail) returns(bool success)
func (_Api *ApiSession) UpdateUserEmail(userAddress common.Address, userEmail [32]byte) (*types.Transaction, error) {
	return _Api.Contract.UpdateUserEmail(&_Api.TransactOpts, userAddress, userEmail)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x81f1d75c.
//
// Solidity: function updateUserEmail(address userAddress, bytes32 userEmail) returns(bool success)
func (_Api *ApiTransactorSession) UpdateUserEmail(userAddress common.Address, userEmail [32]byte) (*types.Transaction, error) {
	return _Api.Contract.UpdateUserEmail(&_Api.TransactOpts, userAddress, userEmail)
}

// UpdateUserTime is a paid mutator transaction binding the contract method 0x1e75c2ec.
//
// Solidity: function updateUserTime(address userAddress, uint256 userTime) returns(bool success)
func (_Api *ApiTransactor) UpdateUserTime(opts *bind.TransactOpts, userAddress common.Address, userTime *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateUserTime", userAddress, userTime)
}

// UpdateUserTime is a paid mutator transaction binding the contract method 0x1e75c2ec.
//
// Solidity: function updateUserTime(address userAddress, uint256 userTime) returns(bool success)
func (_Api *ApiSession) UpdateUserTime(userAddress common.Address, userTime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateUserTime(&_Api.TransactOpts, userAddress, userTime)
}

// UpdateUserTime is a paid mutator transaction binding the contract method 0x1e75c2ec.
//
// Solidity: function updateUserTime(address userAddress, uint256 userTime) returns(bool success)
func (_Api *ApiTransactorSession) UpdateUserTime(userAddress common.Address, userTime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateUserTime(&_Api.TransactOpts, userAddress, userTime)
}

// ApiLogDeleteUserIterator is returned from FilterLogDeleteUser and is used to iterate over the raw logs and unpacked data for LogDeleteUser events raised by the Api contract.
type ApiLogDeleteUserIterator struct {
	Event *ApiLogDeleteUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiLogDeleteUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLogDeleteUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiLogDeleteUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiLogDeleteUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLogDeleteUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLogDeleteUser represents a LogDeleteUser event raised by the Api contract.
type ApiLogDeleteUser struct {
	UserAddress common.Address
	Index       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogDeleteUser is a free log retrieval operation binding the contract event 0x9a67af531bfbfd34f63109e83cfab08e30f92a1827d5c3f50e8fa815469ff9f2.
//
// Solidity: event LogDeleteUser(address indexed userAddress, uint256 index)
func (_Api *ApiFilterer) FilterLogDeleteUser(opts *bind.FilterOpts, userAddress []common.Address) (*ApiLogDeleteUserIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "LogDeleteUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiLogDeleteUserIterator{contract: _Api.contract, event: "LogDeleteUser", logs: logs, sub: sub}, nil
}

// WatchLogDeleteUser is a free log subscription operation binding the contract event 0x9a67af531bfbfd34f63109e83cfab08e30f92a1827d5c3f50e8fa815469ff9f2.
//
// Solidity: event LogDeleteUser(address indexed userAddress, uint256 index)
func (_Api *ApiFilterer) WatchLogDeleteUser(opts *bind.WatchOpts, sink chan<- *ApiLogDeleteUser, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "LogDeleteUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLogDeleteUser)
				if err := _Api.contract.UnpackLog(event, "LogDeleteUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDeleteUser is a log parse operation binding the contract event 0x9a67af531bfbfd34f63109e83cfab08e30f92a1827d5c3f50e8fa815469ff9f2.
//
// Solidity: event LogDeleteUser(address indexed userAddress, uint256 index)
func (_Api *ApiFilterer) ParseLogDeleteUser(log types.Log) (*ApiLogDeleteUser, error) {
	event := new(ApiLogDeleteUser)
	if err := _Api.contract.UnpackLog(event, "LogDeleteUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiLogNewUserIterator is returned from FilterLogNewUser and is used to iterate over the raw logs and unpacked data for LogNewUser events raised by the Api contract.
type ApiLogNewUserIterator struct {
	Event *ApiLogNewUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiLogNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLogNewUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiLogNewUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiLogNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLogNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLogNewUser represents a LogNewUser event raised by the Api contract.
type ApiLogNewUser struct {
	UserAddress common.Address
	Index       *big.Int
	UserEmail   [32]byte
	UserTime    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogNewUser is a free log retrieval operation binding the contract event 0x57f7301fe8e6f5b36f73ccce215bfaee0be1dda153d32992eb9ac05da8668611.
//
// Solidity: event LogNewUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) FilterLogNewUser(opts *bind.FilterOpts, userAddress []common.Address) (*ApiLogNewUserIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "LogNewUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiLogNewUserIterator{contract: _Api.contract, event: "LogNewUser", logs: logs, sub: sub}, nil
}

// WatchLogNewUser is a free log subscription operation binding the contract event 0x57f7301fe8e6f5b36f73ccce215bfaee0be1dda153d32992eb9ac05da8668611.
//
// Solidity: event LogNewUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) WatchLogNewUser(opts *bind.WatchOpts, sink chan<- *ApiLogNewUser, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "LogNewUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLogNewUser)
				if err := _Api.contract.UnpackLog(event, "LogNewUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewUser is a log parse operation binding the contract event 0x57f7301fe8e6f5b36f73ccce215bfaee0be1dda153d32992eb9ac05da8668611.
//
// Solidity: event LogNewUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) ParseLogNewUser(log types.Log) (*ApiLogNewUser, error) {
	event := new(ApiLogNewUser)
	if err := _Api.contract.UnpackLog(event, "LogNewUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiLogUpdateUserIterator is returned from FilterLogUpdateUser and is used to iterate over the raw logs and unpacked data for LogUpdateUser events raised by the Api contract.
type ApiLogUpdateUserIterator struct {
	Event *ApiLogUpdateUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiLogUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLogUpdateUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiLogUpdateUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiLogUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLogUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLogUpdateUser represents a LogUpdateUser event raised by the Api contract.
type ApiLogUpdateUser struct {
	UserAddress common.Address
	Index       *big.Int
	UserEmail   [32]byte
	UserTime    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateUser is a free log retrieval operation binding the contract event 0x24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee.
//
// Solidity: event LogUpdateUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) FilterLogUpdateUser(opts *bind.FilterOpts, userAddress []common.Address) (*ApiLogUpdateUserIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "LogUpdateUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiLogUpdateUserIterator{contract: _Api.contract, event: "LogUpdateUser", logs: logs, sub: sub}, nil
}

// WatchLogUpdateUser is a free log subscription operation binding the contract event 0x24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee.
//
// Solidity: event LogUpdateUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) WatchLogUpdateUser(opts *bind.WatchOpts, sink chan<- *ApiLogUpdateUser, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "LogUpdateUser", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLogUpdateUser)
				if err := _Api.contract.UnpackLog(event, "LogUpdateUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUpdateUser is a log parse operation binding the contract event 0x24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee.
//
// Solidity: event LogUpdateUser(address indexed userAddress, uint256 index, bytes32 userEmail, uint256 userTime)
func (_Api *ApiFilterer) ParseLogUpdateUser(log types.Log) (*ApiLogUpdateUser, error) {
	event := new(ApiLogUpdateUser)
	if err := _Api.contract.UnpackLog(event, "LogUpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

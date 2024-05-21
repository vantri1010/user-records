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

// EmployeeUserStruct is an auto generated low-level Go binding around an user-defined struct.
type EmployeeUserStruct struct {
	UserEmail [32]byte
	UserTime  *big.Int
	Index     *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"LogDeleteUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"LogNewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"LogUpdateUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"deleteUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structEmployee.UserStruct[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"insertUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isIndeed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userEmail\",\"type\":\"bytes32\"}],\"name\":\"updateUserEmail\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userTime\",\"type\":\"uint256\"}],\"name\":\"updateUserTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506109e08061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80636f77926b116100635780636f77926b1461010357806381f1d75c14610131578063b5cb15f714610144578063e2842d791461014c578063ffcc7bbf14610162575f80fd5b80631e75c2ec1461009457806329434e33146100bc5780634209fff1146100dd5780635c60f226146100f0575b5f80fd5b6100a76100a23660046107f7565b61018d565b60405190151581526020015b60405180910390f35b6100cf6100ca36600461081f565b61022d565b6040519081526020016100b3565b6100a76100eb36600461084f565b610308565b6100cf6100fe36600461084f565b6103a6565b61011661011136600461084f565b610573565b604080519384526020840192909252908201526060016100b3565b6100a761013f3660046107f7565b6105c5565b6001546100cf565b61015461064b565b6040516100b392919061086f565b61017561017036600461090e565b6107ae565b6040516001600160a01b0390911681526020016100b3565b5f61019783610308565b6101bc5760405162461bcd60e51b81526004016101b390610925565b60405180910390fd5b6001600160a01b0383165f818152602081815260409182902060018101869055600281015490548351918252918101919091529081018490527f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee906060015b60405180910390a25060015b92915050565b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b0386169081179091555f9081526020819052604081208481558201839055815490916102969161094f565b6001600160a01b0385165f8181526020818152604091829020600201849055815193845283018690528201849052907f57f7301fe8e6f5b36f73ccce215bfaee0be1dda153d32992eb9ac05da86686119060600160405180910390a260018054610300919061094f565b949350505050565b6001545f90810361035b5760405162461bcd60e51b815260206004820152601c60248201527f4c6973742075736572206164647265737320697320656d70747920210000000060448201526064016101b3565b6001600160a01b0382165f8181526020819052604090206002015460018054909190811061038b5761038b61096e565b5f918252602090912001546001600160a01b03161492915050565b5f6103b082610308565b6103cc5760405162461bcd60e51b81526004016101b390610925565b6001600160a01b0382165f90815260208190526040812060020154600180549192916103f990829061094f565b815481106104095761040961096e565b5f91825260209091200154600180546001600160a01b0390921692508291849081106104375761043761096e565b5f91825260208083209190910180546001600160a01b0319166001600160a01b03948516179055838316808352908290526040808320938816835282208354815560018481018054838301556002808701805491909401559284529383905590829055558054806104aa576104aa610982565b5f8281526020902081015f1990810180546001600160a01b03191690550190556040516001600160a01b038516907f9a67af531bfbfd34f63109e83cfab08e30f92a1827d5c3f50e8fa815469ff9f2906105079085815260200190565b60405180910390a26001600160a01b0381165f818152602081815260409182902080546001909101548351878152928301919091528183015290517f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee9181900360600190a25092915050565b5f805f61057f84610308565b61059b5760405162461bcd60e51b81526004016101b390610925565b5050506001600160a01b03165f908152602081905260409020805460018201546002909201549092565b5f6105cf83610308565b6105eb5760405162461bcd60e51b81526004016101b390610925565b6001600160a01b0383165f818152602081815260409182902085815560028101546001909101548351918252918101869052918201527f24659ed7fa20faefd192de94aa063d3b6e668e49b0ca2738ca4d866db1ae9eee9060600161021b565b6060805f60018054905067ffffffffffffffff81111561066d5761066d610996565b6040519080825280602002602001820160405280156106b657816020015b604080516060810182525f80825260208083018290529282015282525f1990920191018161068b5790505b5090505f5b600154811015610747575f80600183815481106106da576106da61096e565b5f9182526020808320909101546001600160a01b031683528281019390935260409182019020815160608101835281548152600182015493810193909352600201549082015282518390839081106107345761073461096e565b60209081029190910101526001016106bb565b508060018080548060200260200160405190810160405280929190818152602001828054801561079e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610780575b5050505050905092509250509091565b5f600182815481106107c2576107c261096e565b5f918252602090912001546001600160a01b031692915050565b80356001600160a01b03811681146107f2575f80fd5b919050565b5f8060408385031215610808575f80fd5b610811836107dc565b946020939093013593505050565b5f805f60608486031215610831575f80fd5b61083a846107dc565b95602085013595506040909401359392505050565b5f6020828403121561085f575f80fd5b610868826107dc565b9392505050565b604080825283518282018190525f9190606090818501906020808901865b838110156108ba57815180518652838101518487015287015187860152938501939082019060010161088d565b5050868303818801528751808452928101945091925050858101905f905b838210156109015782516001600160a01b031685529384019391820191600191909101906108d8565b5092979650505050505050565b5f6020828403121561091e575f80fd5b5035919050565b60208082526010908201526f55736572206e6f7420666f756e64202160801b604082015260600190565b8181038181111561022757634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52604160045260245ffdfea2646970667358221220e1b4da7da3c201a064574d8cea7c820f11844847248426dc24a889d2fa9134a864736f6c63430008190033",
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
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[], address[])
func (_Api *ApiCaller) GetAllUsers(opts *bind.CallOpts) ([]EmployeeUserStruct, []common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllUsers")

	if err != nil {
		return *new([]EmployeeUserStruct), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]EmployeeUserStruct)).(*[]EmployeeUserStruct)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[], address[])
func (_Api *ApiSession) GetAllUsers() ([]EmployeeUserStruct, []common.Address, error) {
	return _Api.Contract.GetAllUsers(&_Api.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns((bytes32,uint256,uint256)[], address[])
func (_Api *ApiCallerSession) GetAllUsers() ([]EmployeeUserStruct, []common.Address, error) {
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

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_Api *ApiTransactor) DeleteUser(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deleteUser", userAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_Api *ApiSession) DeleteUser(userAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DeleteUser(&_Api.TransactOpts, userAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_Api *ApiTransactorSession) DeleteUser(userAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.DeleteUser(&_Api.TransactOpts, userAddress)
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

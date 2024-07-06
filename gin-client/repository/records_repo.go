package repository

import (
	"context"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"user-records/api"
)

type RecordsRepository struct {
	caller *api.Api
}

func NewRecordRepo(caller *api.Api) *RecordsRepository {
	return &RecordsRepository{caller: caller}
}

// Generate 20 random bytes
func genKey() common.Address {
	// Generate 20 random bytes
	addressBytes := make([]byte, 20)
	_, err := rand.Read(addressBytes)
	if err != nil {
		panic(err)
	}

	return common.Address(addressBytes)
}

func (repo *RecordsRepository) Create(ctx context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error {
	tran, err := repo.caller.InsertUser(txOpts, genKey(), data.UserEmail, data.UserTime)
	if err != nil {
		return err
	}

	log.Println("tran value: ", tran)
	return nil
}

func (repo *RecordsRepository) GetUserByAddr(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error) {
	var user api.EmployeeUser
	user, err := repo.caller.GetUser(opts, addr)
	return &user, err
}

func (repo *RecordsRepository) GetAddrByIndex(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error) {
	addr, err := repo.caller.GetUserAtIndex(opts, index)
	return &addr, err
}

func (repo *RecordsRepository) GetAllUsers(ctx context.Context, opts *bind.CallOpts) ([]api.EmployeeUser, error) {
	return repo.caller.GetAllUsers(opts)
}

func (repo *RecordsRepository) GetUserCount(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	return repo.caller.GetUserCount(opts)
}

func (repo *RecordsRepository) IsUser(ctx context.Context, opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	return repo.caller.IsUser(opts, userAddress)
}

func (repo *RecordsRepository) DeleteUser(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return repo.caller.DeleteUser(opts, userAddress)
}

func (repo *RecordsRepository) UpdateUserEmail(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte) (*types.Transaction, error) {
	return repo.caller.UpdateUserEmail(opts, userAddress, userEmail)
}

func (repo *RecordsRepository) UpdateUserTime(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address, userTime *big.Int) (*types.Transaction, error) {
	return repo.caller.UpdateUserTime(opts, userAddress, userTime)
}

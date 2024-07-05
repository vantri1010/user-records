package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"user-records/api"
	"user-records/gin-client/model"
)

type CreateRepo interface {
	Create(context context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error
	GetUserByAddr(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error)
	GetAddrByIndex(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error)
	GetAllUsers(opts *bind.CallOpts) ([]api.EmployeeUser, error)
	GetUserCount(opts *bind.CallOpts) (*big.Int, error)
	IsUser(opts *bind.CallOpts, userAddress common.Address) (bool, error)
	DeleteUser(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error)
	UpdateUserEmail(opts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte) (*types.Transaction, error)
	UpdateUserTime(opts *bind.TransactOpts, userAddress common.Address, userTime *big.Int) (*types.Transaction, error)
}

type CreateBiz struct {
	repo CreateRepo
}

func NewCreateBiz(repo CreateRepo) *CreateBiz {
	return &CreateBiz{repo: repo}
}

func (biz *CreateBiz) CreateRecord(ctx context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error {
	if err := biz.repo.Create(ctx, txOpts, data); err != nil {
		return err
	}

	return nil
}

func (biz *CreateBiz) GetUserByAddr(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error) {
	user, err := biz.repo.GetUserByAddr(ctx, opts, addr)
	if user.UserEmail == *new([32]byte) {
		return nil, model.ErrNotFound
	}
	return user, err
}

func (biz *CreateBiz) GetUserByID(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*api.EmployeeUser, error) {
	addr, err := biz.repo.GetAddrByIndex(ctx, opts, index)
	if *addr == *new(common.Address) {
		return nil, model.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	user, err := biz.repo.GetUserByAddr(ctx, opts, *addr)
	if user.UserEmail == *new([32]byte) {
		return nil, model.ErrNotFound
	}
	return user, err
}

func (biz *CreateBiz) GetAddress(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error) {
	addr, err := biz.repo.GetAddrByIndex(ctx, opts, index)
	if *addr == *new(common.Address) {
		return nil, model.ErrNotFound
	}

	return addr, err
}

func (biz *CreateBiz) GetAll(ctx context.Context, opts *bind.CallOpts) ([]api.EmployeeUser, error) {
	users, err := biz.repo.GetAllUsers(opts)
	if err != nil {
		return nil, err
	}
	return users, nil
}

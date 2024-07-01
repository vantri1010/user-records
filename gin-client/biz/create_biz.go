package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"user-records/api"
	"user-records/gin-client/model"
)

type CreateRepo interface {
	Create(context context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error
	GetUserByAddr(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error)
	GetAddrByIndex(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error)
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

func (biz *CreateBiz) ReadRecord(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error) {
	user, err := biz.repo.GetUserByAddr(ctx, opts, addr)
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

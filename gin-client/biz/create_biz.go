package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"user-records/api"
)

type CreateRepo interface {
	Create(context context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error
	Read(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error)
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
	user, err := biz.repo.Read(ctx, opts, addr)
	if err != nil {
		return &api.EmployeeUser{}, err
	}
	return user, nil
}

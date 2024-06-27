package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"user-records/api"
)

type Record interface {
	Create(context context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error
}

type RecordBiz struct {
	repo Record
}

func NewRecord(repo Record) *RecordBiz {
	return &RecordBiz{repo: repo}
}

func (biz *RecordBiz) CreateRecord(ctx context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error {
	if err := biz.repo.Create(ctx, txOpts, data); err != nil {
		return err
	}

	return nil
}

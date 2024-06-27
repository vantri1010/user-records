package biz

import (
	"context"
	"user-records/api"
)

type Record interface {
	Create(context context.Context, data *api.EmployeeUser) error
}

type RecordBiz struct {
	repo Record
}

func NewRecord(repo Record) *RecordBiz {
	return &RecordBiz{repo: repo}
}

func (biz *RecordBiz) CreateRecord(ctx context.Context, data *api.EmployeeUser) error {
	if err := biz.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

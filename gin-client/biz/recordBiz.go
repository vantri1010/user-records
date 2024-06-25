package biz

import (
	"context"
	"user-records/gin-client/model"
)

type Record interface {
	CreateRecord(context context.Context, data *model.Employee) error
}

type RecordBiz struct {
	repo Record
}

func NewRecord(repo Record) *RecordBiz {
	return &RecordBiz{repo: repo}
}

func (biz *RecordBiz) CreateRecord(ctx context.Context, data *model.Employee) error {
	if err := biz.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

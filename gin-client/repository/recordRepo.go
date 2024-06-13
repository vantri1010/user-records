package repository

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"user-records/api"
	"user-records/gin-client/model"
)

type RecordsRepository struct {
	caller *api.Api
}

func NewRecordRepo(caller *api.Api) *RecordsRepository {
	return &RecordsRepository{caller: caller}
}

func (repo *RecordsRepository) CreateRecord(ctx context.Context, data *api.EmployeeUser) error {
	_t, err := repo.caller.InsertUser(&bind.TransactOpts{}, data., data.UserTime, data.UserEmail)
	if err != nil {
		return err
	}

	log.Println(_t)
}

package repository

import (
	"context"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
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

func (repo *RecordsRepository) Create(ctx context.Context, data api.EmployeeUser) error {
	tran, err := repo.caller.InsertUser(&bind.TransactOpts{}, genKey(), data.UserEmail, data.UserTime)
	if err != nil {
		return err
	}

	log.Println("tran value: ", tran.Value())
	return nil
}

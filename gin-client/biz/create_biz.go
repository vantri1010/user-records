package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
	"user-records/api"
	"user-records/gin-client/model"
)

type CreateRepo interface {
	Create(context context.Context, txOpts *bind.TransactOpts, data *api.EmployeeUser) error
	GetUserByAddr(ctx context.Context, opts *bind.CallOpts, addr common.Address) (*api.EmployeeUser, error)
	GetAddrByIndex(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error)
	GetAllUsers(ctx context.Context, opts *bind.CallOpts) ([]api.EmployeeUser, error)
	GetUserCount(ctx context.Context, opts *bind.CallOpts) (*big.Int, error)
	IsUser(ctx context.Context, opts *bind.CallOpts, userAddress common.Address) (bool, error)
	DeleteUser(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error)
	UpdateUserEmail(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte) (*types.Transaction, error)
	UpdateUserTime(ctx context.Context, opts *bind.TransactOpts, userAddress common.Address, userTime *big.Int) (*types.Transaction, error)
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
	if err != nil {
		if strings.Contains(err.Error(), "User not found !") {
			return nil, model.ErrNotFound
		}
		return nil, err
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

	if err != nil {
		return nil, err
	}

	return user, err
}

func (biz *CreateBiz) GetAddrByID(ctx context.Context, opts *bind.CallOpts, index *big.Int) (*common.Address, error) {
	addr, err := biz.repo.GetAddrByIndex(ctx, opts, index)
	if *addr == *new(common.Address) {
		return nil, model.ErrNotFound
	}

	return addr, err
}

func (biz *CreateBiz) GetAll(ctx context.Context, opts *bind.CallOpts) ([]api.EmployeeUser, error) {
	users, err := biz.repo.GetAllUsers(ctx, opts)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (biz *CreateBiz) DeleteUser(ctx context.Context, txOpts *bind.TransactOpts, userAddress common.Address) (*big.Int, error) {
	tran, err := biz.repo.DeleteUser(ctx, txOpts, userAddress)
	if err != nil {
		if strings.Contains(err.Error(), "User not found !") {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return tran.Value(), nil
}

func (biz *CreateBiz) GetCounts(ctx context.Context, opts *bind.CallOpts) (*big.Int, error) {
	count, err := biz.repo.GetUserCount(ctx, opts)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (biz *CreateBiz) IsAddrExist(ctx context.Context, opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	exist, err := biz.repo.IsUser(ctx, opts, userAddress)
	if err != nil {
		if strings.Contains(err.Error(), "List user address is empty !") {
			return false, model.ErrEmpty
		}
		return false, err
	}

	return exist, nil
}

func (biz *CreateBiz) UpdateMail(ctx context.Context, txOpts *bind.TransactOpts, userAddress common.Address, userEmail [32]byte) (bool, error) {
	tran, err := biz.repo.UpdateUserEmail(ctx, txOpts, userAddress, userEmail)
	if err != nil {
		if strings.Contains(err.Error(), "User not found !") {
			return false, model.ErrNotFound
		}
		return false, err
	}
	log.Println(tran)

	return true, nil
}

func (biz *CreateBiz) UpdateTime(ctx context.Context, txOpts *bind.TransactOpts, userAddress common.Address, userTime *big.Int) (bool, error) {
	tran, err := biz.repo.UpdateUserTime(ctx, txOpts, userAddress, userTime)
	if err != nil {
		if strings.Contains(err.Error(), "User not found !") {
			return false, model.ErrNotFound
		}
		return false, err
	}
	log.Println(tran)

	return true, nil
}

func (biz *CreateBiz) UpdateUser(ctx context.Context, txOpts *bind.TransactOpts, userAddress common.Address, data *api.EmployeeUser) (bool, error) {
	if data.UserEmail != *new([32]byte) {
		return biz.UpdateMail(ctx, txOpts, userAddress, data.UserEmail)
	} else {
		return biz.UpdateTime(ctx, txOpts, userAddress, data.UserTime)
	}
}

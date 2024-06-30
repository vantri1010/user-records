package model

import (
	"bytes"
	"math/big"
	"time"
	"user-records/api"
)

type Employee struct {
	UserEmail string    `json:"user_email" validate:"required"`
	UserTime  time.Time `json:"user_time,omitempty"`
	//Index     *big.Int  `json:"index,omitempty"`
}

type EmployeeRead struct {
	UserEmail string    `json:"user_email,omitempty"`
	UserTime  time.Time `json:"user_time,omitempty"`
	Index     *big.Int  `json:"index,omitempty"`
}

func MapToSolData(emp Employee) (*api.EmployeeUser, error) {
	// Convert UserEmail from string to [32]byte
	var userEmail [32]byte
	copy(userEmail[:], emp.UserEmail)

	// Convert UserTime from time.Time to *big.Int (Unix timestamp)
	if emp.UserEmail == "" {
		return nil, ErrBadParamInput
	}

	userTime := big.NewInt(emp.UserTime.Unix())
	if userTime.Cmp(big.NewInt(0)) <= 0 {
		userTime = big.NewInt(time.Now().Unix())
	}

	return &api.EmployeeUser{
		UserEmail: userEmail,
		UserTime:  userTime,
		Index:     nil,
	}, nil
}
func MapSolDataToEmployee(empUser *api.EmployeeUser) (*EmployeeRead, error) {
	// Convert UserEmail from [32]byte to string
	userEmail := string(bytes.Trim(empUser.UserEmail[:], "\x00"))

	// Convert UserTime from *big.Int to time.Time
	userTime := time.Unix(empUser.UserTime.Int64(), 0)

	return &EmployeeRead{
		UserEmail: userEmail,
		UserTime:  userTime,
		Index:     empUser.Index,
	}, nil
}

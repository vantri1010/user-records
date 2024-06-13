package model

import (
	"math/big"
	"time"
)

type Employee struct {
	UserEmail string    `json:"user_email,omitempty"`
	UserTime  time.Time `json:"user_time,omitempty"`
	Address   *big.Int  `json:"address,omitempty"`
}

type EmployeeUser struct {
	UserEmail string    `json:"user_email,omitempty"`
	UserTime  time.Time `json:"user_time,omitempty"`
	Address   *big.Int  `json:"address,omitempty"`
}

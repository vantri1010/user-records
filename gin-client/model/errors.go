package model

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrEmpty               = errors.New("List user address is empty !")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)

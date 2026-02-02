package domain

import "errors"

var (
	ErrInvalidName      = errors.New("invalid subscription name")
	ErrInvalidPrice     = errors.New("invalid subscription price")
	ErrInvalidStartDate = errors.New("invalid subscription start date")
)

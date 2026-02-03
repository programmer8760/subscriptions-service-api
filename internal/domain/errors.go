package domain

import "errors"

var (
	ErrInvalidID            = errors.New("invalid subscription id")
	ErrInvalidName          = errors.New("invalid subscription name")
	ErrInvalidPrice         = errors.New("invalid subscription price")
	ErrInvalidStartDate     = errors.New("invalid subscription start date")
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

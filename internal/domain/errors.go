package domain

import "errors"

var (
	ErrInvalidID            = errors.New("invalid subscription id")
	ErrInvalidName          = errors.New("invalid subscription name")
	ErrInvalidPrice         = errors.New("invalid subscription price")
	ErrInvalidStartDate     = errors.New("invalid subscription start date")
	ErrInvalidEndDate       = errors.New("invalid subscription end date")
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrNoChanges            = errors.New("request doesn't change anything")
	ErrInvalidFromDate      = errors.New("invalid period start date")
	ErrInvalidToDate        = errors.New("invalid period end date")
)

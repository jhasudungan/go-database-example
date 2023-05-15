package main

import (
	"time"
)

type Transaction struct {
	TransactionId    string
	TransactionTotal float64
	TransactionPIC   string
	TransactionDate  time.Time
	CreatedDate      time.Time
	LastModified     time.Time
}

package domain

import (
	"context"
)

type Transaction struct {
	ID           uint
	DateTime     string
	CustomerName string
	FinalAmount  float64
	Status       string
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction Transaction) error
	GetCustomer(ctx context.Context,id uint) (string,error)
	GetStatusTransaction(ctx context.Context,id uint) (string,error) 
	GetTotalTransactin(ctx context.Context, id uint) (float64,error)
}

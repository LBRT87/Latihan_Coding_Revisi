package usecase

import (
	"context"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/config"
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/transaction-service/core/domain"
)

type TransactionUsecaseStruct struct {
	TransactionRepo domain.TransactionRepository
	Cfg             *config.Config
}

type TransactionUsecaseInterface interface {
	Create(ctx context.Context, trx domain.Transaction) error
}

func NewTransactionUsecase(transRepo domain.TransactionRepository, cfg *config.Config) TransactionUsecaseInterface {
	return &TransactionUsecaseStruct{TransactionRepo: transRepo, Cfg: cfg}
}
func (u *TransactionUsecaseStruct) Create(ctx context.Context, trx domain.Transaction) error {
	return ctx.Err()
}

package repository

import (
	"context"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/transaction-service/core/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (t *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, transaction domain.Transaction) error {
	return t.db.WithContext(ctx).Create(&transaction).Error
}

func (t *TransactionRepositoryImpl) GetCustomer(ctx context.Context, id uint) (string, error) {
	var trans domain.Transaction
	return t.db.WithContext(ctx).Where("id = ?", id).First(&trans).Error

	if err != nil {
		return nil, err
	}

	return &trans, err
}

func (t *TransactionRepositoryImpl) funcGetStatusTransaction(ctx context.Context, id uint) (string, error) {
}
func (t *TransactionRepositoryImpl) GetTotalTransactin(ctx context.Context, id uint) (float64, error) {
}

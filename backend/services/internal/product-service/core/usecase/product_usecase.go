package usecase

import (
	"context"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/config"
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/domain"
)

type IceCreamUsecaseInterface interface {
	Create(ctx context.Context, req CreateIceCreamRequest, photo []byte) error
	Update(ctx context.Context, req UpdateIceCreamRequest) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.IceCream, error)
	GetAll(ctx context.Context, req ListIceCreamRequest) (*ListIceCreamResponse, error)
}

type IceCreamUsecase struct {
	IceCreamRepo domain.IceCreamRepository
	Cfg          *config.Config
}

// Create implements [IceCreamUsecaseInterface].
func (i *IceCreamUsecase) Create(ctx context.Context, req CreateIceCreamRequest, photo []byte) error {
	newIceCream := domain.IceCream{
		Name: req.Name,
		Flavour: req.Flavour,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
	}

	if err := i.IceCreamRepo.Create(ctx, newIceCream, photo); err != nil {
		return err
	}

	return nil
}

// Delete implements [IceCreamUsecaseInterface].
func (i *IceCreamUsecase) Delete(ctx context.Context, id uint) error {
	return i.IceCreamRepo.Delete(ctx, id)
}

// GetAll implements [IceCreamUsecaseInterface].
func (i *IceCreamUsecase) GetAll(ctx context.Context, req ListIceCreamRequest) (*ListIceCreamResponse, error) {

	filter := domain.IceCreamFilter{
		Name: req.Name,
		Page: req.Page,
		Limit: req.Limit,
		Random: req.Random,
	}

	iceCream, total, err := i.IceCreamRepo.GetAll(ctx, filter)

	if err != nil {
		return nil, err
	}

	return &ListIceCreamResponse{
		IceCreams: iceCream,
		TotalData: int(total),
	}, nil
}

// GetByID implements [IceCreamUsecaseInterface].
func (i *IceCreamUsecase) GetByID(ctx context.Context, id uint) (*domain.IceCream, error) {
	return i.IceCreamRepo.GetDetail(ctx, id)
}

// Update implements [IceCreamUsecaseInterface].
func (i *IceCreamUsecase) Update(ctx context.Context, req UpdateIceCreamRequest) error {
	updateData := make(map[string]interface{})

	if req.Description != "" {
		updateData["description"] = req.Description
	}

	if req.Flavour != "" {
		updateData["flavour"] = req.Flavour
	}

	if req.Price > 0 {
		updateData["price"] = req.Price
	}

	if req.Stock >= 0 {
		updateData["stock"] = req.Stock
	}

	return i.IceCreamRepo.Update(ctx, req.ID, updateData)
}

func NewIceCreamUsecase(iceCreamRepo domain.IceCreamRepository, cfg *config.Config) IceCreamUsecaseInterface {
	return &IceCreamUsecase{IceCreamRepo: iceCreamRepo, Cfg: cfg}
}
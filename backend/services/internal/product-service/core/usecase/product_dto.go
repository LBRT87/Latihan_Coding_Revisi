package usecase

import "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/domain"

type CreateIceCreamRequest struct {
	Name        string `json:"name"`
	Flavour     string `json:"flavour"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}

type ListIceCreamRequest struct {
	Name   string `json:"name"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Random bool `json:"random"`
}

type ListIceCreamResponse struct {
	IceCreams []domain.IceCream
	TotalData int
}

type UpdateIceCreamRequest struct {
	ID uint `json:"id"`
	Flavour     string `json:"flavour"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}
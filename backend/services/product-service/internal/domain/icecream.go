package domain

import "context"

type IceCream struct {
	ID          uint
	Name        string
	Flavour     string
	PictureUrl  string
	Price       float64
	Description string
}

type IceCreamRepository interface {
	CreateIceCream(c context.Context, icecream IceCream) error
	GetDetail(c context.Context, id uint) (*IceCream, error)
	UpdatePrice(c context.Context, id uint, newPrice float64) error
	DeleteIceCream(c context.Context, id uint) error
}

type IceCreamUseCase interface {
	CreateIceCream(c context.Context, icecream IceCream) error
	GetDetail(c context.Context, id uint) (*IceCream, error)
	UpdatePrice(c context.Context, id uint, newPrice float64) error
	DeleteIceCream(c context.Context, id uint) error
}

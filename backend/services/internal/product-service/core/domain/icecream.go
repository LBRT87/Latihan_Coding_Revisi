package domain

import (
	"context"
	"time"
)

type IceCream struct {
	ID          uint
	Name        string
	Description string
	Flavour     string
	PictureUrl  string
	Stock		int
	Price       int64
	CreatedAt time.Time
}

type IceCreamFilter struct {
	Name string
	Limit int
	Page int
	Random bool
}

func (f *IceCreamFilter) Offset() int {
	return (f.Page - 1) * f.Limit
}

func (f *IceCreamFilter) Normalize () {
	if f.Page <= 0 {
		f.Page = 1
	}

	if f.Limit < 25 {
		f.Limit = 25
	}else if f.Limit > 100 {
		f.Limit = 100
	}
}

type IceCreamRepository interface {
	Create(c context.Context, icecream IceCream, photo []byte) error
	GetDetail(c context.Context, id uint) (*IceCream, error)
	GetAll(c context.Context, filter IceCreamFilter) ([]IceCream, int64, error)
	Update(c context.Context, id uint, updateData map[string]interface{}) error
	Delete(c context.Context, id uint) error
}
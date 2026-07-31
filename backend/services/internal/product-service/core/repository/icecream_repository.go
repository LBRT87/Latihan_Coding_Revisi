package repository

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/domain"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type PgIceCreamRepository struct {
	client *minio.Client
	bucket string
	publicEndpoint string
	db *gorm.DB
}

// Create implements [domain.IceCreamRepository].
func (p *PgIceCreamRepository) Create(c context.Context, icecream domain.IceCream, photo []byte) error {
	photoName := strings.ReplaceAll(icecream.Name, " ", "_")
	url := fmt.Sprintf("http://%s/%s/%s", p.publicEndpoint, p.bucket, photoName)

	_, err := p.client.PutObject(c, p.bucket, photoName, bytes.NewReader(photo), int64(len(photo)), minio.PutObjectOptions{
		ContentType: http.DetectContentType(photo),
	})

	if err != nil {
		return err
	}

	icecream.PictureUrl = url

	return p.db.WithContext(c).Create(&icecream).Error
}

// Delete implements [domain.IceCreamRepository].
func (p *PgIceCreamRepository) Delete(c context.Context, id uint) error {
	return p.db.WithContext(c).Delete(&domain.IceCream{}, "id = ?", id).Error
}

// GetAll implements [domain.IceCreamRepository].
func (p *PgIceCreamRepository) GetAll(c context.Context, filter domain.IceCreamFilter) ([]domain.IceCream, int64, error) {
	filter.Normalize()

	var iceCreams []domain.IceCream
	var total int64

	query := p.db.WithContext(c).Model(&domain.IceCream{})

	if filter.Name != "" {
		query.Where("name = ?", filter.Name)
	}

	query.Count(&total)

	if filter.Random {
		query.Order("RANDOM()")
	}else {
		query.Order("created_at DESC")
	}

	if err := query.Offset(filter.Offset()).Limit(filter.Limit).Find(&iceCreams).Error; err != nil {
		return nil, 0, err
	}

	return iceCreams, total, nil
}

// GetDetail implements [domain.IceCreamRepository].
func (p *PgIceCreamRepository) GetDetail(c context.Context, id uint) (*domain.IceCream, error) {
	var icecream domain.IceCream
	err := p.db.WithContext(c).Where("id = ?", id).First(&icecream).Error

	if err != nil {
		return nil, err
	}

	return &icecream, err
}

// Update implements [domain.IceCreamRepository].
func (p *PgIceCreamRepository) Update(c context.Context, id uint, updateData map[string]interface{}) error {
	return p.db.WithContext(c).Model(&domain.IceCream{}).Where("id = ?", id).Updates(updateData).Error
}

func NewIceCreamRepository(db *gorm.DB) domain.IceCreamRepository {
	return &PgIceCreamRepository{
		db: db,
	}
}
package dao

import (
	"context"
	"xm-mall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (d *CarouselDao) List() (carousels []*model.Carousel, err error) {
	err = d.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}

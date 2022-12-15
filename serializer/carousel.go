package serializer

import "xm-mall/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []*model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousels = append(carousels, BuildCarousel(item))
	}
	return
}

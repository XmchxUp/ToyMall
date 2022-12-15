package service

import (
	"context"
	"xm-mall/dao"
	"xm-mall/pkg/e"
	"xm-mall/serializer"

	logging "github.com/sirupsen/logrus"
)

type ListCarouselsService struct {
}

func (s *ListCarouselsService) List() serializer.Response {
	code := e.SUCCESS
	dao := dao.NewCarouselDao(context.Background())
	carousels, err := dao.List()
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}

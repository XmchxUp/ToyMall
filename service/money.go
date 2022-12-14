package service

import (
	"context"
	"xm-mall/dao"
	"xm-mall/pkg/e"
	"xm-mall/serializer"

	logging "github.com/sirupsen/logrus"
)

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (s *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.MakeMoney(user, s.Key),
		Msg:    e.GetMsg(code),
	}
}

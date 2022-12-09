package service

import (
	"context"
	"xm-mall/dao"
	"xm-mall/model"
	"xm-mall/pkg/e"
	"xm-mall/pkg/utils"
	"xm-mall/serializer"

	logging "github.com/sirupsen/logrus"
)

type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行验证
}

func (s UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	if s.Key == "" || len(s.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	utils.Encrypt.SetKey(s.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(s.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		NickName: s.NickName,
		UserName: s.UserName,
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"), //初始金额
	}
	if err := user.SetPassword(s.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = "avatar.jpg"
	err = userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

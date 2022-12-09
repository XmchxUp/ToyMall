package dao

import (
	"context"
	"xm-mall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Error
	return
}

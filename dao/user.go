package dao

import (
	"context"

	"gorm.io/gorm"
)

type UserDao struct{
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) Create(ctx context.Context,user User) error {
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
} 

func(u *UserDao) IsExist(ctx context.Context, id string) (bool, error) {
	var count int64
	if err := u.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *UserDao) UpdateName(ctx context.Context, id, name string) error {
	if err := u.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}
package service

import (
	"GuiXinSchool/dao"
	"context"
	"fmt"
)


type UserSvc struct{
	ud *dao.UserDao
}

func(u *UserSvc) CreateIfNotExist(ctx context.Context, user dao.User) error {
	//检查用户是否存在
	exist, err := u.ud.IsExist(ctx, user.ID)
	if err != nil {
		return err
	}
	if exist {
		return nil //用户已存在，不需要创建
	}
	//创建用户
	return u.ud.Create(ctx, user)
}

func(u *UserSvc) UpdateName(ctx context.Context, id, name string) error {
	//检查用户是否存在
	exist, err := u.ud.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("user is not exist") //用户不存在
	}
	//更新用户名
	return u.ud.UpdateName(ctx, id, name)
}
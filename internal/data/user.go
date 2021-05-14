package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hellworld/internal/biz"
)

/**
 * @Date: 2021/5/14 10:48 上午
 * @Desc:
 */
var _ biz.UserRepo =(*userRepo)(nil)

type userRepo struct {
	date *Data
	log *log.Helper
}

func NewUserRepo(data *Data,logger log.Logger) biz.UserRepo  {
	return &userRepo{
		date: data,
		log:  log.NewHelper("data/user",logger),
	}
}

func (u userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	//TODO mock数据->mysql
	return &biz.User{Id:1,UserName: "李四"}  ,nil
}


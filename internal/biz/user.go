package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

/**
 * @Date: 2021/5/14 10:51 上午
 * @Desc:
 */
type User struct {
	Id int64
	UserName string
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64)(*User,error)
}

type UserUseCase struct {
	repo UserRepo
	log *log.Helper
}

func NewUserUseCase(repo UserRepo,logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper("usecase/user",logger),
	}
}
func (u *UserUseCase)GetUser(ctx context.Context,id int64)(*User,error)  {
	return u.repo.GetUser(ctx,id)
}
package data

import (
	"context"
	"kratos-hello/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, loggger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(loggger),
	}
}

func (us *userRepo) UserLogin(cxt context.Context, ul *biz.UserLogin) (ulr *biz.UserloginResponse, err error) {
	return nil, nil
}

func (us *userRepo) UserSignup(ctx context.Context, usp *biz.UserSignup) (msg string, err error) {
	return "", nil
}

func (us *userRepo) CheckUser(ctx context.Context, usp *biz.UserSignup) (rowAff int, err error) {
	return 0, nil
}

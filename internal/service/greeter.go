package service

import (
	"context"

	v1 "kratos-hello/api/helloworld/v1"
	"kratos-hello/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.

func (s *GreeterService) Login(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	return nil, nil
}

func (s *GreeterService) Singup(ctx context.Context, req *v1.UserSignupRequest) (reply *v1.UserSignupResponse, err error) {
	if req.Password != req.RePassword {
		reply.Msg = "请确认密码"
		return
	}
	userSignup := &biz.UserSignup{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
    n,err:=s.uc.CheckUserByName(context.Background(),userSignup)
    if err != nil {
		return nil, err
	} 
	if n==0{
		reply.Msg="注册成功"
	}
	return nil, nil
}

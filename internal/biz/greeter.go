package biz

import (
	"context"

	v1 "kratos-hello/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// user login
type UserLogin struct {
	Username string
	Password string
	Email    string
}

type UserloginResponse struct {
	Msg   string
	Token string
}

type UserSignup struct {
	Username   string
	Password   string
	RePassword string
	Email      string
}

type UserRepo interface {
	UserLogin(context.Context, *UserLogin) (*UserloginResponse, error)
	UserSignup(context.Context, *UserSignup) (string, error)
	CheckUser(context.Context, *UserSignup) (int, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(ur UserRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{ur: ur, log: log.NewHelper(logger)}
}

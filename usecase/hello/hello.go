package usecase_hello

import "context"

type usecase struct {
}

func NewUsecase(ctx context.Context) *usecase {
	return &usecase{}
}

func (u *usecase) Hello() string {
	return "Hello"
}

func (u *usecase) Hi() string {
	return "Hi"
}

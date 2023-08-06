package usecase_auth

import "context"

type usecase struct {
}

func NewUsecase(ctx context.Context) *usecase {
	return &usecase{}
}

func (u *usecase) IsAuthorized() bool {
	return true
}

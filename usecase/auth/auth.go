package usecase_auth

import (
	"context"
	"fmt"

	"github.com/hengkyawijaya/simple-go/repository"
)

type usecase struct {
	repo *repository.Repository
}

func NewUsecase(ctx context.Context, repo *repository.Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) IsAuthorized(ctx context.Context, token string) bool {
	username, err := u.repo.DatabaseRepository.GetUser(context.TODO(), 1)
	if err != nil {
		fmt.Println("error: ", err)
		return false
	}

	fmt.Println("username: ", username)
	return true
}

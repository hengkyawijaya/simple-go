package usecase

import (
	"context"

	"github.com/hengkyawijaya/simple-go/repository"
	usecase_auth "github.com/hengkyawijaya/simple-go/usecase/auth"
	usecase_hello "github.com/hengkyawijaya/simple-go/usecase/hello"
)

// HelloUsecase is an interface for hello
type HelloUsecase interface {
	Hello() string
	Hi() string
}

// AuthUsecase is an interface for authorization
type AuthUsecase interface {
	IsAuthorized(ctx context.Context, token string) bool
}

// Usecase is a struct that contains all usecase interfaces
type Usecase struct {
	HelloUsecase HelloUsecase
	AuthUsecase  AuthUsecase
}

// NewUsecase is a constructor for Usecase
func NewUsecase(ctx context.Context, repo *repository.Repository) *Usecase {
	return &Usecase{
		HelloUsecase: usecase_hello.NewUsecase(ctx),
		AuthUsecase:  usecase_auth.NewUsecase(ctx, repo),
	}
}

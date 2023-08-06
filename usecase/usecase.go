package usecase

import (
	"context"

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
	IsAuthorized() bool
}

// Usecase is a struct that contains all usecase interfaces
type Usecase struct {
	HelloUsecase HelloUsecase
	AuthUsecase  AuthUsecase
}

// NewUsecase is a constructor for Usecase
func NewUsecase(ctx context.Context) *Usecase {
	return &Usecase{
		HelloUsecase: usecase_hello.NewUsecase(ctx),
		AuthUsecase:  usecase_auth.NewUsecase(ctx),
	}
}

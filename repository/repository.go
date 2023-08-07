package repository

import (
	"context"

	"github.com/hengkyawijaya/simple-go/repository/config"
	"github.com/hengkyawijaya/simple-go/repository/database"
	"github.com/hengkyawijaya/simple-go/repository/ipinfo"
)

// ConfigRepository is an interface for reading config
type ConfigRepository interface {
	ReadConfig() *repository_config.Config
}

// IPInfoRepository is an interface for getting IP info
type IPInfoRepository interface {
	GetIPInfo() (string, error)
}

type DatabaseRepository interface {
	GetUser(ctx context.Context, id int) (string, error)
}

// Repository is a struct that contains all repository interfaces
type Repository struct {
	ConfigRepository   ConfigRepository
	IPInfoRepository   IPInfoRepository
	DatabaseRepository DatabaseRepository
}

// NewRepository is a constructor for Repository
func NewRepository(ctx context.Context) *Repository {
	cfgRepo := repository_config.NewConfig(ctx)
	return &Repository{
		ConfigRepository:   cfgRepo,
		IPInfoRepository:   repository_ipinfo.NewIPInfo(ctx, cfgRepo.ReadConfig()),
		DatabaseRepository: repository_database.NewDatabase(ctx, cfgRepo.ReadConfig()),
	}
}

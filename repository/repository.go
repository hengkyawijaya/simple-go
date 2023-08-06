package repository

import (
	"context"

	"github.com/hengkyawijaya/simple-go/repository/config"
	repository_ipinfo "github.com/hengkyawijaya/simple-go/repository/ipinfo"
)

// ConfigRepository is an interface for reading config
type ConfigRepository interface {
	ReadConfig() *repository_config.Config
}

// IPInfoRepository is an interface for getting IP info
type IPInfoRepository interface {
	GetIPInfo() (string, error)
}

// Repository is a struct that contains all repository interfaces
type Repository struct {
	ConfigRepository ConfigRepository
	IPInfoRepository IPInfoRepository
}

// NewRepository is a constructor for Repository
func NewRepository(ctx context.Context) *Repository {
	cfgRepo := repository_config.NewConfig(ctx)
	return &Repository{
		ConfigRepository: cfgRepo,
		IPInfoRepository: repository_ipinfo.NewIPInfo(ctx, cfgRepo.ReadConfig()),
	}
}

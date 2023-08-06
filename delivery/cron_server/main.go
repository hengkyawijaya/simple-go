package cron_server

import (
	"log"
	"sync"

	"github.com/hengkyawijaya/simple-go/repository"
	"github.com/hengkyawijaya/simple-go/usecase"
	"github.com/robfig/cron"
)

// CronServer is an interface for cron server
type CronServer interface {
	Hello()
	GetIP()
}

// cronServer is a struct that contains all dependencies for cron server
type cronServer struct {
	uc   *usecase.Usecase
	repo *repository.Repository
}

// NewCronServer is a constructor for CronServer
func NewCronServer(uc *usecase.Usecase, repo *repository.Repository) *cronServer {
	return &cronServer{
		uc:   uc,
		repo: repo,
	}
}

// DeliverCronServer is a function to deliver cron server
func DeliverCronServer(
	uc *usecase.Usecase,
	wg sync.WaitGroup,
	repo *repository.Repository,
) {
	c := cron.New()

	cfg := repo.ConfigRepository.ReadConfig()
	m := NewCronServer(uc, repo)

	c.AddFunc(cfg.CronServer.CronHello.Schedule, m.Hello)
	c.AddFunc(cfg.CronServer.CronHello.Schedule, m.GetIP)

	log.Println("Starting cron server...")
	c.Start()
}

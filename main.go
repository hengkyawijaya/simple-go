package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/hengkyawijaya/simple-go/constant"
	"github.com/hengkyawijaya/simple-go/delivery/cron_server"
	"github.com/hengkyawijaya/simple-go/delivery/grpc_server"
	"github.com/hengkyawijaya/simple-go/delivery/http_server"
	"github.com/hengkyawijaya/simple-go/repository"
	"github.com/hengkyawijaya/simple-go/usecase"
)

type ServerFunc func(uc *usecase.Usecase, wg sync.WaitGroup, repo *repository.Repository)

func main() {
	ctx := buildContext()
	repo := repository.NewRepository(ctx)
	uc := usecase.NewUsecase(ctx, repo)

	listServerFunc := []ServerFunc{
		http_server.DeliverHTTP,
		grpc_server.DeliverGRPCServer,
		cron_server.DeliverCronServer,
	}
	initiate(listServerFunc, uc, repo)
}

func initiate(
	listServerFunc []ServerFunc,
	uc *usecase.Usecase,
	repo *repository.Repository,
) {
	wg := sync.WaitGroup{}
	wg.Add(len(listServerFunc))
	for _, serverFunc := range listServerFunc {
		go serverFunc(uc, wg, repo)
	}
	wg.Wait()
}

func buildContext() context.Context {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, constant.EnvContextKey, ReadEnv())
	return ctx
}

func ReadEnv() string {
	appEnv := os.Getenv("APP_ENV")
	fmt.Println("APP_ENV:", appEnv)
	_, ok := constant.EnvMap[appEnv]
	if !ok {
		return constant.DevelopmentEnv
	}

	return appEnv
}

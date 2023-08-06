package http_server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/hengkyawijaya/simple-go/repository"
	"github.com/hengkyawijaya/simple-go/usecase"
)

// HTTPServer is an interface for http server
type HTTPServer interface {
	Hello(w http.ResponseWriter, r *http.Request)
	Hi(w http.ResponseWriter, r *http.Request)
}

// HTTPServerModule is a struct that contains all dependencies for http server
type HTTPServerModule struct {
	Uc *usecase.Usecase
}

// NewHTTPServer is a constructor for HTTPServerModule
func NewHTTPServer(uc *usecase.Usecase) *HTTPServerModule {
	return &HTTPServerModule{
		Uc: uc,
	}
}

// DeliverHTTP is a function to deliver http server
func DeliverHTTP(
	uc *usecase.Usecase,
	wg sync.WaitGroup,
	repo *repository.Repository,
) {
	cfg := repo.ConfigRepository.ReadConfig()
	m := NewHTTPServer(uc)
	r := Router(uc, m)
	log.Printf("Starting http server on port %s...\n", cfg.HTTPServer.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPServer.Port), r))
	wg.Done()
}
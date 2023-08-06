package http_server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hengkyawijaya/simple-go/usecase"
)

func Router(uc *usecase.Usecase, m *HTTPServerModule) *mux.Router {
	r := mux.NewRouter()
	r.Use(logging)
	public := r.PathPrefix("/public").Subrouter()
	public.Use(publicMiddleware)

	public.HandleFunc("/", m.Hello).Methods(http.MethodGet)

	authMid := NewAuthMiddleware(uc.AuthUsecase)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.Use(authMid.authMiddleware)
	auth.HandleFunc("/", m.Hi).Methods(http.MethodGet)

	return r
}

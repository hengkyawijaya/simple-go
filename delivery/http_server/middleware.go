package http_server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hengkyawijaya/simple-go/constant"
	"github.com/hengkyawijaya/simple-go/usecase"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func publicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PUBLIC")
		next.ServeHTTP(w, r)
	})
}

type authMiddleware struct {
	authUc usecase.AuthUsecase
}

func NewAuthMiddleware(authUc usecase.AuthUsecase) *authMiddleware {
	return &authMiddleware{
		authUc: authUc,
	}
}

func (a authMiddleware) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// validate bearer token
		if len(authorization) < 7 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// remove bearer prefix
		token := authorization[7:]

		isAuth := a.authUc.IsAuthorized(r.Context(), token)
		fmt.Println("AUTH:", isAuth)
		if !isAuth {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		context.WithValue(r.Context(), constant.AuthContextKey, "true")

		next.ServeHTTP(w, r)
	})
}

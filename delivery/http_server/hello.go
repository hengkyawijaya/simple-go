package http_server

import "net/http"

// Hello is a handler for hello
func (h *HTTPServerModule) Hello(w http.ResponseWriter, r *http.Request) {
	res := h.Uc.HelloUsecase.Hello()
	w.Write([]byte(res))
}

// Hi is a handler for hi
func (h *HTTPServerModule) Hi(w http.ResponseWriter, r *http.Request) {
	res := h.Uc.HelloUsecase.Hi()
	w.Write([]byte(res))
}

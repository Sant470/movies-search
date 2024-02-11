package apis

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/sant470/moviesearch/apis/v1"
)

type Handler func(rw http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Internal Server Error", err)
	}
}

func InitSerachRoutes(r *chi.Mux, sh *v1.SearchHandler) {
	r.Route("/v1/search", func(r chi.Router) {
		r.Method(http.MethodGet, "/helloworld", Handler(sh.HelloWorld))
	})
}

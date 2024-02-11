package v1

import (
	"log"
	"net/http"

	"github.com/sant470/moviesearch/services"
)

type SearchHandler struct {
	lgr *log.Logger
	svc *services.SearchService
}

func NewSearchHandler(lgr *log.Logger, svc *services.SearchService) *SearchHandler {
	return &SearchHandler{lgr, svc}
}

func (sh *SearchHandler) HelloWorld(rw http.ResponseWriter, r *http.Request) error {
	rw.Write([]byte("hello World"))
	return nil
}

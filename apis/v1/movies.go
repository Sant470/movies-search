package v1

import (
	"log"

	"github.com/sant470/moviesearch/services"
)

type SearchHandler struct {
	lgr *log.Logger
	svc *services.SearchService
}

func NewSearchHandler(lgr *log.Logger, svc *services.SearchService) *SearchHandler {
	return &SearchHandler{lgr, svc}
}

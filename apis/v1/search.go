package v1

import (
	"log"
	"net/http"

	"github.com/sant470/moviesearch/services"
	"github.com/sant470/moviesearch/utils/errors"
	"github.com/sant470/moviesearch/utils/res"
)

type SearchHandler struct {
	lgr *log.Logger
	svc *services.SearchService
}

func NewSearchHandler(lgr *log.Logger, svc *services.SearchService) *SearchHandler {
	return &SearchHandler{lgr, svc}
}

func (sh *SearchHandler) HelloWorld(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	return res.OK(rw, "hello, world!")
}

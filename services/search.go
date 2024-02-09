package services

import (
	"log"

	"github.com/sant470/moviesearch/daos"
)

type SearchService struct {
	lgr *log.Logger
	dao *daos.SearchDao
}

func NewSearchService(lgr *log.Logger, dao *daos.SearchDao) *SearchService {
	return &SearchService{lgr, dao}
}

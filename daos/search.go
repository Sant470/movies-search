package daos

import (
	"log"

	"github.com/redis/go-redis/v9"
)

type SearchDao struct {
	lgr *log.Logger
	dao *redis.Client
}

func NewSearchDao(lgr *log.Logger, rdb *redis.Client) *SearchDao {
	return &SearchDao{lgr, rdb}
}

package main

import (
	"fmt"
	"log"

	handlers "github.com/sant470/moviesearch/apis/v1"
	"github.com/sant470/moviesearch/config"
	"github.com/sant470/moviesearch/daos"
	"github.com/sant470/moviesearch/services"
)

func main() {
	lgr := log.Default()
	lgr.Println("info: starting the server")
	appConf := config.GetAppConfig("config", "./")
	rdb := config.GetRedisClient(lgr, &appConf.Redis)
	searchDao := daos.NewSearchDao(lgr, rdb)
	searchSvc := services.NewSearchService(lgr, searchDao)
	searchHlr := handlers.NewSearchHandler(lgr, searchSvc)
	fmt.Println("search svc: ", searchHlr)
}

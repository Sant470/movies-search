package main

import (
	"log"
	"net/http"

	"github.com/sant470/moviesearch/apis"
	handlers "github.com/sant470/moviesearch/apis/v1"
	"github.com/sant470/moviesearch/config"
	"github.com/sant470/moviesearch/daos"
	"github.com/sant470/moviesearch/services"
)

func main() {
	lgr := log.Default()
	lgr.Println("info: starting the server")
	appConf := config.GetAppConfig("config", "./")
	router := config.InitRouters()
	rdb := config.GetRedisClient(lgr, &appConf.Redis)
	searchDao := daos.NewSearchDao(lgr, rdb)
	searchSvc := services.NewSearchService(lgr, searchDao)
	searchHlr := handlers.NewSearchHandler(lgr, searchSvc)
	apis.InitSerachRoutes(router, searchHlr)
	http.ListenAndServe("localhost:8000", router)
}

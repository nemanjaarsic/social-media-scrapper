package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	config "social-media-api/config"
	"social-media-api/controller"
	"social-media-api/repo"
	"social-media-api/routes"
	"social-media-api/service"

	"social-media-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	controllers controller.Controllers
	svcs        service.Services
	repos       repo.Repositories
)

// @title S.M.S. API
// @version 1.0
// @description This is a simple social media scrapper
// @description
// @description NOTE: To get userID required for twitch endpoints(followers, following) use twtitch/{username}
// @description to get details for user. One of filed is ID you can use that value in before mentiond endpoints.

// @BasePath /api

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.Host(config.Conf.Host)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiSubrouter := router.PathPrefix("/api").Subrouter()

	initRepos()
	initServices()
	initControllers()
	setupRoutes(apiSubrouter)
	handleIncomingRequests(router)
}

func initRepos() {
	repo.InitRepos(&repos)
}

func initServices() {
	svcs.Init(&repos)
}

func initControllers() {
	controllers.Init(&svcs)
}

func setupRoutes(apiSubrouter *mux.Router) {
	routes.SetupTwitterRoutes(apiSubrouter, controllers.TwitterController, &svcs)
	routes.SetupTwitchRoutes(apiSubrouter, controllers.TwitchController)
}

func handleIncomingRequests(router *mux.Router) {
	done := make(chan error, 1)
	go func() {
		done <- http.ListenAndServe(fmt.Sprint(":", config.Conf.API_port), router)
	}()
	http.Handle("/", router)
	log.Printf("--- SERVICE STARTED ---")
	if err := <-done; err != nil {
		log.Printf("Failed to serve. Error message: %s", err)
		os.Exit(1)
	}
}

func init() {
	//load default configuration
	if err := config.LoadConfJson(); err != nil {
		log.Fatalf("Failed loading service config. Error message: %s", err)
	}
	// load environment variables
	config.LoadEnv()
	docs.SwaggerInfo.Host = config.Conf.Swagger.Host
}

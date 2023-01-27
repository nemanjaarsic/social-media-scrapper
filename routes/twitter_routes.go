package routes

import (
	"social-media-api/controller"
	"social-media-api/service"

	"github.com/gorilla/mux"
)

func SetupTwitterRoutes(r *mux.Router, ctrl *controller.TwitterController, svcs *service.Services) {
	routerRead := r.PathPrefix("/twitter").Subrouter()

	routerRead.HandleFunc("/{username}/followers", ctrl.GetFollowers).Methods("GET")
	routerRead.HandleFunc("/{username}/likes", ctrl.GetLikes).Methods("GET")
	routerRead.Use(authenticationMiddleware(svcs))
}

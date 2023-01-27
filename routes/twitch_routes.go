package routes

import (
	"social-media-api/controller"

	"github.com/gorilla/mux"
)

func SetupTwitchRoutes(r *mux.Router, ctrl *controller.TwitchController) {
	routerRead := r.PathPrefix("/twitch").Subrouter()

	routerRead.HandleFunc("/{username}", ctrl.GetUser).Methods("GET")
	routerRead.HandleFunc("/{username}/views", ctrl.GetUserViewsCount).Methods("GET")
	routerRead.HandleFunc("/{userID}/followers", ctrl.GetFollowers).Methods("GET")
	routerRead.HandleFunc("/{userID}/following", ctrl.GetFollowing).Methods("GET")
}

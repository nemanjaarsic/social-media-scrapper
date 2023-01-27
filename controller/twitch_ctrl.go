package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media-api/service"

	"github.com/gorilla/mux"
)

type TwitchController struct {
	TwitchService service.TwitchService
}

func NewTwitchController(svcs *service.Services) *TwitchController {
	ctrl := &TwitchController{
		TwitchService: svcs.TwitchService,
	}
	return ctrl
}

//	@Summary		Get user details
//	@Description	Get details for user with provided username
//	@Tags			twitch
//	@Accept			json
//	@Produce		json
//	@Param			username	path		string	true	"Twitch username"
//	@Success		200			{object}	model.TwitchUser
//	@Failure		400
//	@Failure		500
//	@Router			/twitch/{username} [get]
func (c *TwitchController) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := mux.Vars(r)["username"]
	if username == "" {
		msg := fmt.Sprint("Error [TwitchController, GetUser]", "Invalid username")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	user, err := c.TwitchService.GetUser(ctx, username)
	if err != nil {
		msg := fmt.Sprint("Error [TwitchController, GetUser]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(user)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitchController, GetUser]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//	@Summary		Show user views
//	@Description	This endpoint will retrive views count for specified user, if username exists along side username...
//	@Tags			twitch
//	@Accept			json
//	@Produce		json
//	@Param			username	path		string	true	"Twitch username"
//	@Success		200			{object}	model.TwitchViewsCount
//	@Failure		400
//	@Failure		500
//	@Router			/twitch/{username}/views [get]
func (c *TwitchController) GetUserViewsCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := mux.Vars(r)["username"]
	if username == "" {
		msg := fmt.Sprint("Error [TwitchController, GetUserViewsCount]", "Invalid username")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	viewsData, err := c.TwitchService.GetUserViewsCount(ctx, username)
	if err != nil {
		msg := fmt.Sprint("Error [TwitchController, GetUserViewsCount]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(viewsData)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitchController, GetUserViewsCount]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//	@Summary		Show user's followers
//	@Description	This endpoint will retrive list of followers of user with corresponding userId
//	@Tags			twitch
//	@Accept			json
//	@Produce		json
//	@Param			userID	path	string	true	"Twitch userID"
//	@Success		200		{array}	model.TwitchFollowers
//	@Failure		400
//	@Failure		500
//	@Router			/twitch/{userID}/followers [get]
func (c *TwitchController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := mux.Vars(r)["userID"]
	if userID == "" {
		msg := fmt.Sprint("Error [TwitchController, GetFollowers]", "Invalid username")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	followers, err := c.TwitchService.GetFollowers(ctx, userID)
	if err != nil {
		msg := fmt.Sprint("Error [TwitchController, GetFollowers]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(followers)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitchController, GetFollowers]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//	@Summary		Show list of users that user is following
//	@Description	This endpoint will retrive list of users that user with corresponding userId is following
//	@Tags			twitch
//	@Accept			json
//	@Produce		json
//	@Param			userID	path	string	true	"Twitch userID"
//	@Success		200		{array}	model.TwitchFollowing
//	@Failure		400
//	@Failure		500
//	@Router			/twitch/{userID}/following [get]
func (c *TwitchController) GetFollowing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := mux.Vars(r)["userID"]
	if userID == "" {
		msg := fmt.Sprint("Error [TwitchController, GetFollowing]", "Invalid username")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	followings, err := c.TwitchService.GetFollowing(ctx, userID)
	if err != nil {
		msg := fmt.Sprint("Error [TwitchController, GetFollowing]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(followings)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitchController, GetFollowing]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-media-api/service"

	"google.golang.org/grpc/metadata"
)

type TwitterController struct {
	TwitterService service.TwitterService
}

func NewTwitterController(svcs *service.Services) *TwitterController {
	ctrl := &TwitterController{
		TwitterService: svcs.TwitterService,
	}
	return ctrl
}

// @Summary		Show list of followers
// @Description	This endpoint will retrive followers for specified user, if username exists
// @Tags			twitter
// @Accept			json
// @Produce		json
// @Param			username	path	string	true	"Twitter username (case sensitive)"
// @Success		200			{array}	model.TwitterUserFollowers
// @Failure		400
// @Failure		500
// @Router			/twitter/{username}/followers [get]
func (c *TwitterController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	md, _ := metadata.FromOutgoingContext(ctx)
	userIDs := md.Get("userID")
	if len(userIDs) == 0 {
		log.Print("[TwitterController, GetFollowers] User id is missing")
		return
	}
	userID := userIDs[0]

	followers, err := c.TwitterService.GetFallowers(ctx, userID)
	if err != nil {
		msg := fmt.Sprint("Error [TwitterController, GetFollowers]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(followers)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitterController, GetFollowers]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// @Summary		Show list of likes
// @Description	This endpoint will retrive likes for specified user, if username exists
// @Tags			twitter
// @Accept			json
// @Produce		json
// @Param			username	path	string	true	"Twitter username (case sensitive)"
// @Success		200			{array}	model.TwitterUserLikes
// @Failure		400
// @Failure		500
// @Router			/twitter/{username}/likes [get]
func (c *TwitterController) GetLikes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	md, _ := metadata.FromOutgoingContext(ctx)
	userIDs := md.Get("userID")
	if len(userIDs) == 0 {
		log.Print("[TwitterController, GetLikes] User id is missing")
		return
	}
	userID := userIDs[0]

	likes, err := c.TwitterService.GetLikes(ctx, userID)
	if err != nil {
		msg := fmt.Sprint("Error [TwitterController, GetLikes]", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response, mErr := json.Marshal(likes)
	if mErr != nil {
		msg := fmt.Sprint("Error [TwitterController, GetLikes]", mErr.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

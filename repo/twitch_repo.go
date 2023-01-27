package repo

import (
	"context"
	"fmt"
	"net/http"
	"social-media-api/config"
	"social-media-api/model"
	twitch "social-media-api/twitch_client"
)

type twitchRepo struct {
	client *twitch.Client
}

// Check interface compliance
var _ TwitchRepository = (*twitchRepo)(nil)

func NewTwitchRepo() *twitchRepo {
	client := &twitch.Client{
		Authorizer: authorize{
			Token: config.Conf.Twitch.Token,
		},
		ClientID: config.Conf.Twitch.ClientID,
		Client:   http.DefaultClient,
		Host:     config.Conf.Twitch.Host,
	}
	return &twitchRepo{
		client: client,
	}
}

func (r *twitchRepo) GetUser(ctx context.Context, username string) (model.TwitchUser, error) {
	resp, err := r.client.GetUserByUsername(ctx, username)
	if err != nil {
		return model.TwitchUser{}, err
	}

	return resp.Data[0].MapResponseToUser(), nil
}

func (r *twitchRepo) GetUserViewsCount(ctx context.Context, username string) (model.TwitchViewsCount, error) {
	resp, err := r.client.GetUserByUsername(ctx, username)
	if err != nil {
		return model.TwitchViewsCount{}, err
	}

	//Make shure there is only one user in response
	//if lenght of 'raw.Data' is 0 it means that there are no users that match to provided username
	//if lenght of 'raw.Data' is more than one something is not quite right with response
	if len(resp.Data) == 0 || len(resp.Data) > 1 {
		return model.TwitchViewsCount{}, fmt.Errorf("[TwitchRepo GetUserViewsCount] Error: User with provided username does not exist")
	}
	return resp.Data[0].MapResponseToViewCount(), nil
}

func (r *twitchRepo) GetFollowers(ctx context.Context, userID string) (model.TwitchFollowers, error) {
	resp, err := r.client.GetFollowersTo(ctx, userID)
	if err != nil {
		return model.TwitchFollowers{}, err
	}
	res := model.TwitchFollowers{
		Total: resp.Total,
		Users: make([]model.TwitchFollower, 0),
	}
	for _, u := range resp.Data {
		res.Users = append(res.Users, u.MapResponseUserFollowers())
	}
	return res, nil
}

func (r *twitchRepo) GetFollowing(ctx context.Context, userID string) (model.TwitchFollowing, error) {
	resp, err := r.client.GetFollowersFrom(ctx, userID)
	if err != nil {
		return model.TwitchFollowing{}, err
	}
	res := model.TwitchFollowing{
		Total: resp.Total,
		Users: make([]model.TwitchFollower, 0),
	}
	for _, u := range resp.Data {
		res.Users = append(res.Users, u.MapResponseUserFollowsTo())
	}
	return res, nil
}

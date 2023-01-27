package repo

import (
	"context"
	"fmt"
	"net/http"
	"social-media-api/config"
	"social-media-api/model"

	"github.com/g8rswimmer/go-twitter/v2"
)

type twitterRepo struct {
	client *twitter.Client
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

// Check interface compliance
var _ TwitterRepository = (*twitterRepo)(nil)

func NewTwitterRepo() *twitterRepo {
	client := &twitter.Client{
		Authorizer: authorize{
			Token: config.Conf.Twitter.Token,
		},
		Client: http.DefaultClient,
		Host:   config.Conf.Twitter.Host,
	}
	return &twitterRepo{
		client: client,
	}
}

func (r *twitterRepo) AuthenticateUser(ctx context.Context, username string) (string, error) {
	usernames := []string{username}
	users, err := r.client.UserNameLookup(ctx, usernames, twitter.UserLookupOpts{})
	if err != nil {
		return "", err
	}
	var userID string
	if users.Raw.Users[0] == nil {
		return "", fmt.Errorf("invalid username")
	}

	for _, u := range users.Raw.Users {
		if u.UserName == username {
			userID = u.ID
		}
	}
	return userID, nil
}

func (r *twitterRepo) GetFollowers(ctx context.Context, userID string) ([]model.TwitterUserFollowers, error) {
	opts := twitter.UserFollowersLookupOpts{
		UserFields: []twitter.UserField{twitter.UserFieldDescription},
		MaxResults: 10,
	}

	followers, err := r.client.UserFollowersLookup(ctx, userID, opts)
	if err != nil {
		return []model.TwitterUserFollowers{}, err
	}

	return mapGetFallowersResponse(followers), nil
}

func mapGetFallowersResponse(followers *twitter.UserFollowersLookupResponse) []model.TwitterUserFollowers {
	result := make([]model.TwitterUserFollowers, 0)
	for _, u := range followers.Raw.Users {
		result = append(result, model.TwitterUserFollowers{
			ID:          u.ID,
			Name:        u.Name,
			UserName:    u.UserName,
			Description: u.Description,
		})
	}
	return result
}

func (r *twitterRepo) GetLikes(ctx context.Context, userID string) ([]model.TwitterUserLikes, error) {
	opts := twitter.UserLikesLookupOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldAuthorID},
		MaxResults:  10,
	}
	likes, err := r.client.UserLikesLookup(ctx, userID, opts)
	if err != nil {
		return []model.TwitterUserLikes{}, err
	}

	return mapGetLikesResponse(likes), err
}

func mapGetLikesResponse(likes *twitter.UserLikesLookupResponse) []model.TwitterUserLikes {
	result := make([]model.TwitterUserLikes, 0)
	for _, t := range likes.Raw.Tweets {
		result = append(result, model.TwitterUserLikes{
			ID:       t.ID,
			Text:     t.Text,
			AuthorID: t.AuthorID,
		})
	}
	return result
}

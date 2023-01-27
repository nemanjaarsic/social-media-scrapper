package repo

import (
	"context"
	"social-media-api/model"
)

type TwitterRepository interface {
	AuthenticateUser(context.Context, string) (string, error)
	GetFollowers(context.Context, string) ([]model.TwitterUserFollowers, error)
	GetLikes(context.Context, string) ([]model.TwitterUserLikes, error)
}

type TwitchRepository interface {
	GetUser(context.Context, string) (model.TwitchUser, error)
	GetUserViewsCount(context.Context, string) (model.TwitchViewsCount, error)
	GetFollowers(context.Context, string) (model.TwitchFollowers, error)
	GetFollowing(context.Context, string) (model.TwitchFollowing, error)
}

type Repositories struct {
	TwitterRepo TwitterRepository
	TwitchRepo  TwitchRepository
}

func InitRepos(repo *Repositories) {
	repo.TwitterRepo = NewTwitterRepo()
	repo.TwitchRepo = NewTwitchRepo()
}

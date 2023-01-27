package service

import (
	"context"
	"social-media-api/model"
	"social-media-api/repo"
)

type TwitchService struct {
	twitchRepo repo.TwitchRepository
}

func (s *TwitchService) Init(repo *repo.Repositories) {
	s.twitchRepo = repo.TwitchRepo
}

func (s *TwitchService) GetUser(ctx context.Context, username string) (model.TwitchUser, error) {
	return s.twitchRepo.GetUser(ctx, username)
}

func (s *TwitchService) GetUserViewsCount(ctx context.Context, username string) (model.TwitchViewsCount, error) {
	return s.twitchRepo.GetUserViewsCount(ctx, username)
}

func (s *TwitchService) GetFollowers(ctx context.Context, userID string) (model.TwitchFollowers, error) {
	return s.twitchRepo.GetFollowers(ctx, userID)
}

func (s *TwitchService) GetFollowing(ctx context.Context, userID string) (model.TwitchFollowing, error) {
	return s.twitchRepo.GetFollowing(ctx, userID)
}

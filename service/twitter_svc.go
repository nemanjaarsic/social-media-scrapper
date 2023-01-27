package service

import (
	"context"
	"social-media-api/model"
	"social-media-api/repo"
)

type TwitterService struct {
	twitterRepo repo.TwitterRepository
}

func (s *TwitterService) Init(repo *repo.Repositories) {
	s.twitterRepo = repo.TwitterRepo
}

func (s *TwitterService) AuthenticateUser(ctx context.Context, username string) (string, error) {
	return s.twitterRepo.AuthenticateUser(ctx, username)
}

func (s *TwitterService) GetFallowers(ctx context.Context, userID string) ([]model.TwitterUserFollowers, error) {
	return s.twitterRepo.GetFollowers(ctx, userID)
}

func (s *TwitterService) GetLikes(ctx context.Context, userID string) ([]model.TwitterUserLikes, error) {
	return s.twitterRepo.GetLikes(ctx, userID)
}

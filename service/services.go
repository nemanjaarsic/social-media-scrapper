package service

import "social-media-api/repo"

type Services struct {
	TwitterService TwitterService
	TwitchService  TwitchService
}

func (svcs *Services) Init(repos *repo.Repositories) {
	svcs.TwitterService.Init(repos)
	svcs.TwitchService.Init(repos)
}

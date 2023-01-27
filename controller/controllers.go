package controller

import "social-media-api/service"

type Controllers struct {
	TwitterController *TwitterController
	TwitchController  *TwitchController
}

func (c *Controllers) Init(svcs *service.Services) {
	c.TwitterController = NewTwitterController(svcs)
	c.TwitchController = NewTwitchController(svcs)
}

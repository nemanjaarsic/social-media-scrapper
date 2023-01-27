package model

type TwitchFollowersResponse struct {
	Total int
	Data  []FollowersDataResponse
}

type FollowersDataResponse struct {
	From_id     string
	From_login  string
	From_name   string
	To_id       string
	To_login    string
	To_name     string
	Followed_at string
}

type TwitchFollowers struct {
	Total int              `json:"Total" example:"1200"`
	Users []TwitchFollower `json:"Users"`
}

type TwitchFollowing struct {
	Total int              `json:"Total" example:"120"`
	Users []TwitchFollower `json:"Users"`
}

type TwitchFollower struct {
	ID         string `json:"ID" example:"12342342"`
	Name       string `json:"Name" example:"aCoolUsername"`
	FollowedAt string `json:"FollowedAt" example:"2023-01-26T19:34:28Z"`
}

func (f *FollowersDataResponse) MapResponseUserFollowers() TwitchFollower {
	return TwitchFollower{
		ID:         f.From_id,
		Name:       f.From_name,
		FollowedAt: f.Followed_at,
	}
}

func (f *FollowersDataResponse) MapResponseUserFollowsTo() TwitchFollower {
	return TwitchFollower{
		ID:         f.To_id,
		Name:       f.To_name,
		FollowedAt: f.Followed_at,
	}
}

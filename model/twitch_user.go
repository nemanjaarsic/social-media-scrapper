package model

type TwitchUserResponse struct {
	Data []TwitchUserDataResponse
}

type TwitchUserDataResponse struct {
	ID                string
	Login             string
	Display_name      string
	Type              string
	Broadcaster_type  string
	Description       string
	Profile_image_url string
	Offline_image_url string
	View_count        int
	Created_at        string
}

type TwitchViewsCount struct {
	ID    string `json:"ID" example:"12342342"`
	Name  string `json:"Name" example:"aCoolUsername"`
	Views int    `json:"Views" example:"6302"`
}

type TwitchUser struct {
	ID          string `json:"ID" example:"12342342"`
	Name        string `json:"Name" example:"aCoolUsername"`
	Description string `json:"Description" example:"Some description text"`
	CreatedAt   string `json:"Created at" example:"2022-01-26T19:34:28Z"`
}

func (u *TwitchUserDataResponse) MapResponseToUser() TwitchUser {
	return TwitchUser{
		ID:          u.ID,
		Name:        u.Display_name,
		Description: u.Description,
		CreatedAt:   u.Created_at,
	}
}

func (u *TwitchUserDataResponse) MapResponseToViewCount() TwitchViewsCount {
	return TwitchViewsCount{
		ID:    u.ID,
		Name:  u.Display_name,
		Views: u.View_count,
	}
}

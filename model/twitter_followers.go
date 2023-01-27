package model

type TwitterUserFollowers struct {
	ID          string `json:"ID" example:"12342342"`
	Name        string `json:"Name" example:"John"`
	UserName    string `json:"UserName" example:"userNameJohn"`
	Description string `json:"Description" example:"This is some lame description"`
}

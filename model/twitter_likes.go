package model

type TwitterUserLikes struct {
	ID       string `json:"ID" example:"12342342"`
	Text     string `json:"Text" example:"Some tweet text"`
	AuthorID string `json:"AuthorID" example:"3454624643"`
}

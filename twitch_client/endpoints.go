package twitch

import (
	"fmt"
	"strings"
)

type endpoint string

const (
	userViewsData endpoint = "helix/users?login={username}"
	getFollowers  endpoint = "helix/users/follows?to_id={id}"
	getFollowings endpoint = "helix/users/follows?from_id={id}"

	idTag       = "{id}"
	usernameTag = "{username}"
)

func (e endpoint) url(host string) string {
	return fmt.Sprintf("%s/%s", host, string(e))
}

func (e endpoint) urlID(host, id string) string {
	u := fmt.Sprintf("%s/%s", host, string(e))
	return strings.ReplaceAll(u, idTag, id)
}

func (e endpoint) urlUsername(host, username string) string {
	u := fmt.Sprintf("%s/%s", host, string(e))
	return strings.ReplaceAll(u, usernameTag, username)
}

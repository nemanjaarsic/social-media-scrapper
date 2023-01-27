package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"social-media-api/model"
)

type Authorizer interface {
	Add(req *http.Request)
}

type Client struct {
	Authorizer Authorizer
	ClientID   string
	Client     *http.Client
	Host       string
}

func (c *Client) GetUserByUsername(ctx context.Context, username string) (*model.TwitchUserResponse, error) {
	ep := userViewsData.urlUsername(c.Host, username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return &model.TwitchUserResponse{}, fmt.Errorf("[TwitchClient GetUserByUsername] Error: %s", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("client-id", c.ClientID)
	c.Authorizer.Add(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return &model.TwitchUserResponse{}, fmt.Errorf("[TwitchClient GetUserByUsername] Error: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &model.TwitchUserResponse{}, fmt.Errorf("[TwitchClient GetUserByUsername] Error: %s", resp.Status)
	}

	raw := &model.TwitchUserResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&raw); err != nil {
		return &model.TwitchUserResponse{}, fmt.Errorf("[TwitchClient GetUserByUsername] Decoder error: %s", err.Error())
	}
	return raw, nil
}

func (c *Client) GetFollowersTo(ctx context.Context, userID string) (*model.TwitchFollowersResponse, error) {
	ep := getFollowers.urlID(c.Host, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersTo] Error: %s", err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("client-id", c.ClientID)
	c.Authorizer.Add(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersTo] Error: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersTo] Error: %s", resp.Status)
	}

	raw := &model.TwitchFollowersResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&raw); err != nil {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersTo] Decoder error: %s", err.Error())
	}
	return raw, nil
}

func (c *Client) GetFollowersFrom(ctx context.Context, userID string) (*model.TwitchFollowersResponse, error) {
	ep := getFollowings.urlID(c.Host, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("[TwitchClient GetFollowersFrom] Error: %s", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("client-id", c.ClientID)
	c.Authorizer.Add(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[TwitchClient GetFollowersFrom] Error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersFrom] Error: %s", resp.Status)
	}

	raw := &model.TwitchFollowersResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&raw); err != nil {
		return &model.TwitchFollowersResponse{}, fmt.Errorf("[TwitchClient GetFollowersFrom] Decoder error: %s", err.Error())
	}
	return raw, nil
}

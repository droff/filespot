package filespot

import (
	"context"
	"net/http"
)

const playersBasePath = "/1/players"

// PlayersService implements interface with API /players endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#players
type PlayersService interface {
	List(context.Context, interface{}) (*playersRoot, *http.Response, error)
	Get(context.Context, string) (*Player, *http.Response, error)
	Create(context.Context, *PlayerCreateRequest) (*Player, *http.Response, error)
	Update(context.Context, string, *PlayerUpdateRequest) (*http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

// PlayersCli handles communication with API
type PlayersCli struct {
	client *Client
}

// Player represents a platformcraft Player
type Player struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	IsDir         bool   `json:"is_dir"`
	Videos        Videos `json:"videos"`
	ScreenShotURL string `json:"screen_shot_url"`
	VastAdTagURL  string `json:"vast_ad_tag_url"`
	CreateDate    string `json:"create_date"`
	Href          string `json:"href"`
	FrameTag      string `json:"frame_tag"`
	Description   string `json:"description"`
	Tags          Tags   `json:"tags"`
	Geo           Geo    `json:"geo"`
}

type Videos map[string]string
type Tags []string

// playersRoot respresents a List root
type playersRoot struct {
	Players []Player `json:"players"`
}

// playerRoot represents a Get root
type playerRoot struct {
	Player *Player `json:"player"`
}

// PlayerCreateRequest identifies Player for the Create request
type PlayerCreateRequest struct {
	Name         string `json:"name"`
	Folder       string `json:"folder"`
	Videos       Videos `json:"videos"`
	ScreenShotID string `json:"screen_shot_id"`
	VastAdTagURL string `json:"vast_ad_tag_url"`
	Description  string `json:"description"`
	Tags         Tags   `json:"tags"`
	Geo          Geo    `json:"geo"`
}

// PlayerUpdateRequest indentifies Player for the Update request
type PlayerUpdateRequest struct {
	Name         string `json:"name"`
	Folder       string `json:"folder"`
	Videos       Videos `json:"videos"`
	ScreenShotID string `json:"screen_shot_id"`
	Description  string `json:"description"`
	Tags         Tags   `json:"tags"`
	Geo          Geo    `json:"geo"`
}

// List of Players
func (c PlayersCli) List(ctx context.Context, params interface{}) (*playersRoot, *http.Response, error) {
	path, err := addParams(playersBasePath, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(playersRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// Get Player
func (c PlayersCli) Get(ctx context.Context, id string) (*Player, *http.Response, error) {
	endpointURL := playersBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(playerRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Player, resp, err
}

// Create Player
func (c PlayersCli) Create(ctx context.Context, playerCreateRequest *PlayerCreateRequest) (*Player, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, playersBasePath, playerCreateRequest)
	if err != nil {
		return nil, nil, err
	}

	data := new(playerRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Player, resp, err
}

// Update Player
func (c PlayersCli) Update(ctx context.Context, id string, playerUpdateRequest *PlayerUpdateRequest) (*http.Response, error) {
	endpointURL := playersBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPut, endpointURL, playerUpdateRequest)
	if err != nil {
		return nil, err
	}

	data := &struct{}{}
	resp, err := c.client.Do(ctx, req, data)

	return resp, err
}

// Delete Player
func (c PlayersCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := playersBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodDelete, endpointURL, nil)
	if err != nil {
		return nil, err
	}

	data := &struct{}{}
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return resp, err
	}

	return resp, err
}

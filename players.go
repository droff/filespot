package filespot

import (
	"context"
	"net/http"
)

const playersBasePath = "/1/players"

// PlayersService implements interface with API /players endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#players
type PlayersService interface {
	Create(context.Context, *PlayerCreateRequest) (*Player, *http.Response, error)
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

package filespot

import (
	"context"
	"net/http"
)

const objectsBasePath = "/1/objects"

type ObjectsService interface {
	List(context.Context) ([]Object, *http.Response, error)
}

type ObjectsCli struct {
	client *Client
}

type Object struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	IsDir        bool   `json:"is_dir"`
	Size         uint32 `json:"size"`
	ContentType  string `json:"content_type"`
	CreateDate   string `json:"create_date"`
	LatestUpdate string `json:"latest_update"`
	ResourceURL  string `json:"resource_url"`
	CDNURL       string `json:"cdn_url"`
	VODHLS       string `json:"vod_hls"`
	Video        string `json:"video"`
	Private      bool   `json:"private"`
	Status       string `json:"status"`
}

type objectsRoot struct {
	Objects []Object `json:"objects"`
}

func (c ObjectsCli) List(ctx context.Context) ([]Object, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, objectsBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(objectsRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, nil, err
	}

	return data.Objects, resp, err
}

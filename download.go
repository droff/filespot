package filespot

import (
	"context"
	"net/http"
)

const downloadBasePath = "/1/download"

// DownloadService implements interface with API /download endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#download
type DownloadService interface {
	Create(context.Context, interface{}) (*Download, *http.Response, error)
}

// DownloadCli handles communication with API
type DownloadCli struct {
	client *Client
}

// Download represents a platformcraft Download
type Download struct {
	Message     string `json:"message"`
	ActiveTasks int    `json:"active_tasks"`
	TaskID      string `json:"task_id"`
}

// DownloadCreateParams identifies as query params of Create request
type DownloadCreateParams struct {
	URL          string `url:"url,omitempty"`
	Path         string `url:"path,omitempty"`
	Name         string `url:"name,omitempty"`
	Autoencoding bool   `url:"autoencoding,omitempty"`
	Presets      string `url:"presets,omitempty"`
	DelOriginal  bool   `url:"del_original,omitempty"`
	Autoplayer   bool   `url:"autoplayer,omitempty"`
}

// Create Download
func (c DownloadCli) Create(ctx context.Context, params interface{}) (*Download, *http.Response, error) {
	path, err := addParams(downloadBasePath, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(Download)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

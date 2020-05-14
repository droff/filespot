package filespot

import (
	"context"
	"net/http"
)

const transcoderTasksBasePath = "/1/transcoder_tasks"

// TranscoderTasksService implements interface with API /transcoder_tasks endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#transcoder_tasks
type TranscoderTasksService interface {
	List(context.Context) (*tasksRoot, *http.Response, error)
	Get(context.Context, string) (*Task, *http.Response, error)
	HLS(context.Context, string) (*Task, *http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

// TranscoderTasksCli handles communication with API
type TranscoderTasksCli struct {
	client *Client
}

// List returns list of all transcoders tasks
func (c TranscoderTasksCli) List(ctx context.Context) (*tasksRoot, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, transcoderTasksBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(tasksRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// Get Task
func (c TranscoderTasksCli) Get(ctx context.Context, id string) (*Task, *http.Response, error) {
	endpointURL := transcoderTasksBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(taskRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Task, resp, err
}

// HLS Task
func (c TranscoderTasksCli) HLS(ctx context.Context, id string) (*Task, *http.Response, error) {
	endpointURL := transcoderTasksBasePath + "/hls/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(taskRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Task, resp, err
}

// Delete transcoder Task
func (c TranscoderTasksCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := transcoderTasksBasePath + "/" + id

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

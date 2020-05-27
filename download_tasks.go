package filespot

import (
	"context"
	"net/http"
)

const downloadTasksBasePath = "/1/download_tasks"

// DownloadTasksService implements interface with API /download_tasks endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#download_tasks
type DownloadTasksService interface {
	List(context.Context) (*tasksRoot, *http.Response, error)
	Get(context.Context, string) (*Task, *http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

// DownloadTasksCli handles communication with API
type DownloadTasksCli struct {
	client *Client
}

// Task represents a platformcraft Task
type Task struct {
	ID         string `json:"id"`
	Category   string `json:"category"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Status     string `json:"status"`
	TimeStart  string `json:"time_start"`
	TimeFinish string `json:"time_finish"`
	Lock       bool   `json:"lock"`
}

// tasksRoot represents a List root
type tasksRoot struct {
	Tasks []Task `json:"tasks"`
	Count int    `json:"count"`
}

// taskRoot represents a Get root
type taskRoot struct {
	Task *Task `json:"task"`
}

// List of Tasks
func (c DownloadTasksCli) List(ctx context.Context) (*tasksRoot, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, downloadTasksBasePath, nil)
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
func (c DownloadTasksCli) Get(ctx context.Context, id string) (*Task, *http.Response, error) {
	endpointURL := downloadTasksBasePath + "/" + id

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

// Delete Task
func (c DownloadTasksCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := downloadTasksBasePath + "/" + id

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

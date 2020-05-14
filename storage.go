package filespot

import (
	"context"
	"net/http"
)

const storageBasePath = "/1/storage"

// StorageService implements interface with API /storage endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#storage
type StorageService interface {
	Get(context.Context) (*Storage, *http.Response, error)
}

// StorageCli handles communication with API
type StorageCli struct {
	client *Client
}

// Storage represents a platformcraft Storage
type Storage struct {
	Used  int `json:"used"`
	Limit int `json:"limit"`
}

// storageRoot represents a Get root
type storageRoot struct {
	Storage *Storage `json:"storage"`
}

// Get Storage
func (c StorageCli) Get(ctx context.Context) (*Storage, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, storageBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(storageRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Storage, resp, err
}

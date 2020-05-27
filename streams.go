package filespot

import (
	"context"
	"net/http"
)

const streamsBasePath = "/1/streams"

// StreamsService implements interface with API /streams endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#streams
type StreamsService interface {
	List(context.Context) ([]Stream, *http.Response, error)
	Get(context.Context, string) (*Stream, *http.Response, error)
	Create(context.Context, *StreamCreateRequest) (*Stream, *http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
	Start(context.Context, string, *StreamStartRequest) (*http.Response, error)
	Stop(context.Context, string) ([]File, *http.Response, error)
	CreateSchedule(context.Context, string) (string, *http.Response, error)
	Rec(context.Context, string) (*Record, *http.Response, error)
	DeleteSchedule(context.Context, string, string) (*http.Response, error)
}

// StreamsCli handles communication with API
type StreamsCli struct {
	client *Client
}

// Stream represents a platformcraft Stream
type Stream struct {
	ID                 string `json:"id"`
	User               string `json:"user"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	IsInstantRecording bool   `json:"is_instant_recording"`
}

// File represents a platformcraft Record File
type File struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int    `json:"size"`
	ContentType  string `json:"content_type"`
	CreateDate   string `json:"create_date"`
	LatestUpdate string `json:"latest_update"`
	ResourceURL  string `json:"resource_url"`
	Video        string `json:"video"`
	CDNURL       string `json:"cdn_url"`
	Status       string `json:"status"`
}

// Record represents a platformcraft Record
type Record struct {
	Status string   `json:"status"`
	Files  []string `json:"files"`
}

// streamsRoot represents a List root
type streamsRoot struct {
	Streams []Stream `json:"streams"`
}

// streamRoot represents a Get root
type streamRoot struct {
	Stream *Stream `json:"stream"`
}

// stopRoot represents a Stop root
type stopRoot struct {
	Files []File `json:"files"`
}

// recordRoot represents a Rec root
type recordRoot struct {
	Record *Record `json:"record"`
}

// StreamCreateRequest identifies Stream for the Create request
type StreamCreateRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// StreamStartRequest identifies streaming process for the Start request
type StreamStartRequest struct {
	StopTimeout int `json:"stop_timeout"`
}

// List of Streams
func (c StreamsCli) List(ctx context.Context) ([]Stream, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, streamsBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(streamsRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Streams, resp, err
}

// Get Stream
func (c StreamsCli) Get(ctx context.Context, id string) (*Stream, *http.Response, error) {
	endpointURL := streamsBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(streamRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Stream, resp, err
}

// Create Stream
func (c StreamsCli) Create(ctx context.Context, streamCreateRequest *StreamCreateRequest) (*Stream, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, streamsBasePath, streamCreateRequest)
	if err != nil {
		return nil, nil, err
	}

	data := new(streamRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Stream, resp, err
}

// Delete Stream
func (c StreamsCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := streamsBasePath + "/" + id

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

// Start Stream
func (c StreamsCli) Start(ctx context.Context, id string, streamStartRequest *StreamStartRequest) (*http.Response, error) {
	endpointURL := streamsBasePath + "/rec/instant/start/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, streamStartRequest)
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

// Stop Stream
func (c StreamsCli) Stop(ctx context.Context, id string) ([]File, *http.Response, error) {
	endpointURL := streamsBasePath + "/rec/instant/stop/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(stopRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Files, resp, err
}

// CreateSchedule returns record_id
func (c StreamsCli) CreateSchedule(ctx context.Context, id string) (string, *http.Response, error) {
	endpointURL := streamsBasePath + "/rec/schedule/new/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return "", nil, err
	}

	data := &struct {
		RecordID string `json:"record_id"`
	}{}
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return "", resp, err
	}

	return data.RecordID, resp, err
}

// Rec returns Record
func (c StreamsCli) Rec(ctx context.Context, id string) (*Record, *http.Response, error) {
	endpointURL := streamsBasePath + "/rec/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(recordRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Record, resp, err
}

// DeleteSchedule deletes record
func (c StreamsCli) DeleteSchedule(ctx context.Context, streamID string, recordID string) (*http.Response, error) {
	endpointURL := streamsBasePath + "/rec/schedule/del/" + streamID + "/" + recordID

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

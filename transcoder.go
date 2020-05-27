package filespot

import (
	"context"
	"net/http"
)

const transcoderBasePath = "/1/transcoder"

// TranscoderService implements interface with API /transcoder endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#transcoder
type TranscoderService interface {
	Presets(context.Context) (*presetsRoot, *http.Response, error)
	Create(context.Context, string, *TranscoderCreateRequest) (*Transcoder, *http.Response, error)
	Concat(context.Context, *TranscoderConcatRequest) (*Transcoder, *http.Response, error)
	HLS(context.Context, string, *TranscoderHLSRequest) (*Transcoder, *http.Response, error)
}

// TranscoderCli handles communication with API
type TranscoderCli struct {
	client *Client
}

// Transcoder represents a platformcraft Transcoder
type Transcoder struct {
	ActiveTasks int    `json:"active_tasks"`
	TaskID      string `json:"task_id"`
}

// Preset represents a platformcraft Preset
type Preset struct {
	ID         string                     `json:"id"`
	Name       string                     `json:"name"`
	Container  string                     `json:"container"`
	Video      map[string]string          `json:"video"`
	Audio      map[string]string          `json:"audio"`
	Watermarks map[string]WatermarkParams `json:"watermarks"`
}

// presetsRoot represents a Presets root
type presetsRoot struct {
	Presets []Preset `json:"presets"`
	Count   int      `json:"count"`
}

// TranscoderCreateRequest identifies params for the Create request
type TranscoderCreateRequest struct {
	Presets     []string  `json:"presets"`
	Path        string    `json:"path"`
	Watermarks  Watermark `json:"watermarks"`
	DelOriginal bool      `json:"del_original"`
	Start       int       `json:"start"`
	Duration    int       `json:"duration"`
}

// TranscoderConcatRequest identifies params for the Concat request
type TranscoderConcatRequest struct {
	Files []string `json:"files"`
	Path  string   `json:"path"`
	Name  string   `json:"name"`
}

// TranscoderHLSRequest identifies params for the HLS request
type TranscoderHLSRequest struct {
	Presets         []string `json:"presets"`
	SegmentDuration int      `json:"segment_duration"`
}

type Watermark map[string]string
type WatermarkParams map[string]string

// Presets Transcoder
func (c TranscoderCli) Presets(ctx context.Context) (*presetsRoot, *http.Response, error) {
	endpointURL := transcoderBasePath + "/presets"

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(presetsRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// Create Transcoder
func (c TranscoderCli) Create(ctx context.Context, id string, transcoderCreateRequest *TranscoderCreateRequest) (*Transcoder, *http.Response, error) {
	endpointURL := transcoderBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, transcoderCreateRequest)
	if err != nil {
		return nil, nil, err
	}

	trancoder := new(Transcoder)
	resp, err := c.client.Do(ctx, req, trancoder)
	if err != nil {
		return nil, resp, err
	}

	return trancoder, resp, err
}

// Concat Transcoder
func (c TranscoderCli) Concat(ctx context.Context, transcoderConcatRequest *TranscoderConcatRequest) (*Transcoder, *http.Response, error) {
	endpointURL := transcoderBasePath + "?concat"

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, transcoderConcatRequest)
	if err != nil {
		return nil, nil, err
	}

	trancoder := new(Transcoder)
	resp, err := c.client.Do(ctx, req, trancoder)
	if err != nil {
		return nil, resp, err
	}

	return trancoder, resp, err
}

// HLS Transcoder
func (c TranscoderCli) HLS(ctx context.Context, id string, transcoderHLSRequest *TranscoderHLSRequest) (*Transcoder, *http.Response, error) {
	endpointURL := transcoderBasePath + "/hls/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, transcoderHLSRequest)
	if err != nil {
		return nil, nil, err
	}

	trancoder := new(Transcoder)
	resp, err := c.client.Do(ctx, req, trancoder)
	if err != nil {
		return nil, resp, err
	}

	return trancoder, resp, err
}

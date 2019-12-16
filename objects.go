package filespot

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const objectsBasePath = "/1/objects"

type ObjectsService interface {
	List(context.Context, interface{}) ([]Object, *http.Response, error)
	Get(context.Context, string) (*Object, *http.Response, error)
	Create(context.Context, *ObjectCreateRequest) (*Object, *http.Response, error)
	Update(context.Context, string, *ObjectUpdateRequest) (*http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

type ObjectsCli struct {
	client *Client
}

type Object struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Path         string          `json:"path"`
	IsDir        bool            `json:"is_dir"`
	Size         uint32          `json:"size"`
	ContentType  string          `json:"content_type"`
	CreateDate   string          `json:"create_date"`
	LatestUpdate string          `json:"latest_update"`
	ResourceURL  string          `json:"resource_url"`
	CDNURL       string          `json:"cdn_url"`
	VODHLS       string          `json:"vod_hls"`
	Video        string          `json:"video"`
	Private      bool            `json:"private"`
	Status       string          `json:"status"`
	Advanced     *ObjectAdvanced `json:"advanced"`
	Previews     []string        `json:"previews"`
	Description  string          `json:"description"`
}

type ObjectAdvanced struct {
	AudioStreams []ObjectAudioStream `json:"audio_streams"`
	Format       *ObjectFormat       `json:"format"`
	VideoStreams []ObjectVideoStream `json:"video_streams"`
}

type ObjectAudioStream struct {
	BitRate       uint32  `json:"bit_rate"`
	ChannelLayout string  `json:"channel_layout"`
	Channels      uint32  `json:"channels"`
	CodecLongName string  `json:"codec_long_name"`
	CodecType     string  `json:"codec_type"`
	Duration      float32 `json:"duration"`
	Index         uint32  `json:"index"`
	SampleRate    uint32  `json:"sample_rate"`
}

type ObjectFormat struct {
	BitRate        uint32  `json:"bit_rate"`
	Duration       float32 `json:"duration"`
	FormatLongName string  `json:"format_long_name"`
	FormatName     string  `json:"format_name"`
	NBStreams      uint32  `json:"nb_streams"`
}

type ObjectVideoStream struct {
	BitRate            uint32  `json:"bit_rate"`
	CodecName          string  `json:"codec_name"`
	CodecType          string  `json:"codec_type"`
	CodecLongName      string  `json:"codeclongname"`
	DisplayAspectRatio string  `json:"display_aspect_ratio"`
	Duration           float32 `json:"duration"`
	FPS                float32 `json:"fps"`
	Height             uint32  `json:"height"`
	Index              uint32  `json:"index"`
	Width              uint32  `json:"width"`
}

type objectsRoot struct {
	Objects []Object `json:"objects"`
}

type objectRoot struct {
	Object *Object `json:"object"`
}

type ObjectCreateRequest struct {
	File         string
	Name         string
	Private      bool
	Autoencoding bool
	Presets      string
	DelOriginal  bool
	Autoplayer   bool
}

type ObjectUpdateRequest struct {
	Name        string `json:"name"`
	Folder      string `json:"folder"`
	Description string `json:"description"`
	MaxHeight   int    `json:"max_height"`
	MaxWidth    int    `json:"max_width"`
	Private     bool   `json:"private"`
}

func (c ObjectsCli) List(ctx context.Context, params interface{}) ([]Object, *http.Response, error) {
	path, err := addParams(objectsBasePath, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(objectsRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Objects, resp, err
}

func (c ObjectsCli) Get(ctx context.Context, id string) (*Object, *http.Response, error) {
	endpointURL := objectsBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(objectRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Object, resp, err
}

func (c ObjectsCli) Create(ctx context.Context, objectCreateRequest *ObjectCreateRequest) (*Object, *http.Response, error) {
	var fw io.Writer
	var err error
	method := http.MethodPost
	u := c.client.requestURL(method, objectsBasePath)
	buf := new(bytes.Buffer)
	mp := multipart.NewWriter(buf)

	fw, err = mp.CreateFormFile("file", objectCreateRequest.File)
	file, _ := os.Open(objectCreateRequest.File)
	defer file.Close()

	if _, err := io.Copy(fw, file); err != nil {
		return nil, nil, err
	}

	mp.Close()

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", mp.FormDataContentType())
	req.Header.Set("User-Agent", c.client.UserAgent)

	data := new(objectRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Object, resp, err
}

func (c ObjectsCli) Update(ctx context.Context, id string, objectUpdateRequest *ObjectUpdateRequest) (*http.Response, error) {
	endpointURL := objectsBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodPut, endpointURL, objectUpdateRequest)
	if err != nil {
		return nil, err
	}

	data := &struct{}{}
	resp, err := c.client.Do(ctx, req, data)

	return resp, err
}

func (c ObjectsCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := objectsBasePath + "/" + id

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

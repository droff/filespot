package filespot

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	packageVersion = "1.0.0"
	UserAgent      = "platformcraft-filesport/" + packageVersion
	mediaType      = "application/json"
	APIBaseURL     = "https://api.platformcraft.ru/1/"
)

// Client manages communication with platformcraft API
type Client struct {
	client    *http.Client
	UserAgent string
	BaseURL   *url.URL

	// Authentication
	APIUserID  string
	APIUserKey string

	// Services provides communication with API endpoints
	Objects         ObjectsService
	Temp            TempService
	Streams         StreamsService
	Players         PlayersService
	Download        DownloadService
	DownloadTasks   DownloadTasksService
	Transcoder      TranscoderService
	TranscoderTasks TranscoderTasksService
}

// ErrorResponse handles API errors
type ErrorResponse struct {
	Response *http.Response
	Code     uint32 `json:"code"`
	Status   string `json:"status"`
	MsgUser  string `json:"msg_user"`
	MsgDev   string `json:"msg_dev"`
	Doc      string `json:"doc"`
	Advanced string `json:"advanced"`
}

// NewClient returns client API
func NewClient(apiUserId, apiUserKey string) *Client {
	baseURL, _ := url.Parse(APIBaseURL)

	c := &Client{
		client:     http.DefaultClient,
		UserAgent:  UserAgent,
		APIUserID:  apiUserId,
		APIUserKey: apiUserKey,
		BaseURL:    baseURL,
	}

	// Endpoint communications
	c.Objects = &ObjectsCli{c}
	c.Temp = &TempCli{c}
	c.Streams = &StreamsCli{c}
	c.Players = &PlayersCli{c}
	c.Download = &DownloadCli{c}
	c.DownloadTasks = &DownloadTasksCli{c}
	c.Transcoder = &TranscoderCli{c}
	c.TranscoderTasks = &TranscoderTasksCli{c}

	return c
}

// generateHash returns HMAC hash-sum for authentication
func (c *Client) generateHash(method, path, timestamp string) string {
	data := fmt.Sprintf("%v+%v%v?apiuserid=%v&timestamp=%v", method, c.BaseURL.Host, path, c.APIUserID, timestamp)
	mac := hmac.New(sha256.New, []byte(c.APIUserKey))
	mac.Write([]byte(data))

	return hex.EncodeToString(mac.Sum(nil))
}

// requestURL returns URL with formated request
func (c *Client) requestURL(method, endpointURL string) *url.URL {
	endpoint, _ := url.Parse(endpointURL)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	hash := c.generateHash(method, endpoint.Path, timestamp)

	q := endpoint.Query()
	q.Set("apiuserid", c.APIUserID)
	q.Set("timestamp", timestamp)
	q.Set("hash", hash)
	endpoint.RawQuery = q.Encode()

	return c.BaseURL.ResolveReference(endpoint)
}

// NewRequest creates a API request with HTTP method, endpoint path and payload
func (c *Client) NewRequest(ctx context.Context, method, endpointURL string, body interface{}) (*http.Request, error) {
	u := c.requestURL(method, endpointURL)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// DoClientRequest submits request
func DoClientRequest(ctx context.Context, c *Client, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return c.client.Do(req)
}

// Do sends request and returns API response
func (c *Client) Do(ctx context.Context, req *http.Request, data interface{}) (*http.Response, error) {
	resp, err := DoClientRequest(ctx, c, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		resp.Body.Close()
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CheckResponse checks response for errors
func CheckResponse(resp *http.Response) error {
	code := resp.StatusCode
	if code >= 200 && code <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: resp}
	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

// Error returns formated error
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %v - %v\n\t%v\n\t%v", e.Code, e.Status, e.MsgUser,
		e.MsgDev, e.Doc)
}

// addParams returns path with query params
func addParams(path string, params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return path, nil
	}

	pathURL, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	newPath := pathURL.Query()
	newParams, err := query.Values(params)
	if err != nil {
		return path, err
	}

	for k, v := range newParams {
		newPath[k] = v
	}

	pathURL.RawQuery = newParams.Encode()
	return pathURL.String(), nil
}

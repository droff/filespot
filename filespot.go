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
	"strconv"
	"time"
)

const (
	packageVersion = "1.0.0"
	UserAgent      = "platformcraft-filesport/" + packageVersion
	mediaType      = "application/json"
	APIBaseURL     = "https://api.platformcraft.ru/1/"
)

type Client struct {
	client    *http.Client
	UserAgent string
	BaseURL   *url.URL

	APIUserID  string
	APIUserKey string

	Objects ObjectsService
}

type ErrorResponse struct {
	Response *http.Response
	Code     uint32 `json:"code"`
	Status   string `json:"status"`
	MsgUser  string `json:"msg_user"`
	MsgDev   string `json:"msg_dev"`
	Doc      string `json:"doc"`
	Advanced string `json:"advanced"`
}

func NewClient(apiUserId, apiUserKey string) *Client {
	baseURL, _ := url.Parse(APIBaseURL)

	c := &Client{
		client:     http.DefaultClient,
		UserAgent:  UserAgent,
		APIUserID:  apiUserId,
		APIUserKey: apiUserKey,
		BaseURL:    baseURL,
	}

	c.Objects = &ObjectsCli{c}

	return c
}

func (c *Client) requestURL(method, endpointURL string) *url.URL {
	endpoint, _ := url.Parse(endpointURL)

	q := endpoint.Query()
	q.Set("apiuserid", c.APIUserID)
	q.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	endpoint.RawQuery = q.Encode()

	data := method + "+" + c.BaseURL.Host + endpoint.String()
	mac := hmac.New(sha256.New, []byte(c.APIUserKey))
	mac.Write([]byte(data))
	sha := hex.EncodeToString(mac.Sum(nil))

	q.Set("hash", sha)
	endpoint.RawQuery = q.Encode()
	baseURL := c.BaseURL.ResolveReference(endpoint)

	return baseURL
}

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

func DoClientRequest(ctx context.Context, c *Client, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return c.client.Do(req)
}

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

func CheckResponse(resp *http.Response) error {
	code := resp.StatusCode
	if (code >= 200 && code <= 208) || code == 226 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: resp}
	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %v - %v\n\t%v\n\t%v", e.Code, e.Status, e.MsgUser,
		e.MsgDev, e.Doc)
}

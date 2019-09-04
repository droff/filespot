package filespot

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return resp, err
}

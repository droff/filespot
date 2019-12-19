package filespot

import (
	"context"
	"net/http"
)

const tempBasePath = "/1/temp"

// TempService implements interface with API /temp endpoint.
// See https://doc.platformcraft.ru/filespot/api/en/#temp
type TempService interface {
	List(context.Context, interface{}) (*linksRoot, *http.Response, error)
	Get(context.Context, string) (*Link, *http.Response, error)
	Create(context.Context, *LinkCreateRequest) (*Link, *http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
	Secure(context.Context, string, *SecureLinkRequest) (*SecureLink, *http.Response, error)
}

// TempCli handles communication with API
type TempCli struct {
	client *Client
}

// Link represents a platformcraft temporary Link
type Link struct {
	ID       string `json:"id"`
	ObjectID string `json:"object_id"`
	Href     string `json:"href"`
	Secure   bool   `json:"secure"`
	Exp      int    `json:"exp"`
	ForSale  bool   `json:"for_sale"`
	Geo      Geo    `json:"geo"`
}

// LinkCreateRequest identifies Link for the Create request
type LinkCreateRequest struct {
	ObjectID string `json:"object_id"`
	Endless  bool   `json:"endless"`
	Exp      int    `json:"exp"`
	Secure   bool   `json:"secure"`
	Geo      Geo    `json:"geo"`
}

// SecureLink represents a platformcraft Secure Link
type SecureLink struct {
	Hash string `json:"hash"`
	URL  string `json:"url"`
}

// SecureLinkRequest identifies data for Secure request
type SecureLinkRequest struct {
	IP string `json:"ip"`
	TS int    `json:"ts"`
}

// linksRoot represents a List root
type linksRoot struct {
	Links []Link `json:"links"`
	Count int    `json:"count"`
}

type linkRoot struct {
	Link *Link `json:"link"`
}

// TempListParams identifies as query params of List request
type TempListParams struct {
	ObjectID string `url:"object_id,omitempty"`
	ForSale  bool   `json:"for_sale,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
}

// List of Links
func (c TempCli) List(ctx context.Context, params interface{}) (*linksRoot, *http.Response, error) {
	path, err := addParams(tempBasePath, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(linksRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// Get Link
func (c TempCli) Get(ctx context.Context, id string) (*Link, *http.Response, error) {
	endpointURL := tempBasePath + "/" + id

	req, err := c.client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(linkRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Link, resp, err
}

// Create Link
func (c TempCli) Create(ctx context.Context, linkCreateRequest *LinkCreateRequest) (*Link, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, tempBasePath, linkCreateRequest)
	if err != nil {
		return nil, nil, err
	}

	data := new(linkRoot)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data.Link, resp, err
}

// Delete Link
func (c TempCli) Delete(ctx context.Context, id string) (*http.Response, error) {
	endpointURL := tempBasePath + "/" + id

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

// Secure Link
func (c TempCli) Secure(ctx context.Context, id string, secureLinkRequest *SecureLinkRequest) (*SecureLink, *http.Response, error) {
	endpointURL := tempBasePath + "/" + id + "/secure"

	req, err := c.client.NewRequest(ctx, http.MethodPost, endpointURL, secureLinkRequest)
	if err != nil {
		return nil, nil, err
	}

	data := new(SecureLink)
	resp, err := c.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

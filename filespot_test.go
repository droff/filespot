package filespot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	client *Client
	server *httptest.Server
	mux    *http.ServeMux
	ctx    = context.TODO()

	apiuserid  = "test"
	apiuserkey = "APIUserKey"
)

type RequestBody struct {
	Name string `json:"name"`
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	url, _ := url.Parse(server.URL)

	client = NewClient(apiuserid, apiuserkey)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.BaseURL == nil || c.BaseURL.String() != defaultAPIURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL, defaultAPIURL)
	}
}

func testClientDefaultUserAgent(t *testing.T, c *Client) {
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, expected %v", c.UserAgent, userAgent)
	}
}

func testClient(t *testing.T, c *Client) {
	testClientDefaultBaseURL(t, c)
	testClientDefaultUserAgent(t, c)
}

func TestNewClient(t *testing.T) {
	c := NewClient(apiuserid, apiuserkey)
	testClient(t, c)
}

func TestNewRequest(t *testing.T) {
	c := NewClient(apiuserid, apiuserkey)

	endpointURL := "/1/objects"
	requestBody := &RequestBody{Name: "filespot"}

	req, _ := c.NewRequest(ctx, http.MethodGet, endpointURL, requestBody)

	if req.URL.Path != endpointURL {
		t.Errorf("NewRequest URL Path = %v, expected %v", req.URL.Path, endpointURL)
	}

	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest UserAgent = %v, expected %v", userAgent, c.UserAgent)
	}

	q := req.URL.Query()

	if q.Get("apiuserid") != apiuserid {
		t.Errorf("NewRequest query apiuserid = %v, expected %v", q.Get("apiuserid"), apiuserid)
	}

	if q.Get("timestamp") == "" {
		t.Errorf("NewRequest query timestamp is missing")
	}

	if q.Get("hash") == "" {
		t.Errorf("NewRequest query hash is missin")
	}

	body, _ := ioutil.ReadAll(req.Body)
	expectedBody := `{"name":"filespot"}` + "\n"

	if string(body) != expectedBody {
		t.Errorf("NewRequest Body = %v, expected %v", string(body), expectedBody)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Request Method = %v, expected %v", r.Method, http.MethodGet)
		}

		fmt.Fprintf(w, `{"name":"filespot"}`)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)
	body := new(RequestBody)

	_, err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Errorf("Do error %v", err)
	}

	expectedBody := &RequestBody{Name: "filespot"}

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Do response = %v, expected %v", body, expectedBody)
	}
}

func TestDoWithErrorResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{
		    "code": 400,
		    "status": "fail",
		    "msg_user": "Api ID, timestamp and hash are required.",
		    "msg_dev": "Check Api id, timestamp and hash.",
		    "doc": "http://doc.platformcraft.ru/filespot/api/#access",
		    "advanced": null
		}`, 400)
	})

	req, _ := client.NewRequest(ctx, http.MethodGet, "/", nil)

	_, err := client.Do(context.Background(), req, nil)
	if err == nil {
		t.Error("Do returns with expected error")
	}
}

func TestAddParams(t *testing.T) {
	pathURL := "/1/objects"
	params := &struct {
		Name   string `url:"name,omitempty"`
		Folder string `url:"folder,omitempty"`
	}{
		Name:   "test.mp4",
		Folder: "tmp",
	}

	expected := "/1/objects?folder=tmp&name=test.mp4"
	path, err := addParams(pathURL, params)
	if err != nil {
		t.Errorf("addParams returned error = %v", err)
	}

	if path != expected {
		t.Errorf("addParams = %v, expected %v", path, expected)
	}
}

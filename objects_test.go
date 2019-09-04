package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestObjectsList(t *testing.T) {
	setup()
	defer shutdown()

	mux.HandleFunc("/1/objects", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Objects List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "objects": [
                {
                    "id": "56787f0c044dfe226b000001",
                    "name": "test.mp4",
                    "path": "/test.mp4",
                    "is_dir": false,
                    "size": 985781,
                    "content_type": "video/mp4",
                    "create_date": "22.12.2015T01:37:00",
                    "latest_update": "",
                    "resource_url": "api.platformcraft.ru/objects/56787f0c044dfe226b000001",
                    "cdn_url": "cdn.platformcraft.ru/billy/test.mp4",
                    "vod_hls": "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test.mp4/playlist.m3u8",
                    "video": "video.platformcraft.ru/56787f0c044dfe226b000001",
                    "private": true,
                    "status": "ok"
                },
                {
                    "id": "56787f0c044dfe226b000002",
                    "name": "test1.mp4",
                    "path": "/test1.mp4",
                    "is_dir": false,
                    "size": 985781,
                    "content_type": "video/mp4",
                    "create_date": "22.12.2015T01:37:00",
                    "latest_update": "",
                    "resource_url": "api.platformcraft.ru/objects/56787f0c044dfe226b000002",
                    "cdn_url": "cdn.platformcraft.ru/billy/test1.mp4",
                    "vod_hls": "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test1.mp4/playlist.m3u8",
                    "video": "video.platformcraft.ru/56787f0c044dfe226b000001",
                    "private": false,
                    "status": "ok"
                }
            ]
        }`)
	})

	objects, _, err := client.Objects.List(ctx)
	if err != nil {
		t.Errorf("Objects.List returned error: %v", err)
	}

	expected := []Object{
		{
			ID:           "56787f0c044dfe226b000001",
			Name:         "test.mp4",
			Path:         "/test.mp4",
			IsDir:        false,
			Size:         985781,
			ContentType:  "video/mp4",
			CreateDate:   "22.12.2015T01:37:00",
			LatestUpdate: "",
			ResourceURL:  "api.platformcraft.ru/objects/56787f0c044dfe226b000001",
			CDNURL:       "cdn.platformcraft.ru/billy/test.mp4",
			VODHLS:       "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test.mp4/playlist.m3u8",
			Video:        "video.platformcraft.ru/56787f0c044dfe226b000001",
			Private:      true,
			Status:       "ok",
		},
		{
			ID:           "56787f0c044dfe226b000002",
			Name:         "test1.mp4",
			Path:         "/test1.mp4",
			IsDir:        false,
			Size:         985781,
			ContentType:  "video/mp4",
			CreateDate:   "22.12.2015T01:37:00",
			LatestUpdate: "",
			ResourceURL:  "api.platformcraft.ru/objects/56787f0c044dfe226b000002",
			CDNURL:       "cdn.platformcraft.ru/billy/test1.mp4",
			VODHLS:       "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test1.mp4/playlist.m3u8",
			Video:        "video.platformcraft.ru/56787f0c044dfe226b000001",
			Private:      false,
			Status:       "ok",
		},
	}

	if !reflect.DeepEqual(objects, expected) {
		t.Errorf("Objects.List = %v, expected %v", objects, expected)
	}
}

func TestObjectsGet(t *testing.T) {
	setup()
	defer shutdown()

	mux.HandleFunc("/1/objects/56787f0c044dfe226b000001", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Objects List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "object": {
                "id": "56787f0c044dfe226b000001",
                "name": "test.mp4",
                "path": "/test.mp4",
                "is_dir": false,
                "size": 985781,
                "content_type": "video/mp4",
                "create_date": "22.12.2015T01:37:00",
                "latest_update": "",
                "resource_url": "api.platformcraft.ru/objects/56787f0c044dfe226b000001",
                "cdn_url": "cdn.platformcraft.ru/billy/test.mp4",
                "vod_hls": "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test.mp4/playlist.m3u8",
                "video": "video.platformcraft.ru/56787f0c044dfe226b000001",
                "private": true,
                "status": "ok"
            }
        }`)
	})

	object, _, err := client.Objects.Get(ctx, "56787f0c044dfe226b000001")
	if err != nil {
		t.Errorf("Objects.Get returned error: %v", err)
	}

	expected := &Object{
		ID:           "56787f0c044dfe226b000001",
		Name:         "test.mp4",
		Path:         "/test.mp4",
		IsDir:        false,
		Size:         985781,
		ContentType:  "video/mp4",
		CreateDate:   "22.12.2015T01:37:00",
		LatestUpdate: "",
		ResourceURL:  "api.platformcraft.ru/objects/56787f0c044dfe226b000001",
		CDNURL:       "cdn.platformcraft.ru/billy/test.mp4",
		VODHLS:       "customer.cdn.ru/customer-vod/_definst_/mp4:billy/test.mp4/playlist.m3u8",
		Video:        "video.platformcraft.ru/56787f0c044dfe226b000001",
		Private:      true,
		Status:       "ok",
	}

	if !reflect.DeepEqual(object, expected) {
		t.Errorf("Objects.Get = %v, expected %v", object, expected)
	}
}

func TestObjectsDelete(t *testing.T) {
	setup()
	defer shutdown()

	mux.HandleFunc("/1/objects/56787f0c044dfe226b000001", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("Objects List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	resp, err := client.Objects.Delete(ctx, "56787f0c044dfe226b000001")
	if err != nil {
		t.Errorf("Objects.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		fmt.Errorf("Objects.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

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
		if r.Method != http.MethodGet {
			t.Errorf("Objects List request method = %v, expected %v", r.Method, http.MethodGet)
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

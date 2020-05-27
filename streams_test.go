package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStreamsList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Streams.List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "streams": [
                {
                    "id": "56edc536534b4478d3a83b7f",
                    "user": "56bdf4a53f4f716301b09ba3",
                    "name": "example",
                    "url": "https://example.ru/example/example.m3u8",
                    "is_instant_recording": false
                }
            ]
        }`)
	})

	expected := []Stream{
		{
			ID:                 "56edc536534b4478d3a83b7f",
			User:               "56bdf4a53f4f716301b09ba3",
			Name:               "example",
			URL:                "https://example.ru/example/example.m3u8",
			IsInstantRecording: false,
		},
	}

	streams, _, err := client.Streams.List(ctx)
	if err != nil {
		t.Errorf("Streams.List returned error: %v", err)
	}

	if !reflect.DeepEqual(streams, expected) {
		t.Errorf("Streams.List = %v, expected %v", streams, expected)
	}
}

func TestStreamsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/56edc536534b4478d3a83b7f", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Streams.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "stream": {
                "id": "56edc536534b4478d3a83b7f",
                "user": "56bdf4a53f4f716301b09ba3",
                "name": "example",
                "url": "https://example.ru/example/example.m3u8",
                "is_instant_recording": false
            }
        }`)
	})

	expected := &Stream{
		ID:                 "56edc536534b4478d3a83b7f",
		User:               "56bdf4a53f4f716301b09ba3",
		Name:               "example",
		URL:                "https://example.ru/example/example.m3u8",
		IsInstantRecording: false,
	}

	stream, _, err := client.Streams.Get(ctx, "56edc536534b4478d3a83b7f")
	if err != nil {
		t.Errorf("Streams.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(stream, expected) {
		t.Errorf("Streams.Get = %v, expected %v", stream, expected)
	}
}

func TestStreamsCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Streams.Create request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "stream": {
                "id": "56edc536534b4478d3a83b7f",
                "user": "56bdf4a53f4f716301b09ba3",
                "name": "example",
                "url": "https://example.ru/example/example.m3u8",
                "is_instant_recording": false
            }
        }`)
	})

	expected := &Stream{
		ID:                 "56edc536534b4478d3a83b7f",
		User:               "56bdf4a53f4f716301b09ba3",
		Name:               "example",
		URL:                "https://example.ru/example/example.m3u8",
		IsInstantRecording: false,
	}

	streamCreateRequest := &StreamCreateRequest{
		Name: "example",
		URL:  "https://example.ru/example/example.m3u8",
	}

	stream, _, err := client.Streams.Create(ctx, streamCreateRequest)
	if err != nil {
		t.Errorf("Streams.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(stream, expected) {
		t.Errorf("Streams.Create = %v, expected %v", stream, expected)
	}
}

func TestStreamsDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/56cec7e2fa63afd0f843567d", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("Streams.Delete request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	resp, err := client.Streams.Delete(ctx, "56cec7e2fa63afd0f843567d")
	if err != nil {
		t.Errorf("Streams.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("Streams.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

func TestStreamsStart(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/rec/instant/start/56cec7e2fa63afd0f843567d", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Streams.Start request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	streamStartRequest := &StreamStartRequest{
		StopTimeout: 3600,
	}

	_, err := client.Streams.Start(ctx, "56cec7e2fa63afd0f843567d", streamStartRequest)
	if err != nil {
		t.Errorf("Streams.Start returned error: %v", err)
	}
}

func TestStreamsStop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/rec/instant/stop/56f19106534b44355afd96e1", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Streams.Stop request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "files":[
                {
                    "id":"56f19106534b44355afd96e1",
                    "name":"test 20160322.mp4",
                    "path":"/records/test 20160322.mp4",
                    "size":610768,
                    "content_type":"video/mp4",
                    "create_date":"22.03.2016T21:37:58",
                    "latest_update":"",
                    "resource_url":"api.platformcraft.ru/objects/56f19106534b44355afd96e1",
                    "video":"video.platformcraft.ru/56f19106534b44355afd96e1",
                    "cdn_url":"cdn.platformcraft.ru/test/records/test20160322.mp4",
                    "status":"ok"
                }
            ]
        }`)
	})

	expected := []File{
		{
			ID:           "56f19106534b44355afd96e1",
			Name:         "test 20160322.mp4",
			Path:         "/records/test 20160322.mp4",
			Size:         610768,
			ContentType:  "video/mp4",
			CreateDate:   "22.03.2016T21:37:58",
			LatestUpdate: "",
			ResourceURL:  "api.platformcraft.ru/objects/56f19106534b44355afd96e1",
			Video:        "video.platformcraft.ru/56f19106534b44355afd96e1",
			CDNURL:       "cdn.platformcraft.ru/test/records/test20160322.mp4",
			Status:       "ok",
		},
	}

	files, _, err := client.Streams.Stop(ctx, "56f19106534b44355afd96e1")
	if err != nil {
		t.Errorf("Streams.Stop returned error: %v", err)
	}

	if !reflect.DeepEqual(files, expected) {
		t.Errorf("Streams.Stop = %v, expected %v", files, expected)
	}
}

func TestStreamsCreateSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/rec/schedule/new/56cec7e2fa63afd0f843567d", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Streams.CreateSchedule request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "record_id": "5624cd5ac9a492f8b979b63f"
        }`)
	})

	recordID, _, err := client.Streams.CreateSchedule(ctx, "56cec7e2fa63afd0f843567d")
	if err != nil {
		t.Errorf("Streams.CreateSchedule returned error: %v", err)
	}

	expected := "5624cd5ac9a492f8b979b63f"

	if recordID != expected {
		t.Errorf("Streams.CreateSchedule = %v, expected %v", recordID, expected)
	}
}

func TestStreamsRec(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/rec/5624cd5ac9a492f8b979b63f", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Streams.Rec request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "record": {
                "status": "Finish",
                "files": ["5bd37808534b441c4acf7415"]
            }
        }`)
	})

	record, _, err := client.Streams.Rec(ctx, "5624cd5ac9a492f8b979b63f")
	if err != nil {
		t.Errorf("Streams.Rec returned error: %v", err)
	}

	expected := &Record{
		Status: "Finish",
		Files: []string{
			"5bd37808534b441c4acf7415",
		},
	}

	if !reflect.DeepEqual(record, expected) {
		t.Errorf("Streams.Rec = %v, expected %v", record, expected)
	}
}

func TestStreamsDeleteSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/streams/rec/schedule/del/56cec7e2fa63afd0f843567d/5624cd5ac9a492f8b979b63f", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("Streams.DeleteSchedule request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	resp, err := client.Streams.DeleteSchedule(ctx, "56cec7e2fa63afd0f843567d", "5624cd5ac9a492f8b979b63f")
	if err != nil {
		t.Errorf("Streams.DeleteSchedule returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("Streams.DeleteSchedule request code = %v, expected %v", resp.StatusCode, expected)
	}
}

package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTranscoderPresets(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder/presets", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Transcoder.Presets request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprint(w, `{
            "code": 200,
            "count": 3,
            "presets": [
                {
                    "id": "566b0fbf044dfe64f2000002",
                    "name": "System preset: Generic 1080p",
                    "container": "mp4",
                    "video": {
                        "codec": "H.264",
                        "bit_rate": "5400",
                        "fps": "29.97",
                        "max_width": "1920",
                        "max_height": "1080"
                    },
                    "audio": {
                        "codec": "AAC",
                        "bit_rate": "160",
                        "sample_rate": "44100",
                        "channels": "2"
                    },
                    "watermarks": {
                        "BottomLeft": {
                            "horizontal_align": "Left",
                            "horizontal_offset": "10%",
                            "vertical_align": "Bottom",
                            "vertical_offset": "10%",
                            "max_height": "10%",
                            "max_width": "10%",
                            "opacity": "100",
                            "sizing_policy": "ShrinkToFit"
                        },
                        "BottomRight": {
                            "horizontal_align": "Right",
                            "horizontal_offset": "10%",
                            "vertical_align": "Bottom",
                            "vertical_offset": "10%",
                            "max_height": "10%",
                            "max_width": "10%",
                            "opacity": "100",
                            "sizing_policy": "ShrinkToFit"
                        },
                        "Full": {
                            "horizontal_align": "Left",
                            "horizontal_offset": "0%",
                            "vertical_align": "Top",
                            "vertical_offset": "0%",
                            "max_height": "100%",
                            "max_width": "100%",
                            "opacity": "100",
                            "sizing_policy": "Fit"
                        },
                        "TopRight": {
                            "horizontal_align": "Right",
                            "horizontal_offset": "10%",
                            "vertical_align": "Top",
                            "vertical_offset": "10%",
                            "max_height": "10%",
                            "max_width": "10%",
                            "opacity": "100",
                            "sizing_policy": "ShrinkToFit"
                        }
                    }
                },
                {
                    "id": "566b0fbf044dfe64f2000003",
                    "name": "System preset: Generic 720p",
                    "container": "mp4",
                    "video": {
                        "codec": "H.264",
                        "bit_rate": "2400",
                        "fps": "29.97",
                        "max_width": "1280",
                        "max_height": "720"
                    },
                    "audio": {
                        "codec": "AAC",
                        "bit_rate": "160",
                        "sample_rate": "44100",
                        "channels": "2"
                    },
                    "watermarks": null
                },
                {
                    "id": "566b0fbf044dfe64f2000004",
                    "name": "System preset: Generic 480p 16:9",
                    "container": "mp4",
                    "video": {
                        "codec": "H.264",
                        "bit_rate": "1200",
                        "fps": "29.97",
                        "max_width": "854",
                        "max_height": "480"
                    },
                    "audio": {
                        "codec": "AAC",
                        "bit_rate": "128",
                        "sample_rate": "44100",
                        "channels": "2"
                    },
                    "watermarks": null
                }
            ]
        }`)
	})

	presets, _, err := client.Transcoder.Presets(ctx)
	if err != nil {
		t.Errorf("Transcoder.Presets returned error: %v", err)
	}

	expected := []Preset{
		{
			ID:        "566b0fbf044dfe64f2000002",
			Name:      "System preset: Generic 1080p",
			Container: "mp4",
			Video: map[string]string{
				"codec":      "H.264",
				"bit_rate":   "5400",
				"fps":        "29.97",
				"max_width":  "1920",
				"max_height": "1080",
			},
			Audio: map[string]string{
				"codec":       "AAC",
				"bit_rate":    "160",
				"sample_rate": "44100",
				"channels":    "2",
			},
			Watermarks: map[string]watermarkParams{
				"BottomLeft": {
					"horizontal_align":  "Left",
					"horizontal_offset": "10%",
					"vertical_align":    "Bottom",
					"vertical_offset":   "10%",
					"max_height":        "10%",
					"max_width":         "10%",
					"opacity":           "100",
					"sizing_policy":     "ShrinkToFit",
				},
				"BottomRight": {
					"horizontal_align":  "Right",
					"horizontal_offset": "10%",
					"vertical_align":    "Bottom",
					"vertical_offset":   "10%",
					"max_height":        "10%",
					"max_width":         "10%",
					"opacity":           "100",
					"sizing_policy":     "ShrinkToFit",
				},
				"Full": {
					"horizontal_align":  "Left",
					"horizontal_offset": "0%",
					"vertical_align":    "Top",
					"vertical_offset":   "0%",
					"max_height":        "100%",
					"max_width":         "100%",
					"opacity":           "100",
					"sizing_policy":     "Fit",
				},
				"TopRight": {
					"horizontal_align":  "Right",
					"horizontal_offset": "10%",
					"vertical_align":    "Top",
					"vertical_offset":   "10%",
					"max_height":        "10%",
					"max_width":         "10%",
					"opacity":           "100",
					"sizing_policy":     "ShrinkToFit",
				},
			},
		},
		{
			ID:        "566b0fbf044dfe64f2000003",
			Name:      "System preset: Generic 720p",
			Container: "mp4",
			Video: map[string]string{
				"codec":      "H.264",
				"bit_rate":   "2400",
				"fps":        "29.97",
				"max_width":  "1280",
				"max_height": "720",
			},
			Audio: map[string]string{
				"codec":       "AAC",
				"bit_rate":    "160",
				"sample_rate": "44100",
				"channels":    "2",
			},
			Watermarks: nil,
		},
		{
			ID:        "566b0fbf044dfe64f2000004",
			Name:      "System preset: Generic 480p 16:9",
			Container: "mp4",
			Video: map[string]string{
				"codec":      "H.264",
				"bit_rate":   "1200",
				"fps":        "29.97",
				"max_width":  "854",
				"max_height": "480",
			},
			Audio: map[string]string{
				"codec":       "AAC",
				"bit_rate":    "128",
				"sample_rate": "44100",
				"channels":    "2",
			},
			Watermarks: nil,
		},
	}

	if !reflect.DeepEqual(presets, expected) {
		t.Errorf("Transcoder.Presets = %v, expected %v", presets, expected)
	}
}

func TestTranscoderCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder/56787f0c044dfe226b000001", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Transcoder.Create request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "message": "Task create success.",
            "active_tasks": 2,
            "task_id": "56ea79e2534b4423135ebfd3"
        }`)
	})

	transcoderCreateRequest := &TranscoderCreateRequest{
		Presets:     []string{"5676a27cf9cb101634000002", "5676a27cf9cb101634000003"},
		Path:        "/test",
		Watermarks:  watermark{"Full": "5adfa939534b446a607d9937"},
		DelOriginal: false,
	}

	transcoder, resp, err := client.Transcoder.Create(ctx, "56787f0c044dfe226b000001", transcoderCreateRequest)
	if err != nil {
		t.Errorf("Transcoder.Create returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Transcoder.Create request code = %v, expected %v", resp.StatusCode, http.StatusOK)
	}

	expected := &Transcoder{
		ActiveTasks: 2,
		TaskID:      "56ea79e2534b4423135ebfd3",
	}
	if !reflect.DeepEqual(transcoder, expected) {
		t.Errorf("Transcoder.Create = %v, expected %v", transcoder, expected)
	}
}

func TestTranscoderConcat(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Transcoder.Concat request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "message": "Task create success.",
            "active_tasks": 2,
            "task_id": "56ea79e2534b4423135ebfd3"
        }`)
	})

	transcoderConcatRequest := &TranscoderConcatRequest{
		Files: []string{"5adfa9bc534b446a7f5d3c0c", "5ade3ad1534b445140d49539"},
		Path:  "/dir",
		Name:  "concat.mp4",
	}

	transcoder, resp, err := client.Transcoder.Concat(ctx, transcoderConcatRequest)
	if err != nil {
		t.Errorf("Transcoder.Concat returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Transcoder.Concat request code = %v, expected %v", resp.StatusCode, http.StatusOK)
	}

	expected := &Transcoder{
		ActiveTasks: 2,
		TaskID:      "56ea79e2534b4423135ebfd3",
	}
	if !reflect.DeepEqual(transcoder, expected) {
		t.Errorf("Transcoder.Concat = %v, expected %v", transcoder, expected)
	}
}

func TestTranscoderHLS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder/hls/56787f0c044dfe226b000001", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Transcoder.HLS request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "message": "Task create success.",
            "active_tasks": 2,
            "task_id": "56ea79e2534b4423135ebfd3"
        }`)
	})

	transcoderHLSRequest := &TranscoderHLSRequest{
		Presets:         []string{"5693b900534b440cd053f5a5", "5693b900534b440cd053f5a3"},
		SegmentDuration: 20,
	}

	transcoder, resp, err := client.Transcoder.HLS(ctx, "56787f0c044dfe226b000001", transcoderHLSRequest)
	if err != nil {
		t.Errorf("Transcoder.HLS returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Transcoder.HLS request code = %v, expected %v", resp.StatusCode, http.StatusOK)
	}

	expected := &Transcoder{
		ActiveTasks: 2,
		TaskID:      "56ea79e2534b4423135ebfd3",
	}
	if !reflect.DeepEqual(transcoder, expected) {
		t.Errorf("Transcoder.HLS = %v, expected %v", transcoder, expected)
	}
}

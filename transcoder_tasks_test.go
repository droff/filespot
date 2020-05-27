package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTranscoderTasksList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder_tasks", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("TranscoderTasks.List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "count": 3,
            "tasks": [
                {
                    "id": "56365b04044dfe6917000002",
                    "category": "encoding",
                    "title": "Encoding Копия test.360p.mp4",
                    "body": "Please wait.",
                    "status": "Progress",
                    "time_start": "01.11.2015T21:33:40",
                    "time_finish": "",
                    "lock": true
                },
                {
                    "id": "566b3a9d044dfe7381000002",
                    "category": "encoding",
                    "title": "Encoding test (480p).mp4",
                    "body": "Internal server error. Please try again later.",
                    "status": "Error",
                    "time_start": "12.12.2015T00:05:33",
                    "time_finish": "12.12.2015T00:05:47",
                    "lock": false
                },
                {
                    "id": "566b4173044dfe771c000001",
                    "category": "encoding",
                    "title": "Encoding test (480p).mp4",
                    "body": "Encoding video success.",
                    "status": "Completed",
                    "time_start": "12.12.2015T00:34:43",
                    "time_finish": "12.12.2015T00:35:56",
                    "lock": false
                }
            ]
        }`)
	})

	tasks, _, err := client.TranscoderTasks.List(ctx)
	if err != nil {
		t.Errorf("TranscoderTasks.List returned error: %v", err)
	}

	expected := []Task{
		{
			ID:         "56365b04044dfe6917000002",
			Category:   "encoding",
			Title:      "Encoding Копия test.360p.mp4",
			Body:       "Please wait.",
			Status:     "Progress",
			TimeStart:  "01.11.2015T21:33:40",
			TimeFinish: "",
			Lock:       true,
		},
		{
			ID:         "566b3a9d044dfe7381000002",
			Category:   "encoding",
			Title:      "Encoding test (480p).mp4",
			Body:       "Internal server error. Please try again later.",
			Status:     "Error",
			TimeStart:  "12.12.2015T00:05:33",
			TimeFinish: "12.12.2015T00:05:47",
			Lock:       false,
		},
		{
			ID:         "566b4173044dfe771c000001",
			Category:   "encoding",
			Title:      "Encoding test (480p).mp4",
			Body:       "Encoding video success.",
			Status:     "Completed",
			TimeStart:  "12.12.2015T00:34:43",
			TimeFinish: "12.12.2015T00:35:56",
			Lock:       false,
		},
	}

	if !reflect.DeepEqual(tasks, expected) {
		t.Errorf("TranscoderTasks.List = %v, expected %v", tasks, expected)
	}
}

func TestTranscoderTasksGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder_tasks/56787f0c044dfe226b000001", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("TranscoderTasks.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "files":["56ea7c4e534b44353536586c"],
            "status": "success",
            "task": {
                "id": "56365b04044dfe6917000002",
                "category": "encoding",
                "title": "Encoding Копия test.360p.mp4",
                "body": "Please wait.",
                "status": "Progress",
                "time_start": "01.11.2015T21:33:40",
                "time_finish": "",
                "lock": true
            }
        }`)
	})

	task, _, err := client.TranscoderTasks.Get(ctx, "56787f0c044dfe226b000001")
	if err != nil {
		t.Errorf("TranscoderTasks.Get returned error: %v", err)
	}

	expected := &Task{
		ID:         "56365b04044dfe6917000002",
		Category:   "encoding",
		Title:      "Encoding Копия test.360p.mp4",
		Body:       "Please wait.",
		Status:     "Progress",
		TimeStart:  "01.11.2015T21:33:40",
		TimeFinish: "",
		Lock:       true,
	}

	if !reflect.DeepEqual(task, expected) {
		t.Errorf("TranscoderTasks.Get = %v, expected %v", task, expected)
	}
}

func TestTranscoderTasksHLS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder_tasks/hls/56ea7c2f534b440c83fe85bc", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("TranscoderTasks.HLS request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "files": {
                "master_playlist": "56ea7c4e534b44353536586c",
                "media": [
                    {
                        "playlist": "56ea7c4e534b44353536586d",
                        "files": [
                            "56ea7c4f534b44353536586e",
                            "56ea7c50534b44353536586f",
                            "56ea7c50534b443535365870"
                        ]
                    },
                    {
                        "playlist": "56ea7c55534b443535365878",
                        "files": [
                            "56ea7c55534b443535365879",
                            "56ea7c55534b44353536587a",
                            "56ea7c56534b44353536587b"
                        ]
                    }
                ]
            },
            "status": "success",
            "task": {
                "id": "56ea7c2f534b440c83fe85bc",
                "category": "encoding",
                "title": "Encoding abc.mp4",
                "body": "Encoding video success.",
                "status": "Completed",
                "time_start": "17.03.2016T12:43:11",
                "time_finish": "17.03.2016T12:43:53",
                "lock": false
            }
        }`)
	})

	task, _, err := client.TranscoderTasks.HLS(ctx, "56ea7c2f534b440c83fe85bc")
	if err != nil {
		t.Errorf("TranscoderTasks.HLS returned error: %v", err)
	}

	expected := &Task{
		ID:         "56ea7c2f534b440c83fe85bc",
		Category:   "encoding",
		Title:      "Encoding abc.mp4",
		Body:       "Encoding video success.",
		Status:     "Completed",
		TimeStart:  "17.03.2016T12:43:11",
		TimeFinish: "17.03.2016T12:43:53",
		Lock:       false,
	}

	if !reflect.DeepEqual(task, expected) {
		t.Errorf("TranscoderTasks.HLS = %v, expected %v", task, expected)
	}
}

func TestTranscoderTasksDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/transcoder_tasks/56ea7c2f534b440c83fe85bc", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("TranscoderTasks.Delete request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "message": "Task delete success.",
            "status": "success"
        }`)
	})

	resp, err := client.TranscoderTasks.Delete(ctx, "56ea7c2f534b440c83fe85bc")
	if err != nil {
		t.Errorf("TranscoderTasks.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("TranscoderTasks.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

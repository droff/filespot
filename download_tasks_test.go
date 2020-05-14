package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDownloadTasksList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/download_tasks", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("DowloadTasks.List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "count": 2,
            "tasks": [
                {
                    "id": "56365b04044dfe6917000002",
                    "category": "download",
                    "title": "Download test1.mp4",
                    "body": "Please wait.",
                    "status": "Progress",
                    "time_start": "01.11.2015T21:33:40",
                    "time_finish": "",
                    "lock": true
                },
                {
                    "id": "57011e2a534b44741fc67880",
                    "category": "download",
                    "title": "Download abc.mp4",
                    "body": "Download success.",
                    "status": "Completed",
                    "time_start": "03.04.2016T16:44:10",
                    "time_finish": "03.04.2016T16:44:11",
                    "lock": false
                }
            ]
        }`)
	})

	expected := []Task{
		{
			ID:         "56365b04044dfe6917000002",
			Category:   "download",
			Title:      "Download test1.mp4",
			Body:       "Please wait.",
			Status:     "Progress",
			TimeStart:  "01.11.2015T21:33:40",
			TimeFinish: "",
			Lock:       true,
		},
		{
			ID:         "57011e2a534b44741fc67880",
			Category:   "download",
			Title:      "Download abc.mp4",
			Body:       "Download success.",
			Status:     "Completed",
			TimeStart:  "03.04.2016T16:44:10",
			TimeFinish: "03.04.2016T16:44:11",
			Lock:       false,
		},
	}

	root, _, err := client.DownloadTasks.List(ctx)
	if err != nil {
		t.Errorf("DownloadTasks.List returned error: %v", err)
	}

	if !reflect.DeepEqual(root.Tasks, expected) {
		t.Errorf("DownloadTasks.List = %v, expected %v", root.Tasks, expected)
	}
}

func TestDownloadTasksGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/download_tasks/57011e2a534b44741fc67880", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("DownloadTasks.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "task": {
                "id": "57011e2a534b44741fc67880",
                "category": "download",
                "title": "Download abc.mp4",
                "body": "Download success.",
                "status": "Completed",
                "time_start": "03.04.2016T16:44:10",
                "time_finish": "03.04.2016T16:44:11",
                "lock": false
            }
        }`)
	})

	expected := &Task{
		ID:         "57011e2a534b44741fc67880",
		Category:   "download",
		Title:      "Download abc.mp4",
		Body:       "Download success.",
		Status:     "Completed",
		TimeStart:  "03.04.2016T16:44:10",
		TimeFinish: "03.04.2016T16:44:11",
		Lock:       false,
	}

	task, _, err := client.DownloadTasks.Get(ctx, "57011e2a534b44741fc67880")
	if err != nil {
		t.Errorf("DownloadTasks.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(task, expected) {
		t.Errorf("DownloadTasks.Get = %v, expected %v", task, expected)
	}
}

func TestDownloadTasksDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/download_tasks/57011e2a534b44741fc67880", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("DownloadTasks.Delete request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "message": "Task delete success.",
            "status": "success"
        }`)
	})

	resp, err := client.DownloadTasks.Delete(ctx, "57011e2a534b44741fc67880")
	if err != nil {
		t.Errorf("DownloadTasks.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("DownloadTasks.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

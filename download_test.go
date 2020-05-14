package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDownloadCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/download", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Download.Create request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "message": "Task create success.",
            "active_tasks": 1,
            "task_id": "56ea79e2534b4423135ebfaa"
        }`)
	})

	params := &DownloadCreateParams{
		URL:  "http://test-domain.ru/test/test.mp4",
		Path: "/test",
	}
	download, resp, err := client.Download.Create(ctx, params)
	if err != nil {
		t.Errorf("Download.Create returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Download.Create request code = %v, expected %v", resp.StatusCode, http.StatusOK)
	}

	expected := &Download{
		Message:     "Task create success.",
		ActiveTasks: 1,
		TaskID:      "56ea79e2534b4423135ebfaa",
	}
	if !reflect.DeepEqual(download, expected) {
		t.Errorf("Download.Create = %v, expected %v", download, expected)
	}
}

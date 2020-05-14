package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStorageGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/storage", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Storage.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "storage": {
                "used": 97537237,
                "limit": 107374182
            }
        }`)
	})

	expected := &Storage{
		Used:  97537237,
		Limit: 107374182,
	}

	storage, _, err := client.Storage.Get(ctx)
	if err != nil {
		t.Errorf("Storage.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(storage, expected) {
		t.Errorf("Storage.Get = %v, expected %v", storage, expected)
	}
}

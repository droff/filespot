package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTempList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/temp", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Temp.List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "count": 1,
            "links": [
                {
                    "id": "58edef62534b4466a3ec333e",
                    "object_id": "58d51270534b440de6b0d075",
                    "href": "cdn.platformcraft.ru/temp/58edef62534b4466a3ec333e",
                    "secure": true,
                    "exp": 0,
                    "for_sale": false,
                    "geo": {"EU": {"RU": true}}
                }
            ]
        }`)
	})

	expected := []Link{
		{
			ID:       "58edef62534b4466a3ec333e",
			ObjectID: "58d51270534b440de6b0d075",
			Href:     "cdn.platformcraft.ru/temp/58edef62534b4466a3ec333e",
			Secure:   true,
			Exp:      0,
			ForSale:  false,
			Geo: Geo{
				"EU": {
					"RU": true,
				},
			},
		},
	}

	params := &TempListParams{
		Secure: true,
	}
	root, _, err := client.Temp.List(ctx, params)
	if err != nil {
		t.Errorf("Temp.List returned error: %v", err)
	}

	if !reflect.DeepEqual(root.Links, expected) {
		t.Errorf("Temp.List = %v, expected %v", root.Links, expected)
	}
}

func TestTempGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/temp/58edef62534b4466a3ec333e", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Temp.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "link": {
                "id":"58edef62534b4466a3ec333e",
                "object_id":"56787f0c044dfe226b000001",
                "href":"cdn.platformcraft.ru/temp/58ee48ca534b4409844c8f7a",
                "secure":true,
                "exp":1492008763,
                "for_sale":false,
                "geo": {"EU": {"RU": true}}
            }
        }`)
	})

	expected := &Link{
		ID:       "58edef62534b4466a3ec333e",
		ObjectID: "56787f0c044dfe226b000001",
		Href:     "cdn.platformcraft.ru/temp/58ee48ca534b4409844c8f7a",
		Secure:   true,
		Exp:      1492008763,
		ForSale:  false,
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	link, _, err := client.Temp.Get(ctx, "58edef62534b4466a3ec333e")
	if err != nil {
		t.Errorf("Temp.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(link, expected) {
		t.Errorf("Temp.Get = %v, expected %v", link, expected)
	}
}

func TestTempCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/temp", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Temp.Create request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "link": {
                "id":"58ee48ca534b4409844c8f7a",
                "object_id":"56787f0c044dfe226b000001",
                "href":"cdn.platformcraft.ru/temp/58ee48ca534b4409844c8f7a",
                "secure":true,
                "exp":1492008763,
                "for_sale":false,
                "geo": {"EU": {"RU": true}}
            }
        }`)
	})

	expected := &Link{
		ID:       "58ee48ca534b4409844c8f7a",
		ObjectID: "56787f0c044dfe226b000001",
		Href:     "cdn.platformcraft.ru/temp/58ee48ca534b4409844c8f7a",
		Secure:   true,
		Exp:      1492008763,
		ForSale:  false,
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	linkCreateRequest := &LinkCreateRequest{
		ObjectID: "56787f0c044dfe226b000001",
		Endless:  false,
		Exp:      1492008763,
		Secure:   true,
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	link, _, err := client.Temp.Create(ctx, linkCreateRequest)
	if err != nil {
		t.Errorf("Temp.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(link, expected) {
		t.Errorf("Temp.Create = %v, expected %v", link, expected)
	}
}

func TestTempDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/temp/58edef62534b4466a3ec333e", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("Temp.Delete request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	resp, err := client.Temp.Delete(ctx, "58edef62534b4466a3ec333e")
	if err != nil {
		t.Errorf("Temp.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("Temp.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

func TestTempSecure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/temp/58ee48ca534b4409844c8f7a/secure", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("Temp.Secure request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "hash": "79cc441e5f602b37a1294a59ea8ae3deddeed63c1de8580f66c1323f91487aa9",
            "url": "https://example.com/temp/5cdd8082ef3db56742cd704a?hash=79cc441e5f602b37a1294a59ea8ae3deddeed63c1de8580f66c1323f91487aa9&timestamp=1558030392"
        }`)
	})

	secureLinkRequest := &SecureLinkRequest{
		IP: "188.111.110.11",
		TS: 1558030392,
	}

	expected := &SecureLink{
		Hash: "79cc441e5f602b37a1294a59ea8ae3deddeed63c1de8580f66c1323f91487aa9",
		URL:  "https://example.com/temp/5cdd8082ef3db56742cd704a?hash=79cc441e5f602b37a1294a59ea8ae3deddeed63c1de8580f66c1323f91487aa9&timestamp=1558030392",
	}

	secureLink, _, err := client.Temp.Secure(ctx, "58ee48ca534b4409844c8f7a", secureLinkRequest)
	if err != nil {
		t.Errorf("Temp.Secure returned error: %v", err)
	}

	if !reflect.DeepEqual(secureLink, expected) {
		t.Errorf("Temp.Secure = %v, expected %v", secureLink, expected)
	}
}

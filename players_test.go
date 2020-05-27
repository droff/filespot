package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPlayersList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/players", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("PLayers.List request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "count": 4,
            "count_on_page": 2,
            "players": [
                {
                    "id": "567d3643534b4474087c221e",
                    "name": "player_name_1",
                    "path": "/player_name_1",
                    "is_dir": false,
                    "videos": {
                        "360": "cdn.platformcraft.ru/alex/example_360.mp4",
                        "480": "cdn.platformcraft.ru/alex/example_480.mp4"
                    },
                    "screen_shot_url": "cdn.platformcraft.ru/alex/example.jpg",
                    "vast_ad_tag_url": "http://example.com/example-vast.xml",
                    "create_date": "25.12.2015T15:27:48",
                    "href": "video.platformcraft.ru/embed/567d3643534b4474087c221e",
                    "frame_tag": "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221e\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
                    "description": "test player",
                    "tags": [],
                    "geo": null
                }
            ],
            "paging": {
                "next": "api.platformcraft.ru/1/players?pagingts=1525858519&limit=2&start=2",
                "prev": null
            }
        }`)
	})

	expected := []Player{
		{
			ID:    "567d3643534b4474087c221e",
			Name:  "player_name_1",
			Path:  "/player_name_1",
			IsDir: false,
			Videos: videos{
				"360": "cdn.platformcraft.ru/alex/example_360.mp4",
				"480": "cdn.platformcraft.ru/alex/example_480.mp4",
			},
			ScreenShotURL: "cdn.platformcraft.ru/alex/example.jpg",
			VastAdTagURL:  "http://example.com/example-vast.xml",
			CreateDate:    "25.12.2015T15:27:48",
			Href:          "video.platformcraft.ru/embed/567d3643534b4474087c221e",
			FrameTag:      "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221e\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
			Description:   "test player",
			Tags:          tags{},
			Geo:           nil,
		},
	}

	players, _, err := client.Players.List(ctx, nil)
	if err != nil {
		t.Errorf("Players.List returned error: %v", err)
	}

	if !reflect.DeepEqual(players, expected) {
		t.Errorf("Players.List = %v, expected %v", players, expected)
	}
}

func TestPlayersGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/players/567d3643534b4474087c221e", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodGet
		if r.Method != m {
			t.Errorf("Players.Get request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "player": {
                "id": "567d3643534b4474087c221e",
                "name": "player_name_1",
                "path": "/player_name_1",
                "is_dir": false,
                "videos": {
                    "360": "cdn.platformcraft.ru/alex/example_360.mp4",
                    "480": "cdn.platformcraft.ru/alex/example_480.mp4"
                },
                "screen_shot_url": "cdn.platformcraft.ru/alex/example.jpg",
                "vast_ad_tag_url": "http://example.com/example-vast.xml",
                "create_date": "25.12.2015T15:27:48",
                "href": "video.platformcraft.ru/embed/567d3643534b4474087c221e",
                "frame_tag": "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221e\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
                "description": "test player",
                "tags": [],
                "geo": null
            }
        }`)
	})

	expected := &Player{
		ID:    "567d3643534b4474087c221e",
		Name:  "player_name_1",
		Path:  "/player_name_1",
		IsDir: false,
		Videos: videos{
			"360": "cdn.platformcraft.ru/alex/example_360.mp4",
			"480": "cdn.platformcraft.ru/alex/example_480.mp4",
		},
		ScreenShotURL: "cdn.platformcraft.ru/alex/example.jpg",
		VastAdTagURL:  "http://example.com/example-vast.xml",
		CreateDate:    "25.12.2015T15:27:48",
		Href:          "video.platformcraft.ru/embed/567d3643534b4474087c221e",
		FrameTag:      "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221e\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
		Description:   "test player",
		Tags:          tags{},
		Geo:           nil,
	}

	player, _, err := client.Players.Get(ctx, "567d3643534b4474087c221e")
	if err != nil {
		t.Errorf("Players.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Get = %v, expected %v", player, expected)
	}
}

func TestPlayersCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/players", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPost
		if r.Method != m {
			t.Errorf("PLayers.Create request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success",
            "player": {
                "id": "567d3643534b4474087c221d",
                "name": "player_name",
                "path": "/test/player_name",
                "is_dir": false,
                "videos": {
                    "360": "cdn.platformcraft.ru/alex/example_360.mp4",
                    "480": "cdn.platformcraft.ru/alex/example_480.mp4"
                },
                "screen_shot_url": "cdn.platformcraft.ru/alex/example.jpg",
                "vast_ad_tag_url": "http://example.com/example-vast.xml",
                "create_date": "25.12.2015T15:27:47",
                "href": "video.platformcraft.ru/embed/567d3643534b4474087c221d",
                "frame_tag": "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221d\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
                "description": "test description",
                "tags": ["tag1","tag2"],
                "geo": {"EU": {"RU": true}}
            }
        }`)
	})

	expected := &Player{
		ID:    "567d3643534b4474087c221d",
		Name:  "player_name",
		Path:  "/test/player_name",
		IsDir: false,
		Videos: videos{
			"360": "cdn.platformcraft.ru/alex/example_360.mp4",
			"480": "cdn.platformcraft.ru/alex/example_480.mp4",
		},
		ScreenShotURL: "cdn.platformcraft.ru/alex/example.jpg",
		VastAdTagURL:  "http://example.com/example-vast.xml",
		CreateDate:    "25.12.2015T15:27:47",
		Href:          "video.platformcraft.ru/embed/567d3643534b4474087c221d",
		FrameTag:      "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221d\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
		Description:   "test description",
		Tags:          tags{"tag1", "tag2"},
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	playerCreateRequest := &PlayerCreateRequest{
		Name:   "player_name",
		Folder: "/test",
		Videos: videos{
			"360": "566bf3e9534b447f07e2baef",
			"480": "5624cd5ac9a492f8b979b63f",
		},
		ScreenShotID: "56238961c9a492f8b979b633",
		VastAdTagURL: "http://example.com/example-vast.xml",
		Description:  "test description",
		Tags:         tags{"tag1", "tag2"},
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	player, _, err := client.Players.Create(ctx, playerCreateRequest)
	if err != nil {
		t.Errorf("Players.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Create = %v, expected %v", player, expected)
	}
}

func TestPlayerUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/players/567d3643534b4474087c221d", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodPut
		if r.Method != m {
			t.Errorf("Players.Update request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	playerUpdateRequest := &PlayerUpdateRequest{
		Name:        "new_player_name",
		Description: "this description was updated",
	}
	resp, err := client.Players.Update(ctx, "567d3643534b4474087c221d", playerUpdateRequest)
	if err != nil {
		t.Errorf("Players.Update returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("Players.Update request code = %v, expected %v", resp.StatusCode, expected)
	}
}

func TestPlayersDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/1/players/567d3643534b4474087c221d", func(w http.ResponseWriter, r *http.Request) {
		m := http.MethodDelete
		if r.Method != m {
			t.Errorf("Players.Delete request method = %v, expected %v", r.Method, m)
		}

		fmt.Fprintf(w, `{
            "code": 200,
            "status": "success"
        }`)
	})

	resp, err := client.Players.Delete(ctx, "567d3643534b4474087c221d")
	if err != nil {
		t.Errorf("Players.Delete returned error: %v", err)
	}

	expected := http.StatusOK
	if resp.StatusCode != expected {
		t.Errorf("Players.Delete request code = %v, expected %v", resp.StatusCode, expected)
	}
}

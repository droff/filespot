package filespot

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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
		Videos: Videos{
			"360": "cdn.platformcraft.ru/alex/example_360.mp4",
			"480": "cdn.platformcraft.ru/alex/example_480.mp4",
		},
		ScreenShotURL: "cdn.platformcraft.ru/alex/example.jpg",
		VastAdTagURL:  "http://example.com/example-vast.xml",
		CreateDate:    "25.12.2015T15:27:47",
		Href:          "video.platformcraft.ru/embed/567d3643534b4474087c221d",
		FrameTag:      "<iframe width=\"558\" height=\"264\" src=\"video.platformcraft.ru/embed/567d3643534b4474087c221d\" frameBorder=\"0\" scrolling=\"no\" allowFullScreen></iframe>",
		Description:   "test description",
		Tags:          Tags{"tag1", "tag2"},
		Geo: Geo{
			"EU": {
				"RU": true,
			},
		},
	}

	playerCreateRequest := &PlayerCreateRequest{
		Name:   "player_name",
		Folder: "/test",
		Videos: Videos{
			"360": "566bf3e9534b447f07e2baef",
			"480": "5624cd5ac9a492f8b979b63f",
		},
		ScreenShotID: "56238961c9a492f8b979b633",
		VastAdTagURL: "http://example.com/example-vast.xml",
		Description:  "test description",
		Tags:         Tags{"tag1", "tag2"},
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

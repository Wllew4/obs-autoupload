package ttv

import "encoding/json"

type (
	ttv_access_token struct {
		Access_token string
		Expires_in   json.Number
		Token_type   string
	}

	ttv_res[T any] struct {
		Data []T
	}

	ttv_user struct {
		Id                string
		Login             string
		Display_name      string
		Type              string
		Broadcaster_type  string
		Description       string
		Profile_image_url string
		Offline_image_url string
		View_count        int
		Email             string
		Created_at        string
	}

	ttv_video struct {
		Id             string
		Stream_id      string
		User_id        string
		User_login     string
		User_name      string
		Title          string
		Description    string
		Created_at     string
		Published_at   string
		Url            string
		Thumbnail_url  string
		Viewable       string
		View_count     int
		Language       string
		Type           string
		Duration       string
		Muted_segments []ttv_video_muted_segment
	}

	ttv_video_muted_segment struct {
		Duration int
		Offset   int
	}
)

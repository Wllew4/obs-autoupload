from .yt_oauth import get_authenticated_service
from googleapiclient.errors import HttpError

def add_to_playlist(playlist_id: str, video_id: str):
	youtube = get_authenticated_service()
	try:
		body={
			"snippet": {
				"playlistId": playlist_id,
				"resourceId": {
					"videoId": video_id,
					"kind": "youtube#video"
				},
				"position": 0
			}
		}
		return youtube.playlistItems().insert(
			part=",".join(body.keys()),
			body=body
		).execute()
	except HttpError as e:
		print("An HTTP error %d occurred:\n%s" % (e.resp.status, e.content))

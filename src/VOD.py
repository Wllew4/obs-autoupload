import os
import glob
import requests
import json

from .secrets import TTV_CREDS, CONFIG

class VOD:
	title		: str
	date		: str
	path		: str
	stream_id	: str
	ttv_url		: str
	yt_url		: str

	def __init__(self):
		self.__fetch_info()
		self.__fetch_latest_vid()
	
	def __fetch_info(self):
		APP_ACCESS_TOKEN = json.loads(requests.post('https://id.twitch.tv/oauth2/token',
			json={
				'client_id': TTV_CREDS['client_id'],
				'client_secret': TTV_CREDS['client_secret'],
				'grant_type': 'client_credentials'
			}).text)['access_token']

		USER_ID = json.loads(requests.get('https://api.twitch.tv/helix/users?login=%s' % CONFIG['ttv']['LOGIN'],
			headers={
				'Authorization': 'Bearer %s' % APP_ACCESS_TOKEN,
				'Client-Id': TTV_CREDS['client_id']
			}).text)['data'][0]['id']
		VIDEO = json.loads(requests.get('https://api.twitch.tv/helix/videos?user_id=%s&first=1&sort=time&type=archive' % USER_ID,
			headers={
				'Authorization': 'Bearer %s' % APP_ACCESS_TOKEN,
				'Client-Id': TTV_CREDS['client_id']
			}).text)['data'][0]
		
		self.stream_id	= VIDEO['stream_id']
		self.title		= VIDEO['title']
		self.ttv_url	= VIDEO['url']
		self.date		= VIDEO['created_at'][:10]

	def __fetch_latest_vid(self):
		vods: list = glob.glob(CONFIG['files']['VOD_DIR'])
		self.path = max(vods, key=os.path.getctime)


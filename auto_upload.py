import obspython as obs

def script_description(): return "Automatically upload VODs to YouTube"

def script_load(settings):
	print("Loading script...")
	obs.obs_frontend_add_event_callback(on_event)
	print("Successfully loaded")

def on_event(e):
	if(e == obs.OBS_FRONTEND_EVENT_STREAMING_STOPPED):
		print("Uploading!")
		main()

import time
import os
import shutil
from src.VOD import VOD
from src.yt_upload import upload
from src.yt_playlist import add_to_playlist

from src.secrets import CONFIG

def main():
	# Fetch VOD info
	vod: VOD = VOD()
	print(vod.title, vod.date, vod.ttv_url, vod.stream_id, vod.path, sep='\n')

	# Upload VOD
	vod.yt_url = upload(vod.path, vod.title,
		CONFIG['upload']['DESCRIPTION'].format(
			TITLE=vod.title,
			DATE=vod.date,
			URL=vod.ttv_url,
			ID=vod.stream_id,
			PATH=vod.path
		),
		CONFIG['upload']['VISIBILITY'],
		CONFIG['upload']['CATEGORY_ID'],
		CONFIG['upload']['TAGS'])
	
	# Cleanup
	if CONFIG['upload']['PLAYLIST_ID']:
		time.sleep(2)
		add_to_playlist(CONFIG['upload']['PLAYLIST_ID'], vod.yt_url)
	if CONFIG['files']['POST_UPLOAD_STEP'] == "delete":
		os.remove(vod.path)
	if CONFIG['files']['POST_UPLOAD_STEP'] == "move":
		shutil.move(
			vod.path,
			CONFIG['files']['VOD_ARCHIVE_DIR'] + '\\' + os.path.basename(vod.path))
		open(
			CONFIG['files']['VOD_ARCHIVE_DIR']
			+ '\\'
			+ os.path.basename(vod.path)
			+ '.meta', 'w', encoding='utf-8').write('%s\n%s\n%s\n%s\n' % (
				vod.title, vod.date, vod.yt_url, vod.ttv_url
			))

# main()

# Modified from sample here
# https://developers.google.com/youtube/v3/guides/uploading_a_video

import httplib2
import random
import time

from googleapiclient.errors import HttpError
from googleapiclient.http import MediaFileUpload

from .yt_oauth import get_authenticated_service

httplib2.RETRIES = 1
MAX_RETRIES = 10
RETRIABLE_STATUS_CODES = [500, 502, 503, 504]
VALID_PRIVACY_STATUSES = ("public", "private", "unlisted")

# This method implements an exponential backoff strategy to resume a
# failed upload.
def resumable_upload(insert_request):
	response = None
	error = None
	retry = 0
	while response is None:
		try:
			print("Uploading file...")
			status, response = insert_request.next_chunk()
			if response is not None:
				if 'id' in response:
					print("Video id '%s' was successfully uploaded." % response['id'])
					return response['id']
				else:
					exit("The upload failed with an unexpected response: %s" % response)
		except HttpError as e:
			if e.resp.status in RETRIABLE_STATUS_CODES:
				error = "A retriable HTTP error %d occurred:\n%s" % (e.resp.status, e.content)
			else:
				raise

	if error is not None:
		print(error)
		retry += 1
		if retry > MAX_RETRIES:
			exit("No longer attempting to retry.")

		max_sleep = 2 ** retry
		sleep_seconds = random.random() * max_sleep
		print("Sleeping %f seconds and then retrying..." % sleep_seconds)
		time.sleep(sleep_seconds)

def upload(
	filename		: str,
	title			: str,
	description		: str	= 'Livestream archive',
	privacyStatus	: str	= VALID_PRIVACY_STATUSES[1],
	category		: str	= '22',
	keywords		: list	= '',):

	youtube = get_authenticated_service()
	try:
		body = {
			'snippet': {
				'title': title,
				'description': description,
				'tags': keywords,
				'category': category,
			},
			'status': {
				'privacyStatus': privacyStatus
			}
		}

		# Call the API's videos.insert method to create and upload the video.
		insert_request = youtube.videos().insert(
			part=",".join(body.keys()),
			body=body,
			# The chunksize parameter specifies the size of each chunk of data, in
			# bytes, that will be uploaded at a time. Set a higher value for
			# reliable connections as fewer chunks lead to faster uploads. Set a lower
			# value for better recovery on less reliable connections.
			#
			# Setting "chunksize" equal to -1 in the code below means that the entire
			# file will be uploaded in a single HTTP request. (If the upload fails,
			# it will still be retried where it left off.) This is usually a best
			# practice, but if you're using Python older than 2.6 or if you're
			# running on App Engine, you should set the chunksize to something like
			# 1024 * 1024 (1 megabyte).
			media_body=MediaFileUpload(filename, chunksize=-1, resumable=True)
		)

		return resumable_upload(insert_request)
	except HttpError as e:
		print("An HTTP error %d occurred:\n%s" % (e.resp.status, e.content))

import httplib2
import os
import sys

from googleapiclient.discovery import build
from oauth2client.client import flow_from_clientsecrets
from oauth2client.file import Storage
from oauth2client.tools import run_flow

YOUTUBE_SCOPE = "https://www.googleapis.com/auth/youtube"
YOUTUBE_API_SERVICE_NAME = "youtube"
YOUTUBE_API_VERSION = "v3"
CLIENT_SECRETS_FILE = ".credentials/.managed/google_secrets.json"

# This variable defines a message to display if the CLIENT_SECRETS_FILE is
# missing.
MISSING_CLIENT_SECRETS_MESSAGE = """
WARNING: Please configure OAuth 2.0

To make this sample run you will need to populate the client_secrets.json file
found at:

   %s

with information from the API Console
https://console.developers.google.com/

For more information about the client_secrets.json file format, please visit:
https://developers.google.com/api-client-library/python/guide/aaa_client_secrets
""" % os.path.abspath(os.path.join(os.path.dirname(__file__), CLIENT_SECRETS_FILE))


def get_authenticated_service():
	flow = flow_from_clientsecrets(CLIENT_SECRETS_FILE,
		scope=YOUTUBE_SCOPE,
		message=MISSING_CLIENT_SECRETS_MESSAGE)

	storage = Storage(".credentials/.managed/%s-oauth2.json" % sys.argv[0])
	credentials = storage.get()

	if credentials is None or credentials.invalid:
		credentials = run_flow(flow, storage)

	return build(YOUTUBE_API_SERVICE_NAME, YOUTUBE_API_VERSION,
		http=credentials.authorize(httplib2.Http()))

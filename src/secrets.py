import os
import toml
import json

def read_toml(path: str) -> dict: return toml.loads(open(PROJ_DIR + path).read())

PROJ_DIR = os.path.dirname(os.path.realpath(__file__)) + '\\..\\'

CONFIG = read_toml('config.toml')
TTV_CREDS = read_toml('.credentials\\twitch_secrets.toml')
GOOGLE_CREDS = read_toml('.credentials\\google_secrets.toml')
open(PROJ_DIR + '.credentials\\.managed\\google_secrets.json', 'w').write(json.dumps({
	"web": {
		"client_id": GOOGLE_CREDS['client_id'],
		"client_secret": GOOGLE_CREDS['client_secret'],
		"redirect_uris": [],
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://accounts.google.com/o/oauth2/token"
	}
}))

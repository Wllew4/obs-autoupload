# obs-autoupload
Automatically upload your VODs to YouTube when your stream ends!

## Prerequisites:
* OBS
* Twitch API credentials
* Google OAuth2 credentials

## Features:
* Automatically starts an upload when you "Stop Streaming"
* Prompts for confirmation before uploading
* Uses your stream title and other metadata from Twitch
* Highly configurable
	* Open source (you're free to make modifications)
* Can automatically add VODs to a YouTube playlist
* Can delete or move your old VODs to an archive directory

# Setup:
1. Download an unzip the latest release for your system.
1. Create a file `config.toml` based on `config.template.toml`
and provide your configuration.
1. Twitch credentials.
	1. Create an application in the [Twitch Developer Console](https://dev.twitch.tv/console/apps).
	1. Create a file `.credentials/twitch_secrets.toml` based on `.credentials/twitch_secrets.template.toml`.
and provide your client id and secret.
1. Google credentials.
	1. Create a project in the [Google Cloud Console](https://console.cloud.google.com/welcome).
	1. Add an Oauth 2.0 Client ID.
	1. Create a file `.credentials/google_secrets.toml` based on `.credentials/google_secrets.template.toml`.
1. Add script to OBS.
	1. In OBS, navigate to Tools > Scripts.
	1. In the "Scripts" tab, add "auto_upload.lua".
	1. Copy and paste the location of auto_upload.exe into the "auto_upload.exe path" field.
		* Absolute path, including the filename and extension
1. Make sure OBS is set to record video to your disk while you stream!
	* If you use the same encoder for recording and streaming, this shouldn't cause a significant performance hit
1. The next time you end a stream, your VOD will be uploaded to YouTube!

# Having difficulties?
Open an issue and I will respond when I get the chance:
https://github.com/Wllew4/obs-autoupload/issues.

# Consider supporting free and open source software 💖
If this script becomes part of your workflow, consider supporting free and open source software by throwing a dollar or two my way on [ko-fi](https://ko-fi.com/soupsu).

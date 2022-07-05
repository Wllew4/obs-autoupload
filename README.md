to-do:
* test
* publish/release
* yt video

# obs-autoupload
Automatically upload your VODs to YouTube when your stream ends!

## Prerequisites:
* OBS
* Python 3.6.8
* Twitch API credentials
* Google OAuth2 credentials

## Features:
* Automatically starts an upload in the background when you "Stop Streaming"
* Uses your stream title and other metadata from Twitch
* Highly configurable
	* Open source (you're free to make modifications)
* Can automatically add VODs to a YouTube playlist
* Can delete or move your old VODs to an archive

## Setup:
1. Clone this repository.
1. Install [Python 3.6.8](https://www.python.org/downloads/release/python-368/).
	* The version matters, 3.7+ are not supported.
1. Install Python packages.
	1. `pip install -r requirements.txt`.
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
	1. Select the "Python Settings" tab, and add the path to your Python installation
		* On Windows, this is `C:/Users/yourusername/AppData/Local/Programs/Python/Python36` by default.
	1. In the "Scripts" tab, add a script. Select "auto_upload.py" from the root of the project.
1. Make sure OBS is set to record video to your disk while you stream!
	* If you use the same encoder for recording and streaming, this shouldn't cause a significant performance hit
1. The next time you end a stream, your VOD will be uploaded to YouTube!

## Having difficulties?
Open an issue and I will respond when I get the chance:
https://github.com/Wllew4/obs-autoupload/issues.

## Consider supporting free software 💖
If this script becomes part of your workflow, consider supporting free and open source software by throwing a dollar or two my way on [my ko-fi page](https://ko-fi.com/soupsu).

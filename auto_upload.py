DEBUG = True

from src.main import main

if not DEBUG:
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

else: main()

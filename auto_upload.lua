obs = obslua

function script_description()
	return "Automatically upload your VODs to YouTube"
end

function script_properties()
	props = obs.obs_properties_create()
	obs.obs_properties_add_text(props, "exe_location", "auto_upload.exe path:", obs.OBS_TEXT_DEFAULT)
	return props
end

function script_load(settings)
	obs.obs_frontend_add_event_callback(on_event)
end

function on_event(e)
	if e == obs.OBS_FRONTEND_EVENT_STREAMING_STOPPED then
		local command = obs.obs_data_get_string(settings, "exe_location")
	
		if package.config:sub(1,1) == '\\' then
			command = "start " .. command
		end
		os.execute(command)
	end
end

<script lang="ts">
	import { api_call } from "../lib/api_call"

	type Config = {
		Files: {
			VOD_DIR: string
			POST_UPLOAD_STEP: string
			VOD_ARCHIVE_DIR: string
		},
		TTV: {
			LOGIN: string
		},
		Upload: {
			VISIBILITY: string
			DESCRIPTION: string
			CATEGORY_ID: string
			TAGS: string[]
			PLAYLIST_ID: string
		}
	}

	async function getVOD(): Promise<Config> {
		return JSON.parse(await api_call('/api/config'))
	}

	let data = getVOD()

</script>

{#await data} ...
{:then data}
	<tr>Files</tr>
	<tr>
		<td><label for="VOD_DIR">VOD_DIR</label></td>
		<td><input name="VOD_DIR" type="text" value={data.Files.VOD_DIR}/></td>
	</tr>
	<tr>
		<td><label for="POST_UPLOAD_STEP">POST_UPLOAD_STEP</label></td>
		<td><input name="POST_UPLOAD_STEP" type="text" value={data.Files.POST_UPLOAD_STEP}/></td>
	</tr>
	<tr>
		<td><label for="VOD_ARCHIVE_DIR">VOD_ARCHIVE_DIR</label></td>
		<td><input name="VOD_ARCHIVE_DIR" type="text" value={data.Files.VOD_ARCHIVE_DIR}/></td>
	</tr>
{/await}

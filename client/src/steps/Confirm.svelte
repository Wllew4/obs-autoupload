<script lang="ts">
	import { api_call } from "../lib/api_call"

	type VOD = {
		Title: string
		Date: string
		Path: string
		Stream_id: string
		Ttv_url: string
	}

	async function getVOD(): Promise<VOD> {
		return JSON.parse(await api_call('/api/vod'))
	}

	let data = getVOD()

</script>

{#await data} ...
{:then data} 
<form method="POST" action="/api/upload">
	<table>
		<tr>
			<td><label for="title">Title</label></td>
			<td><input name="title" type="text" value={data.Title}/></td>
		</tr>
		<tr>
			<td><label for="date">Date</label></td>
			<td><input name="date" type="text" value={data.Date}/></td>
		</tr>
		<tr>
			<td><label for="path">Path</label></td>
			<td><input name="path" type="text" value={data.Path}/></td>
		</tr>
		<tr>
			<td><label for="url">Url</label></td>
			<td><input name="url" type="text" value={data.Ttv_url}/></td>
		</tr>
		<tr>
			<td><button>Confirm</button></td>
		</tr>
	</table>
</form>
{/await}


<style lang="scss">
	input {
		width: 20rem;
	}
</style>

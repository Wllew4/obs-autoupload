<script lang="ts">

	let data = {
		title: "💖 FINISHING STRAY | @Soupsu_",
		date: "2022-08-15",
		path: "D:\\Videos\\VODs\\2022-08-15 19-50-42.mkv",
		url: "https://www.twitch.tv/videos/1563109387",
	}

	function confirm() {
		data.title	= document.getElementsByTagName("input")[0].value
		data.date	= document.getElementsByTagName("input")[1].value
		data.path	= document.getElementsByTagName("input")[2].value
		data.url	= document.getElementsByTagName("input")[3].value
		console.log(data)
	}

</script>


<table>
	<tr>
		<td>Title</td>
		<td><input value={data.title}></td>
	</tr>
	<tr>
		<td>Date</td>
		<td><input value={data.date}></td>
	</tr>
	<tr>
		<td>Path</td>
		<td><input value={data.path}></td>
	</tr>
	<tr>
		<td>Url</td>
		<td><input value={data.url}></td>
	</tr>
</table>

<button on:click={confirm}>Confirm</button>

<style lang="scss">
	input {
		width: 20rem;
	}
</style>

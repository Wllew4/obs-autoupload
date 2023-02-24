<script lang="ts">
	import Config from "./Config.svelte"
import VOD from "./VOD.svelte"
</script>

<form method="POST" action="/api/upload">
	<table>
		<VOD/>
		<Config/>
		<tr>
			<td><button>Confirm</button></td>
		</tr>
	</table>
</form>

<style lang="scss">
	:global input {
		width: 20rem;
	}
</style>

<script lang="ts">
	import { api_call } from "src/lib/api_call";
	export let valid = false

	let config: Promise<string> = api_call('/api/verify')

	let validated_elements = 0;

	function set_valid(value: string) {
		validated_elements++;
		if (validated_elements == 1) valid = true;
		return value
	}
</script>

Configuration: 
{#await config} Verifying...
	{:then data} { data == "OK" ?  set_valid(data) : "Error: " + data }
	{:catch error} Internal Error: {error}
{/await}


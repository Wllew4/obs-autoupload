<script lang="ts">
  import { api_call } from "./lib/api_call";
  import Confirm from "./steps/Confirm.svelte";
  import Verify from "./steps/Verify.svelte";

	let version: Promise<string> = api_call('/api/version')

	let valid_config = false
</script>

<svelte:head>
	<title>OBS-AutoUpload</title>
</svelte:head>

<div id="header">
	<div id="title">obs-autoupload
		{#await version} ...
		{:then data} { data }
		{:catch error} uh oh! {error}
		{/await}
	</div>
	<div id="credit">by <a href="https://twitter.com/soupsu_" target="_blank">Wllew4</a></div>
	<div id="links">
		<a href="https://ko-fi.com/soupsu" target="blank">
			<img class="logo" src="/ko-fi.png" alt="ko-fi logo" />Support me on Ko-fi
		</a>
		|
		<a href="https://github.com/Wllew4/obs-autoupload" target="blank">
			<img class="logo" src="/github.png" alt="github logo" />Source code
		</a>
	</div>
	<hr>

	<Verify bind:valid={valid_config}/>

	{#if valid_config}
		<Confirm/>
	{/if}

</div>

<style lang="scss">
	@import 'https://fonts.googleapis.com/css?family=Noto+Sans';

	:global(body) {
		font-family: 'Noto Sans';
	}

	#title {
		font-size: 2rem;
	}

	#credit {
		margin-bottom: 0.5rem;
	}

	.logo {
		height: 2rem;
		vertical-align: middle;
		position: relative;
		top: -3px;
		margin-right: 0.2rem;
	}
</style>

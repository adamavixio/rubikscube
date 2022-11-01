<script lang="ts">
	import { onMount } from "svelte";
	import { labelRenderer, renderer } from "./three";

	onMount(() => {
		document
			.querySelector("#three-container")
			.appendChild(renderer.domElement);
		document
			.querySelector("#three-container")
			.appendChild(labelRenderer.domElement);
	});
</script>

<main>
	<div id="three-container" />
</main>

<style>
</style>

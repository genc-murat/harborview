<script>
	import { onMount } from "svelte";
	import { images, fetchImages } from "../../stores/images";
	onMount(fetchImages);
</script>

<main>
	<h1>Docker Images</h1>
	<ul>
		{#each $images as image}
			<li>{image}</li>
		{/each}
	</ul>
</main>

<style>
	/* Ana sayfa stili */
	main {
		max-width: 800px;
		margin: 2rem auto;
		font-family: Arial, sans-serif;
		padding: 1rem;
		background-color: #f8f9fa;
		border-radius: 8px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	}

	h1 {
		text-align: center;
		color: #343a40;
		font-size: 2rem;
		margin-bottom: 1.5rem;
	}

	/* Liste stili */
	ul {
		list-style-type: none;
		padding: 0;
	}

	li {
		padding: 0.75rem 1rem;
		margin-bottom: 0.5rem;
		background-color: #ffffff;
		color: #495057;
		border-radius: 6px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		transition: transform 0.2s, box-shadow 0.2s;
		cursor: pointer;
	}

	li:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
		background-color: #e9ecef;
	}

	li:nth-child(odd) {
		background-color: #f1f3f5;
	}
</style>

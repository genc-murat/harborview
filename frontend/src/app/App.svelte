<script>
	import ImageList from "../components/ImageList.svelte";
	import SearchBar from "../components/SearchBar.svelte";
	import { fetchImages, deleteUnusedImages } from "../stores/images";

	// Fetch images on mount
	import { onMount } from "svelte";
	onMount(fetchImages);

	// Delete unused images
	const handleDeleteUnused = async () => {
		if (confirm("Are you sure you want to delete unused images?")) {
			await deleteUnusedImages();
			alert("Unused images deleted successfully!");
		}
	};
</script>

<main>
	<h1>Harbor View</h1>
	<SearchBar />
	<ImageList />
	<button class="delete-unused" on:click={handleDeleteUnused}>
		Delete Unused Images
	</button>
</main>

<style>
	main {
		padding: 2rem;
		font-family: Arial, sans-serif;
		max-width: 800px;
		margin: auto;
		background-color: #f9f9f9;
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
		border-radius: 8px;
	}
	h1 {
		text-align: center;
		color: #333;
	}
	.delete-unused {
		background-color: #d9534f;
		color: white;
		padding: 0.5rem 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		margin-top: 1rem;
		display: block;
		margin-left: auto;
		margin-right: auto;
	}
	.delete-unused:hover {
		background-color: #c9302c;
	}
</style>

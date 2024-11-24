<script>
	import { searchImages, searchResults } from "../stores/images";

	let query = "";

	// Search for images
	const handleSearch = async () => {
		if (query.trim() !== "") {
			await searchImages(query);
		} else {
			alert("Please enter a search query!");
		}
	};
</script>

<div class="search-bar">
	<input
		type="text"
		placeholder="Search images..."
		bind:value={query}
	/>
	<button on:click={handleSearch}>Search</button>
</div>

<div class="search-results">
	<h2>Search Results</h2>
	{#if $searchResults.length > 0}
		<ul>
			{#each $searchResults as result}
				<li>{result}</li>
			{/each}
		</ul>
	{:else}
		<p>No results found. Try a different query.</p>
	{/if}
</div>

<style>
	.search-bar {
		display: flex;
		justify-content: center;
		margin-bottom: 1rem;
	}
	input {
		width: 70%;
		padding: 0.5rem;
		border: 1px solid #ccc;
		border-radius: 4px;
	}
	button {
		background-color: #28a745;
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		margin-left: 0.5rem;
		border-radius: 4px;
		cursor: pointer;
	}
	button:hover {
		background-color: #218838;
	}
	.search-results {
		margin-top: 2rem;
		text-align: center;
	}
	.search-results h2 {
		color: #333;
	}
	.search-results ul {
		list-style: none;
		padding: 0;
	}
	.search-results li {
		background-color: #f9f9f9;
		padding: 0.5rem;
		margin: 0.5rem auto;
		width: 60%;
		border: 1px solid #ddd;
		border-radius: 4px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}
</style>

<script>
	import { images, fetchImageHistory, imageHistory, deleteImage, fetchImages } from "../stores/images";

	let selectedImage = ""; // Seçilen imaj adı
	let isModalOpen = false; // Modal görünürlüğünü kontrol eder

	// Fetch history and open modal
	const handleShowHistory = async (imageName) => {
		selectedImage = imageName;
		try {
			await fetchImageHistory(imageName);
			isModalOpen = true;
		} catch (error) {
			alert(`Failed to fetch history for ${imageName}`);
		}
	};

	// Delete an image
	const handleDeleteImage = async (imageName) => {
		if (confirm(`Are you sure you want to delete the image "${imageName}"?`)) {
			try {
				await deleteImage(imageName);
				alert(`Image "${imageName}" deleted successfully.`);
				await fetchImages(); // Listeyi yenile
			} catch (error) {
				alert(`Failed to delete image "${imageName}".`);
			}
		}
	};

	// Close modal
	const closeModal = () => {
		isModalOpen = false;
		selectedImage = "";
	};
</script>

<h2>Images</h2>
<ul>
	{#each $images as image}
		<li>
			<div class="image-name">{image}</div>
			<div class="image-actions">
				<button on:click={() => handleShowHistory(image)}>View History</button>
				<button class="delete-button" on:click={() => handleDeleteImage(image)}>Delete</button>
			</div>
		</li>
	{/each}
</ul>

{#if isModalOpen}
	<div class="overlay" on:click={closeModal}>
		<div class="modal" on:click|stopPropagation>
			<h3>History for {selectedImage}</h3>
			{#if Object.keys($imageHistory).length > 0}
				<ul>
					{#each Object.entries($imageHistory) as [key, value]}
						<li>
							<strong>{key}:</strong> {JSON.stringify(value)}
						</li>
					{/each}
				</ul>
			{:else}
				<p>No history available.</p>
			{/if}
			<button on:click={closeModal}>Close</button>
		</div>
	</div>
{/if}

<style>
	h2 {
		color: #555;
		margin-bottom: 1rem;
	}
	ul {
		list-style: none;
		padding: 0;
	}
	li {
		padding: 0.5rem;
		border-bottom: 1px solid #ddd;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.image-name {
		flex: 1;
	}

	.image-actions {
		display: flex;
		gap: 0.5rem;
	}

	button {
		background-color: #007bff;
		color: white;
		border: none;
		padding: 0.3rem 0.6rem;
		border-radius: 4px;
		cursor: pointer;
	}
	button:hover {
		background-color: #0056b3;
	}

	.delete-button {
		background-color: #d9534f;
	}
	.delete-button:hover {
		background-color: #c9302c;
	}

	/* Modal Styles */
	.overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: rgba(0, 0, 0, 0.5);
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 1000;
	}

	.modal {
		background: white;
		padding: 1.5rem;
		border-radius: 8px;
		width: 400px;
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
		position: relative;
	}

	h3 {
		margin-bottom: 1rem;
		color: #333;
	}

	.modal ul {
		list-style-type: none;
		padding: 0;
		margin: 0;
	}

	.modal li {
		margin-bottom: 0.5rem;
		color: #555;
	}

	.modal button {
		background-color: #007bff;
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 4px;
		cursor: pointer;
		display: block;
		margin: 1rem auto 0;
	}

	.modal button:hover {
		background-color: #0056b3;
	}

	p {
		color: #666;
		text-align: center;
	}
</style>

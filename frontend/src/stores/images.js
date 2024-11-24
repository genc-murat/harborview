import { writable } from "svelte/store";

// Store for images
export const images = writable([]);
export const imageHistory = writable([]);
export const searchResults = writable([]);
export const apiBaseURL = "http://localhost:3000/images"; // Backend URL

// Fetch all images
export const fetchImages = async () => {
	try {
		const res = await fetch(`${apiBaseURL}/`);
		const data = await res.json();
		images.set(data);
	} catch (error) {
		console.error("Failed to fetch images:", error);
	}
};

// Fetch history of a specific image
export const fetchImageHistory = async (imageName) => {
	try {
		const res = await fetch(`${apiBaseURL}/${imageName}/history`);
		const data = await res.json();
		imageHistory.set(data);
	} catch (error) {
		console.error("Failed to fetch image history:", error);
	}
};

// Search images
export const searchImages = async (query) => {
	try {
		const res = await fetch(`${apiBaseURL}/search?q=${query}`);
		const data = await res.json();
		searchResults.set(data);
	} catch (error) {
		console.error("Failed to search images:", error);
	}
};

// Delete unused images
export const deleteUnusedImages = async () => {
	try {
		await fetch(`${apiBaseURL}/unused`, { method: "DELETE" });
		await fetchImages(); // Refresh the image list
	} catch (error) {
		console.error("Failed to delete unused images:", error);
	}
};

// Delete a specific image
export const deleteImage = async (imageName) => {
	try {
		await fetch(`${apiBaseURL}/${imageName}`, { method: "DELETE" });
		await fetchImages(); // Refresh the image list after deletion
		console.log(`Image "${imageName}" deleted successfully.`);
	} catch (error) {
		console.error(`Failed to delete image "${imageName}":`, error);
	}
};

import { API_BASE_URL } from './constants';

// Fetch için default options
const defaultOptions = {
  headers: {
    'Cache-Control': 'no-cache, no-store, must-revalidate',
    'Pragma': 'no-cache',
    'Expires': '0'
  }
};

export const api = {
  // Containers
  getContainers: async () => {
    const response = await fetch(`${API_BASE_URL}/containers`, defaultOptions);
    if (!response.ok) throw new Error('Failed to fetch containers');
    return response.json();
  },

  getContainerById: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}`, defaultOptions);
    if (!response.ok) throw new Error('Failed to fetch container');
    return response.json();
  },

  // Container actions
  startContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/start`, {
      method: 'POST',
      ...defaultOptions
    });
    if (!response.ok) throw new Error('Failed to start container');
    return response.json();
  },

  stopContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/stop`, {
      method: 'POST',
      ...defaultOptions
    });
    if (!response.ok) throw new Error('Failed to stop container');
    return response.json();
  },

  restartContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/restart`, {
      method: 'POST',
      ...defaultOptions
    });
    if (!response.ok) throw new Error('Failed to restart container');
    return response.json();
  },

  // Images
  getImages: async () => {
    const response = await fetch(`${API_BASE_URL}/images`, defaultOptions);
    if (!response.ok) throw new Error('Failed to fetch images');
    return response.json();
  },

  getImageById: async (id) => {
    const response = await fetch(`${API_BASE_URL}/images/${id}/inspect`, defaultOptions);
    if (!response.ok) throw new Error('Failed to fetch image');
    return response.json();
  },

  // Image actions
  deleteImage: async (id) => {
    const response = await fetch(`${API_BASE_URL}/images/${id}`, {
      method: 'DELETE',
      ...defaultOptions
    });
    if (!response.ok) throw new Error('Failed to delete image');
    return response.json();
  },
};
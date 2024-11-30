// src/lib/api.js
import { API_BASE_URL } from './constants';

export const api = {
  // Containers
  getContainers: async () => {
    const response = await fetch(`${API_BASE_URL}/containers`);
    if (!response.ok) throw new Error('Failed to fetch containers');
    return response.json();
  },

  getContainerById: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}`);
    if (!response.ok) throw new Error('Failed to fetch container');
    return response.json();
  },

  // Container actions
  startContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/start`, {
      method: 'POST',
    });
    if (!response.ok) throw new Error('Failed to start container');
    return response.json();
  },

  stopContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/stop`, {
      method: 'POST',
    });
    if (!response.ok) throw new Error('Failed to stop container');
    return response.json();
  },

  restartContainer: async (id) => {
    const response = await fetch(`${API_BASE_URL}/containers/${id}/restart`, {
      method: 'POST',
    });
    if (!response.ok) throw new Error('Failed to restart container');
    return response.json();
  },

  // Images
  getImages: async () => {
    const response = await fetch(`${API_BASE_URL}/images`);
    if (!response.ok) throw new Error('Failed to fetch images');
    return response.json();
  },

  getImageById: async (id) => {
    const response = await fetch(`${API_BASE_URL}/images/${id}/inspect`);
    if (!response.ok) throw new Error('Failed to fetch image');
    return response.json();
  },

  // Image actions
  deleteImage: async (id) => {
    const response = await fetch(`${API_BASE_URL}/images/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) throw new Error('Failed to delete image');
    return response.json();
  },
};
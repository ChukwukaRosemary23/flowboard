import axios from 'axios';

// Use environment variable or fallback to localhost
const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8082/api/v1';

console.log('🌐 API Base URL:', API_BASE_URL);

// Create axios instance
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth
export const register = (data) => api.post('/auth/register', data);
export const login = (data) => api.post('/auth/login', data);

// Fixed: getCurrentUser - decode JWT token client-side
export const getCurrentUser = () => {
  const token = localStorage.getItem('token');
  if (!token) {
    return Promise.reject(new Error('No token found'));
  }
  
  try {
    // Decode JWT payload (middle part of token)
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const payload = JSON.parse(window.atob(base64));
    
    return Promise.resolve({ 
      data: {
        user_id: payload.user_id,
        username: payload.username,
        email: payload.email
      }
    });
  } catch (error) {
    return Promise.reject(new Error('Invalid token'));
  }
};

// Boards - FIXED: Changed 'name' to 'title' to match Go backend
export const getBoards = () => api.get('/boards');
export const getBoard = (id) => api.get(`/boards/${id}`);
export const createBoard = (data) => {
  // Convert 'name' to 'title' for Go backend
  const payload = {
    title: data.name || data.title,
    description: data.description || ''
  };
  return api.post('/boards', payload);
};
export const updateBoard = (id, data) => {
  // Convert 'name' to 'title' for Go backend
  const payload = {
    title: data.name || data.title,
    description: data.description
  };
  return api.put(`/boards/${id}`, payload);
};
export const deleteBoard = (id) => api.delete(`/boards/${id}`);

// Lists
export const getLists = (boardId) => api.get(`/lists/board/${boardId}`);
export const createList = (data) => api.post('/lists', data);
export const updateList = (id, data) => api.put(`/lists/${id}`, data);
export const moveList = (id, position) => api.post(`/lists/${id}/move`, { position });
export const deleteList = (id) => api.delete(`/lists/${id}`);

// Cards
export const getCards = (listId) => api.get(`/cards/list/${listId}`);
export const getCard = (id) => api.get(`/cards/${id}`);
export const createCard = (data) => api.post('/cards', data);
export const updateCard = (id, data) => api.put(`/cards/${id}`, data);
export const moveCard = (id, data) => api.post(`/cards/${id}/move`, data);
export const deleteCard = (id) => api.delete(`/cards/${id}`);

// Comments
export const getComments = (cardId) => api.get(`/comments/card/${cardId}`);
export const createComment = (cardId, content) => api.post(`/comments/card/${cardId}`, { content });
export const updateComment = (id, content) => api.put(`/comments/${id}`, { content });
export const deleteComment = (id) => api.delete(`/comments/${id}`);

// Labels
export const getLabels = (boardId) => api.get(`/labels/board/${boardId}`);
export const createLabel = (boardId, data) => api.post(`/labels/board/${boardId}`, data);
export const addLabelToCard = (cardId, labelId) => api.post(`/labels/card/${cardId}`, { label_id: labelId });
export const removeLabelFromCard = (cardId, labelId) => api.delete(`/labels/card/${cardId}/${labelId}`);

// Attachments
export const uploadAttachment = (cardId, file) => {
  const formData = new FormData();
  formData.append('file', file);
  return api.post(`/attachments/card/${cardId}`, formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  });
};
export const downloadAttachment = (id) => api.get(`/attachments/${id}/download`, { responseType: 'blob' });
export const deleteAttachment = (id) => api.delete(`/attachments/${id}`);

// Search
export const searchCards = (params) => api.get('/search/cards', { params });

export default api;
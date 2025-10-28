import { useEffect, useRef, useState, useCallback } from 'react';

export const useWebSocket = (boardId, onMessage) => {
  const ws = useRef(null);
  const [connected, setConnected] = useState(false);
  const reconnectTimeoutRef = useRef(null);

  // Memoize onMessage to prevent reconnection loops
  const stableOnMessage = useCallback(onMessage, []);

  useEffect(() => {
    if (!boardId) return;

    const token = localStorage.getItem('token');
    if (!token) {
      console.error('No authentication token found');
      return;
    }

    // Use environment variable or fallback to localhost
    const WS_BASE_URL = import.meta.env.VITE_WS_URL || 'ws://localhost:8082/api/v1/ws';
    
    // Include token in URL for authentication
    const wsUrl = `${WS_BASE_URL}?board_id=${boardId}&token=${token}`;
    
    console.log('🔌 Connecting to WebSocket:', WS_BASE_URL);
    ws.current = new WebSocket(wsUrl);

    ws.current.onopen = () => {
      console.log('✅ WebSocket connected successfully');
      setConnected(true);
      // Clear any reconnection attempts
      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current);
      }
    };

    ws.current.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log('📨 WebSocket message received:', message);
        stableOnMessage(message);
      } catch (error) {
        console.error('❌ WebSocket message parse error:', error);
      }
    };

    ws.current.onerror = (error) => {
      console.error('❌ WebSocket error:', error);
      setConnected(false);
    };

    ws.current.onclose = (event) => {
      console.log('🔌 WebSocket disconnected', event.code, event.reason);
      setConnected(false);
    };

    // Cleanup function
    return () => {
      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current);
      }
      if (ws.current && ws.current.readyState === WebSocket.OPEN) {
        ws.current.close();
      }
    };
  }, [boardId, stableOnMessage]);

  // Send message function
  const sendMessage = useCallback((message) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(message));
      console.log('📤 Sent WebSocket message:', message);
    } else {
      console.warn('⚠️ WebSocket is not connected');
    }
  }, []);

  return { 
    connected, 
    isConnected: connected, // Alias for compatibility
    sendMessage 
  };
};
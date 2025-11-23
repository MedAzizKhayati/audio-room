import { useEffect, useRef, useState, useCallback } from "react";
import { WSMessage } from "../types";

interface UseAudioWebSocketProps {
  url: string;
  username: string;
  onMessage: (msg: WSMessage) => void;
  onOpen?: () => void;
  onClose?: () => void;
}

export const useAudioWebSocket = ({
  url,
  username,
  onMessage,
  onOpen,
  onClose,
}: UseAudioWebSocketProps) => {
  const wsRef = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);
  const reconnectTimeoutRef = useRef<NodeJS.Timeout | null>(null);

  const connect = useCallback(() => {
    try {
      if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
        return; // Already connected
      }
      const ws = new WebSocket(url);
      wsRef.current = ws;

      ws.onopen = () => {
        console.log("WS Connected");
        setIsConnected(true);
        // Send join message immediately
        ws.send(
          JSON.stringify({
            type: "join",
            sender: username,
          })
        );
        onOpen?.();
      };

      ws.onmessage = (event) => {
        try {
          const msg: WSMessage = JSON.parse(event.data);
          onMessage(msg);
        } catch (e) {
          console.error("Failed to parse WS message", e);
        }
      };

      ws.onclose = (event) => {
        console.log("WS Disconnected", event.code, event.reason);
        setIsConnected(false);
        onClose?.();

        // Attempt reconnect if not closed cleanly (1000)
        if (event.code !== 1000) {
          reconnectTimeoutRef.current = setTimeout(() => {
            console.log("Attempting reconnect...");
            connect();
          }, 3000);
        }
      };

      ws.onerror = (error) => {
        console.error("WS Error", error);
        ws.close();
      };
    } catch (e) {
      console.error("WS Connection error", e);
    }
  }, [url, username, onMessage, onOpen, onClose]);

  useEffect(() => {
    connect();
    return () => {
      if (reconnectTimeoutRef.current)
        clearTimeout(reconnectTimeoutRef.current);
      wsRef.current?.close(1000, "Component unmounting");
    };
  }, [connect]);

  const sendMessage = useCallback((msg: WSMessage) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify(msg));
    }
  }, []);

  return { isConnected, sendMessage };
};

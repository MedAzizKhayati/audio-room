"use client";

import React, { useRef, useState, useCallback } from "react";
import { Mic, MicOff, Users, LogOut, Volume2 } from "lucide-react";
import { useAudioWebSocket } from "../hooks/useAudioWebSocket";
import { useAudioRecorder } from "../hooks/useAudioRecorder";
import { useAudioPlayer } from "../hooks/useAudioPlayer";
import { useVisualizer } from "../hooks/useVisualizer";
import { useUsername } from "@/contexts/UsernameContext";
import { API_CONFIG } from "@/config/api";

export const AudioRoom: React.FC = () => {
  const { username, setUsername: setGlobalUsername } = useUsername();
  const [userCount, setUserCount] = useState(1);
  const canvasRef = useRef<HTMLCanvasElement>(null);

  // 1. Audio Player (Receiving & Playing)
  const { playAudio, activeSpeakers } = useAudioPlayer();

  // 2. Audio Recorder (Microphone & Sending)
  const { isRecording, startRecording, stopRecording, analyser } =
    useAudioRecorder({
      onAudioData: (base64Data) => {
        sendMessage({
          type: "audio",
          payload: base64Data,
        });
      },
    });

  // 3. WebSocket Connection
  const { sendMessage } = useAudioWebSocket({
    url: API_CONFIG.wsUrl,
    username: username || '',
    onMessage: (msg) => {
      if (msg.type === "users" && msg.count) {
        setUserCount(msg.count);
      } else if (msg.type === "audio" && msg.payload && msg.sender) {
        playAudio(msg.payload, msg.sender);
      }
    },
    // onClose: onLeave,
  });

  // 4. Visualizer
  useVisualizer(analyser, canvasRef);

  const toggleMic = useCallback(async () => {
    if (isRecording) {
      stopRecording();
    } else {
      try {
        await startRecording();
      } catch (e) {
        alert("Could not access microphone. Please check permissions.");
      }
    }
  }, [isRecording, startRecording, stopRecording]);

  return (
    <div className="flex flex-col h-dvh bg-gray-900 text-white overflow-hidden">
      {/* Header */}
      <header className="flex items-center justify-between p-4 bg-gray-800 border-b border-gray-700 shrink-0">
        <div className="flex items-center gap-2">
          <Volume2 className="text-indigo-500" />
          <h1 className="font-bold text-lg">Audio Room</h1>
        </div>
        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2 px-3 py-1 bg-gray-700 rounded-full">
            <Users size={16} className="text-gray-400" />
            <span className="text-sm font-medium">{userCount} Online</span>
          </div>
          <button
            onClick={() => setGlobalUsername(null)}
            className="p-2 hover:bg-gray-700 rounded-full transition-colors text-red-400"
            title="Leave Room"
          >
            <LogOut size={20} />
          </button>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 flex flex-col items-center justify-center p-4 relative overflow-hidden gap-6">
        {/* Active Speakers List */}
        <div className="flex flex-wrap justify-center gap-2 min-h-10">
          {Array.from(activeSpeakers).map((speaker) => (
            <div
              key={speaker}
              className="bg-indigo-600/80 px-4 py-2 rounded-full backdrop-blur-sm animate-pulse flex items-center gap-2 shadow-lg shadow-indigo-500/20"
            >
              <div className="w-2 h-2 bg-white rounded-full animate-bounce" />
              <span className="font-medium text-sm">{speaker}</span>
            </div>
          ))}
        </div>

        {/* Visualizer Canvas */}
        <div className="w-full max-w-2xl h-64 bg-gray-800/50 rounded-2xl border border-gray-700 overflow-hidden relative flex items-center justify-center shadow-2xl">
          {!isRecording && (
            <div className="absolute text-gray-500 flex flex-col items-center gap-2 z-10">
              <MicOff size={32} />
              <p>Microphone is off</p>
            </div>
          )}
          <canvas
            ref={canvasRef}
            width={600}
            height={256}
            className={`w-full h-full transition-opacity duration-300 ${
              !isRecording ? "opacity-20" : "opacity-100"
            }`}
          />
        </div>
      </main>

      {/* Controls */}
      <footer className="p-6 bg-gray-800 border-t border-gray-700 flex justify-center shrink-0 pb-safe">
        <button
          onClick={toggleMic}
          className={`
            p-6 rounded-full shadow-xl transition-all transform hover:scale-105 active:scale-95
            ${
              isRecording
                ? "bg-red-500 hover:bg-red-600 shadow-red-500/30 ring-4 ring-red-500/20"
                : "bg-indigo-600 hover:bg-indigo-700 shadow-indigo-500/30"
            }
          `}
        >
          {isRecording ? (
            <MicOff size={32} className="text-white" />
          ) : (
            <Mic size={32} className="text-white" />
          )}
        </button>
      </footer>
    </div>
  );
};

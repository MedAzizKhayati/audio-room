"use client";

import React, { useState } from 'react';
import { Mic } from 'lucide-react';
import { useUsername } from '@/contexts/UsernameContext';

export const JoinScreen: React.FC = () => {
  const { setUsername: setGlobalUsername } = useUsername();
  const [username, setUsername] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (username.trim()) {
      setGlobalUsername(username.trim());
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-dvh bg-gray-900 p-4">
      <div className="bg-gray-800 p-8 rounded-2xl shadow-2xl w-full max-w-md border border-gray-700">
        <div className="flex justify-center mb-8">
          <div className="p-4 bg-indigo-600 rounded-full shadow-lg shadow-indigo-500/30">
            <Mic size={48} className="text-white" />
          </div>
        </div>
        
        <h1 className="text-3xl font-bold text-center text-white mb-2">Audio Room</h1>
        <p className="text-gray-400 text-center mb-8">Enter your name to join the conversation</p>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label htmlFor="username" className="block text-sm font-medium text-gray-300 mb-2">
              Username
            </label>
            <input
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-3 bg-gray-700 border border-gray-600 rounded-xl text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all"
              placeholder="e.g. Alice"
              autoFocus
            />
          </div>

          <button
            type="submit"
            disabled={!username.trim()}
            className="w-full py-3.5 px-4 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-xl shadow-lg shadow-indigo-500/30 transition-all disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-[1.02] active:scale-[0.98]"
          >
            Join Room
          </button>
        </form>
      </div>
    </div>
  );
};

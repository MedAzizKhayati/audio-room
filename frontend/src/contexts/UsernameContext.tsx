"use client";

import React, { createContext, useContext, useEffect, useState } from 'react';

interface UsernameContextType {
  username: string | null;
  setUsername: (username: string | null) => void;
}

const UsernameContext = createContext<UsernameContextType | undefined>(undefined);

export const UsernameProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [username, setUsernameState] = useState<string | null>(null);
  const [isHydrated, setIsHydrated] = useState(false);

  // Load from localStorage on mount
  useEffect(() => {
    const savedUsername = localStorage.getItem('audioroom_username');
    if (savedUsername) {
      setUsernameState(savedUsername);
    }
    setIsHydrated(true);
  }, []);

  const setUsername = (newUsername: string | null) => {
    setUsernameState(newUsername);
    if (newUsername) {
      localStorage.setItem('audioroom_username', newUsername);
    } else {
      localStorage.removeItem('audioroom_username');
    }
  };

  // Prevent rendering until hydrated to avoid hydration mismatch
  if (!isHydrated) {
    return null;
  }

  return (
    <UsernameContext.Provider value={{ username, setUsername }}>
      {children}
    </UsernameContext.Provider>
  );
};

export const useUsername = () => {
  const context = useContext(UsernameContext);
  if (!context) {
    throw new Error('useUsername must be used within UsernameProvider');
  }
  return context;
};

"use client";

import { useUsername } from '@/contexts/UsernameContext';
import { JoinScreen } from './JoinScreen';
import { AudioRoom } from './AudioRoom';

export const HomeContent: React.FC = () => {
  const { username } = useUsername();

  if (!username) {
    return <JoinScreen />;
  }

  return <AudioRoom />;
};

"use client";

import { UsernameProvider } from '@/contexts/UsernameContext';
import { HomeContent } from '@/components/HomeContent';

export default function Home() {
  return (
    <UsernameProvider>
      <HomeContent />
    </UsernameProvider>
  );
}

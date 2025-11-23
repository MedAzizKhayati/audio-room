import { useRef, useEffect, useCallback, useState } from 'react';

export const useAudioPlayer = () => {
  const playbackContextRef = useRef<AudioContext | null>(null);
  const [activeSpeakers, setActiveSpeakers] = useState<Set<string>>(new Set());
  const speakerTimeoutsRef = useRef<Map<string, NodeJS.Timeout>>(new Map());

  useEffect(() => {
    playbackContextRef.current = new (window.AudioContext || (window as any).webkitAudioContext)();
    return () => {
      playbackContextRef.current?.close();
    };
  }, []);

  const playAudio = useCallback(async (base64Data: string, sender: string) => {
    try {
      const ctx = playbackContextRef.current;
      if (!ctx) return;

      if (ctx.state === 'suspended') {
        await ctx.resume();
      }

      // Update active speakers
      setActiveSpeakers(prev => {
        const newSet = new Set(prev);
        newSet.add(sender);
        return newSet;
      });

      // Clear existing timeout for this speaker
      if (speakerTimeoutsRef.current.has(sender)) {
        clearTimeout(speakerTimeoutsRef.current.get(sender)!);
      }

      // Set new timeout to remove speaker
      const timeout = setTimeout(() => {
        setActiveSpeakers(prev => {
          const newSet = new Set(prev);
          newSet.delete(sender);
          return newSet;
        });
        speakerTimeoutsRef.current.delete(sender);
      }, 1000); // 1 second "talking" status

      speakerTimeoutsRef.current.set(sender, timeout);

      // Decode and play
      const binaryString = atob(base64Data);
      const len = binaryString.length;
      const bytes = new Uint8Array(len);
      for (let i = 0; i < len; i++) {
        bytes[i] = binaryString.charCodeAt(i);
      }

      const int16Data = new Int16Array(bytes.buffer);
      const float32Data = new Float32Array(int16Data.length);
      for (let i = 0; i < int16Data.length; i++) {
        float32Data[i] = int16Data[i] / (int16Data[i] < 0 ? 0x8000 : 0x7FFF);
      }

      const buffer = ctx.createBuffer(1, float32Data.length, ctx.sampleRate);
      buffer.copyToChannel(float32Data, 0);

      const source = ctx.createBufferSource();
      source.buffer = buffer;
      source.connect(ctx.destination);
      source.start();

    } catch (e) {
      console.error("Error playing audio", e);
    }
  }, []);

  return { playAudio, activeSpeakers };
};
